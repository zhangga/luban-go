package schema

import (
	"fmt"
	"github.com/zhangga/luban/core/manager"
	"github.com/zhangga/luban/core/pipeline"
	"github.com/zhangga/luban/pkg/logger"
)

var _ manager.IManager = (*Manager)(nil)

type Manager struct {
	logger logger.Logger
}

func (m *Manager) Init(logger logger.Logger) {
	m.logger = logger
}

func (m *Manager) PostInit() {
}

func (m *Manager) CreateSchemaCollector(name string, pipeline pipeline.IPipeline) ISchemaCollector {
	creator := getCollectorCreator(name)
	if creator == nil {
		panic(fmt.Errorf("SchemaCollector %s not found", name))
	}
	return creator(m.logger, pipeline)
}

func (m *Manager) CreateSchemaLoader(extName, dataType string, schemaCollector ISchemaCollector) ISchemaLoader {
	loaderInfo := getSchemaLoaderInfo(dataType, extName)
	if loaderInfo == nil {
		panic(fmt.Errorf("SchemaLoader type:%s, extName:%s not found", dataType, extName))
	}
	loader := loaderInfo.Creator(m.logger, dataType, schemaCollector)
	return loader
}

func (m *Manager) CreateBeanLoader(name string, collector ISchemaCollector) IBeanSchemaLoader {
	creator := getBeanLoaderCreator(name)
	if creator == nil {
		panic(fmt.Errorf("BeanLoader %s not found", name))
	}
	return creator(m.logger, collector)
}
