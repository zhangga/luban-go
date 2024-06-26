package schema

import (
	"github.com/zhangga/luban/core/schema"
	"github.com/zhangga/luban/pkg/logger"
)

var _ schema.ISchemaLoader = (*ExcelSchemaLoader)(nil)

type ExcelSchemaLoader struct {
	logger    logger.Logger
	dataType  string
	collector schema.ISchemaCollector
}

func NewExcelSchemaLoader(logger logger.Logger, dataType string, collector schema.ISchemaCollector) schema.ISchemaLoader {
	return &ExcelSchemaLoader{
		logger:    logger,
		dataType:  dataType,
		collector: collector,
	}
}

func (e ExcelSchemaLoader) Load(fileName string) {
	//TODO implement me
	panic("implement me")
}
