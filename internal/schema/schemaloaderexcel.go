package schema

import (
	"fmt"
	"github.com/zhangga/luban/core/manager"
	"github.com/zhangga/luban/core/pctx"
	"github.com/zhangga/luban/core/schema"
	"github.com/zhangga/luban/internal/defs"
	"github.com/zhangga/luban/internal/rawdefs"
	"github.com/zhangga/luban/internal/types"
	"github.com/zhangga/luban/internal/utils"
	"github.com/zhangga/luban/pkg/logger"
)

var _ schema.ISchemaLoader = (*ExcelSchemaLoader)(nil)

type ExcelSchemaLoader struct {
	schemaLoaderBase
}

func NewExcelSchemaLoader(logger logger.Logger, dataType string, collector schema.ISchemaCollector) schema.ISchemaLoader {
	return &ExcelSchemaLoader{
		schemaLoaderBase: schemaLoaderBase{
			logger:    logger,
			dataType:  dataType,
			collector: collector,
		},
	}
}

func (e *ExcelSchemaLoader) Load(ctx pctx.Context, fileName string) {
	switch e.dataType {
	case "table":
		e.loadTableListFromFile(ctx, fileName)
	case "bean":
		e.loadBeanListFromFile(ctx, fileName)
	case "enum":
		e.loadEnumListFromFile(ctx, fileName)
	default:
		panic(fmt.Errorf("加载文件: %s, 未知的数据类型: %s", fileName, e.dataType))
	}
}

func (e *ExcelSchemaLoader) loadTableListFromFile(ctx pctx.Context, fileName string) {
	defTableRecordType := defs.NewDefBean(rawdefs.RawBean{
		Namespace:   "__intern__",
		Name:        "__TableRecord__",
		IsValueType: false,
		Fields: []*rawdefs.RawField{
			{Name: "full_name", Type: "string"},
			{Name: "value_type", Type: "string"},
			{Name: "index", Type: "string"},
			{Name: "mode", Type: "string"},
			{Name: "group", Type: "string"},
			{Name: "comment", Type: "string"},
			{Name: "read_schema_from_file", Type: "bool"},
			{Name: "input", Type: "string"},
			{Name: "output", Type: "string"},
			{Name: "tags", Type: "string"},
		},
	})
	defTableRecordType.DefTypeBase.Assembly = defs.NewDefAssembly(ctx, e.logger, &rawdefs.RawAssembly{
		Targets: []*rawdefs.RawTarget{{Name: "default", Manager: "Tables"}},
	}, "default", nil)
	defTableRecordType.PreCompile(ctx)
	defTableRecordType.Compile(ctx)
	defTableRecordType.PostCompile(ctx)

	tableRecordType := types.NewTBean(defTableRecordType, false, nil)
	actualFile, sheetName := utils.SplitFileAndSheetName(utils.StandardizePath(fileName))
	// 读取table数据
	records, err := manager.MustGetIface[manager.IDataLoaderManager]().LoadTableFile(tableRecordType, actualFile, sheetName, nil)
	if err != nil {
		panic(err)
	}

	panic("implement me" + fmt.Sprintf("%v", records))
}

func (e *ExcelSchemaLoader) loadBeanListFromFile(ctx pctx.Context, fileName string) {
	panic("implement me")
}

func (e *ExcelSchemaLoader) loadEnumListFromFile(ctx pctx.Context, fileName string) {
	panic("implement me")
}
