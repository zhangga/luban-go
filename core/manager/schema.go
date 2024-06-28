package manager

import (
	"github.com/modern-go/reflect2"
	"github.com/zhangga/luban/core/pipeline"
	"github.com/zhangga/luban/core/schema"
)

type ISchemaManager interface {
	IManager
	MustEmbedSchemaManager

	CreateSchemaCollector(name string, pipeline pipeline.IPipeline) schema.ISchemaCollector
	CreateSchemaLoader(extName, dataType string, schemaCollector schema.ISchemaCollector) schema.ISchemaLoader
	CreateBeanLoader(name string, collector schema.ISchemaCollector) schema.IBeanSchemaLoader
}

type MustEmbedSchemaManager interface {
	mustEmbedSchemaManager()
}

type EmbedSchemaManager struct {
}

func (EmbedSchemaManager) mustEmbedSchemaManager() {}

func (EmbedSchemaManager) key() uintptr {
	rtype := reflect2.RTypeOf((*ISchemaManager)(nil))
	return rtype
}
