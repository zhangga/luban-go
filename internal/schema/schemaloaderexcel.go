package schema

import (
	"fmt"
	"github.com/zhangga/luban/core/schema"
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
	panic("implement me")
}

func (e *ExcelSchemaLoader) loadBeanListFromFile(fileName string) {
	panic("implement me")
}

func (e *ExcelSchemaLoader) loadEnumListFromFile(fileName string) {
	panic("implement me")
}
