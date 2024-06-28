package schema

import (
	"fmt"
	"github.com/zhangga/luban/core/manager"
	"github.com/zhangga/luban/core/pipeline"
	"github.com/zhangga/luban/core/schema"
	"github.com/zhangga/luban/pkg/logger"
)

var _ manager.ISchemaManager = (*Manager)(nil)

type Manager struct {
	manager.EmbedSchemaManager
	logger logger.Logger
}

func (m *Manager) Init(logger logger.Logger) {
	m.logger = logger
}

func (m *Manager) PostInit() {
}

func (m *Manager) CreateSchemaCollector(name string, pipeline pipeline.IPipeline) schema.ISchemaCollector {
	creator := schema.GetCollectorCreator(name)
	if creator == nil {
		panic(fmt.Errorf("SchemaCollector %s not found", name))
	}
	return creator(m.logger, pipeline)
}

func (m *Manager) CreateSchemaLoader(extName, dataType string, schemaCollector schema.ISchemaCollector) schema.ISchemaLoader {
	loaderInfo := schema.GetSchemaLoaderInfo(dataType, extName)
	if loaderInfo == nil {
		panic(fmt.Errorf("SchemaLoader type:%s, extName:%s not found", dataType, extName))
	}
	loader := loaderInfo.Creator(m.logger, dataType, schemaCollector)
	return loader
}

func (m *Manager) CreateBeanLoader(name string, collector schema.ISchemaCollector) schema.IBeanSchemaLoader {
	creator := schema.GetBeanLoaderCreator(name)
	if creator == nil {
		panic(fmt.Errorf("BeanLoader %s not found", name))
	}
	return creator(m.logger, collector)
}
