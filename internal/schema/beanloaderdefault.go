package schema

import (
	"github.com/zhangga/luban/core/refs"
	"github.com/zhangga/luban/core/schema"
	"github.com/zhangga/luban/pkg/logger"
)

var _ schema.IBeanSchemaLoader = (*BeanSchemaFromExcelHeaderLoader)(nil)

type BeanSchemaFromExcelHeaderLoader struct {
	logger    logger.Logger
	collector schema.ISchemaCollector
}

func NewBeanSchemaFromExcelHeaderLoader(logger logger.Logger, collector schema.ISchemaCollector) schema.IBeanSchemaLoader {
	return &BeanSchemaFromExcelHeaderLoader{
		logger:    logger,
		collector: collector,
	}
}

func (b *BeanSchemaFromExcelHeaderLoader) Name() string {
	return "default"
}

func (b *BeanSchemaFromExcelHeaderLoader) Load(fileName, beanFullName string) refs.UnimplementedBean {
	//TODO implement me
	panic("implement me")
}
