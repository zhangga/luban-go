package pipeline

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/zhangga/luban-go/core/codetarget"
	"github.com/zhangga/luban-go/core/dataloader"
	"github.com/zhangga/luban-go/core/datatarget"
	"github.com/zhangga/luban-go/core/defs"
	"github.com/zhangga/luban-go/core/schema"
	"github.com/zhangga/luban-go/core/types"
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
		// 假设数据文件在 args.InputDataDir 下与表名相同，不分大小写处理或直接映射
		// 在我们的测试例子里表名叫 TbItem, 所以文件是 TbItem.xlsx
		dataPath := filepath.Join(args.InputDataDir, "Datas", table.Name()+".xlsx")
		fmt.Printf("Loading data for table %s from %s\n", table.Name(), dataPath)
		if _, err := os.Stat(dataPath); err == nil {
			f, _ := os.Open(dataPath)

			loader := dataloader.NewExcelDataLoader(typeFactory)
			if err := loader.Load(dataPath, "", f); err == nil {
				// 获取此表关联的 bean
				tType, errType := defs.NewTTypeFactory(assembly).CreateType(table.Raw.ValueType)
				if errType != nil {
					fmt.Printf("Failed to create type for %s: %v\n", table.Raw.ValueType, errType)
				}
				if tType != nil {
					if beanType, ok := tType.(*types.TBean); ok {
						records := loader.ReadMulti(beanType)
						tableRecordsMap[table.FullName()] = records
						fmt.Printf("Loaded %d records for table %s\n", len(records), table.Name())
					} else {
						fmt.Printf("Type %s is not a bean type, it is %s\n", table.Raw.ValueType, tType.TypeName())
					}
				}
			} else {
				fmt.Printf("Failed to parse excel %s: %v\n", dataPath, err)
			}
			f.Close()
		} else {
			fmt.Printf("Data file %s not found\n", dataPath)
		}
	}

	// 4. 生成代码 (CodeTarget)
	fmt.Println("Step 4: Generate codes...")
	var target codetarget.ICodeTarget

	lang := "go"
	if len(args.CodeTargets) > 0 && args.CodeTargets[0] != "" {
		lang = args.CodeTargets[0]
	}

	switch lang {
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

	// 5. 导出数据 (DataTarget)
	fmt.Println("Step 5: Export data...")
	jsonTarget := datatarget.NewJsonDataTarget()
	for _, table := range assembly.ExportTables {
		records := tableRecordsMap[table.FullName()]
		out, err := jsonTarget.ExportTable(table, records)
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
