package pipeline

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/zhangga/luban-go/core/codetarget"
	"github.com/zhangga/luban-go/core/dataloader"
	"github.com/zhangga/luban-go/core/datatarget"
	"github.com/zhangga/luban-go/core/defs"
	"github.com/zhangga/luban-go/core/l10n"
	"github.com/zhangga/luban-go/core/schema"
	"github.com/zhangga/luban-go/core/types"
	"github.com/zhangga/luban-go/core/validator"
)

type DefaultPipeline struct {
}

func NewDefaultPipeline() *DefaultPipeline {
	return &DefaultPipeline{}
}

func (p *DefaultPipeline) Process(args *PipelineArguments) error {
	fmt.Println("Start pipeline processing...")

	// 1. 加载配置 Schema (RawDefs)
	fmt.Println("Step 1: Load schemas...")
	mockCollector := newMockSchemaCollector()
	xmlLoader := schema.NewXmlSchemaLoader(mockCollector)

	// 寻找入口 xml (例如 args.Conf)
	// 暂时为了避免文件不存在报错，我们如果在测试时，可能跳过这步，直接模拟
	if args.Target != "" && args.InputDataDir != "" {
		// 模拟从定义文件加载
		confPath := filepath.Join(args.InputDataDir, "Defines", "__root__.xml")
		if _, err := os.Stat(confPath); err == nil {
			if err := xmlLoader.Load(confPath); err != nil {
				return fmt.Errorf("failed to load schema: %w", err)
			}
		}
	}

	rawAssembly := mockCollector.CreateRawAssembly()

	// 2. 将 RawDefs 转化为 DefAssembly (DefTypes)
	fmt.Println("Step 2: Build def assembly...")
	assembly, err := defs.NewDefAssemblyImpl(rawAssembly, args.Target, nil)
	if err != nil {
		return fmt.Errorf("failed to create assembly: %w", err)
	}

	manifest := codetarget.NewOutputFileManifest()

	// 3. 读取数据 (DataLoader)
	fmt.Println("Step 3: Load data...")
	// 这里真实项目中要通过 DataLoaderManager 加载
	// DataLoader 应该遍历 assembly.ExportTables 获取对应的数据
	tableRecordsMap := make(map[string][]*defs.Record)
	typeFactory := defs.NewTTypeFactory(assembly)

	for _, table := range assembly.ExportTables {
		// 收集表对应的数据文件
		var inputFiles []string
		if len(table.Raw.InputFiles) > 0 {
			inputFiles = table.Raw.InputFiles
		} else {
			inputFiles = []string{table.Name()}
		}

		var allRecords []*defs.Record
		for _, rawPattern := range inputFiles {
			// 支持 filename.xlsx@sheetName 语法
			filePattern := rawPattern
			subAsset := ""
			if idx := strings.LastIndex(rawPattern, "@"); idx != -1 {
				filePattern = rawPattern[:idx]
				subAsset = rawPattern[idx+1:]
			}

			// 支持统配符和目录读取
			fullPattern := filepath.Join(args.InputDataDir, "Datas", filePattern)

			// 如果没有扩展名，我们默认尝试匹配 .xlsx 和 .csv
			if filepath.Ext(fullPattern) == "" {
				fullPattern += ".*"
			}

			matches, err := filepath.Glob(fullPattern)
			if err != nil {
				fmt.Printf("Warning: failed to glob pattern %s: %v\n", fullPattern, err)
				continue
			}

			if len(matches) == 0 {
				// 如果直接匹配不到，尝试加扩展名
				matchesXlsx, _ := filepath.Glob(filepath.Join(args.InputDataDir, "Datas", filePattern+".xlsx"))
				matchesCsv, _ := filepath.Glob(filepath.Join(args.InputDataDir, "Datas", filePattern+".csv"))
				matches = append(matches, matchesXlsx...)
				matches = append(matches, matchesCsv...)
			}

			for _, matchPath := range matches {
				// 排除临时文件，比如 ~$xxxx.xlsx
				if strings.HasPrefix(filepath.Base(matchPath), "~$") {
					continue
				}

				f, err := os.Open(matchPath)
				if err != nil {
					fmt.Printf("Warning: failed to open file %s: %v\n", matchPath, err)
					continue
				}

				var loader dataloader.IDataLoader
				ext := strings.ToLower(filepath.Ext(matchPath))
				if ext == ".xlsx" || ext == ".xlsm" || ext == ".xls" {
					loader = dataloader.NewExcelDataLoader(typeFactory)
				} else if ext == ".csv" {
					loader = dataloader.NewCsvDataLoader(typeFactory)
				}

				if loader != nil {
					fmt.Printf("Loading data for table %s from %s (sheet: %s)\n", table.Name(), matchPath, subAsset)

					if err := loader.Load(matchPath, subAsset, f); err == nil {
						tType, errType := typeFactory.CreateType(table.Raw.ValueType)
						if errType != nil {
							fmt.Printf("Failed to create type for %s: %v\n", table.Raw.ValueType, errType)
						} else if beanType, ok := tType.(*types.TBean); ok {
							records := loader.ReadMulti(beanType)
							allRecords = append(allRecords, records...)
							fmt.Printf("Loaded %d records from %s\n", len(records), matchPath)
						} else {
							fmt.Printf("Type %s is not a bean type, it is %s\n", table.Raw.ValueType, tType.TypeName())
						}
					} else {
						fmt.Printf("Failed to parse file %s: %v\n", matchPath, err)
					}
				}
				f.Close()
			}
		}

		if len(allRecords) > 0 {
			tableRecordsMap[table.FullName()] = allRecords
			fmt.Printf("Total loaded %d records for table %s\n", len(allRecords), table.Name())
		} else {
			fmt.Printf("No data found for table %s\n", table.Name())
		}
	}

	// 3.5 校验数据 (Validator)
	fmt.Println("Step 3.5: Validate data...")
	valCtx := validator.NewValidatorContext(assembly, tableRecordsMap)
	valMgr := validator.NewDataValidatorManager()
	valMgr.Register(&validator.RefValidator{})
	valMgr.Register(&validator.RangeValidator{})
	valMgr.Register(&validator.PathValidator{BaseDir: args.InputDataDir})

	valErrs := valMgr.Validate(valCtx)
	if len(valErrs) > 0 {
		fmt.Println("Validation failed with the following errors:")
		for _, err := range valErrs {
			fmt.Printf(" - %v\n", err)
		}
		return fmt.Errorf("data validation failed with %d errors", len(valErrs))
	}

	// 4. 生成代码 (CodeTarget)
	fmt.Println("Step 4: Generate codes...")
	var target codetarget.ICodeTarget

	lang := "go"
	if len(args.CodeTargets) > 0 && args.CodeTargets[0] != "" {
		lang = args.CodeTargets[0]
	}

	switch lang {
	case "ts":
		t, err := codetarget.NewTSTarget(filepath.Join(args.TemplateDir, "ts"))
		if err != nil {
			return fmt.Errorf("failed to create ts code target: %w", err)
		}
		target = t
	case "lua":
		t, err := codetarget.NewLuaCodeTarget(filepath.Join(args.TemplateDir, "lua"))
		if err != nil {
			return fmt.Errorf("failed to create lua code target: %w", err)
		}
		target = t
	case "cpp":
		t, err := codetarget.NewCppTarget(filepath.Join(args.TemplateDir, "cpp"))
		if err != nil {
			return fmt.Errorf("failed to create cpp code target: %w", err)
		}
		target = t
	case "java":
		t, err := codetarget.NewJavaTarget(filepath.Join(args.TemplateDir, "java"))
		if err != nil {
			return fmt.Errorf("failed to create java code target: %w", err)
		}
		target = t
	case "pb":
		t, err := codetarget.NewPbTarget(filepath.Join(args.TemplateDir, "pb"))
		if err != nil {
			return fmt.Errorf("failed to create pb code target: %w", err)
		}
		target = t
	case "cs":
		t, err := codetarget.NewCSTarget(filepath.Join(args.TemplateDir, "cs"))
		if err != nil {
			return fmt.Errorf("failed to create cs code target: %w", err)
		}
		target = t
	case "go":
		fallthrough
	default:
		target = codetarget.NewGoCodeTarget()
	}

	if err := target.Handle(assembly, manifest); err != nil {
		return fmt.Errorf("failed to generate code: %w", err)
	}

	// 4.5 导出多语言文本 (如果使用了 text 类型)
	fmt.Println("Step 4.5: Export l10n texts...")
	l10nData, err := l10n.GetManager().Export()
	if err != nil {
		fmt.Printf("Warning: failed to export l10n data: %v\n", err)
	} else if string(l10nData) != "{}" && string(l10nData) != "{\n}" {
		manifest.AddDataFile(codetarget.NewOutputFile("l10n.json", l10nData))
	}

	// 5. 导出数据 (DataTarget)
	fmt.Println("Step 5: Export data...")
	targetGroups := assembly.Target.Groups

	dataTargetLang := "json"
	if len(args.DataTargets) > 0 && args.DataTargets[0] != "" {
		dataTargetLang = args.DataTargets[0]
	}

	var dTarget datatarget.IDataTarget
	if dataTargetLang == "lua" {
		dTarget = datatarget.NewLuaDataTarget(targetGroups)
	} else if dataTargetLang == "bin" {
		dTarget = datatarget.NewBinDataTarget(targetGroups)
	} else {
		dTarget = datatarget.NewJsonDataTarget(targetGroups)
	}

	for _, table := range assembly.ExportTables {
		records := tableRecordsMap[table.FullName()]
		out, err := dTarget.ExportTable(table, records)
		if err != nil {
			return fmt.Errorf("failed to export data: %w", err)
		}
		manifest.AddDataFile(out)
	}

	// 6. 保存产出文件
	fmt.Println("Step 6: Save output files...")
	if err := saveFiles(args.OutputCodeDir, manifest.CodeFiles); err != nil {
		return err
	}
	if err := saveFiles(args.OutputDataDir, manifest.DataFiles); err != nil {
		return err
	}

	return nil
}

func saveFiles(baseDir string, files []*codetarget.OutputFile) error {
	if baseDir == "" {
		return nil
	}
	os.MkdirAll(baseDir, os.ModePerm)
	for _, f := range files {
		path := filepath.Join(baseDir, f.File)
		if err := ioutil.WriteFile(path, f.Content, 0666); err != nil {
			return fmt.Errorf("failed to write file %s: %w", path, err)
		}
	}
	return nil
}
