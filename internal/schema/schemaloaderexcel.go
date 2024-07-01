package schema

import (
	"fmt"
	"github.com/zhangga/luban/core/schema"
	"github.com/zhangga/luban/internal/defs"
	"github.com/zhangga/luban/internal/rawdefs"
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

func (e *ExcelSchemaLoader) Load(fileName string) {
	switch e.dataType {
	case "table":
		e.loadTableListFromFile(fileName)
	case "bean":
		e.loadBeanListFromFile(fileName)
	case "enum":
		e.loadEnumListFromFile(fileName)
	default:
		panic(fmt.Errorf("加载文件: %s, 未知的数据类型: %s", fileName, e.dataType))
	}
}

func (e *ExcelSchemaLoader) loadTableListFromFile(fileName string) {
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
	defTableRecordType.DefTypeBase.Assembly = defs.NewDefAssembly(e.logger, &rawdefs.RawAssembly{
		Targets: []*rawdefs.RawTarget{{Name: "default", Manager: "Tables"}},
	}, "default", nil)
	defTableRecordType.PreCompile()
	defTableRecordType.Compile()
	defTableRecordType.PostCompile()
}

func (e *ExcelSchemaLoader) loadBeanListFromFile(fileName string) {
	panic("implement me")
}

func (e *ExcelSchemaLoader) loadEnumListFromFile(fileName string) {
	panic("implement me")
}
