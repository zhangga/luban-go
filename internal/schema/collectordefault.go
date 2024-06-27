package schema

import (
	"github.com/zhangga/luban/core/lubanconf"
	"github.com/zhangga/luban/core/manager"
	"github.com/zhangga/luban/core/pipeline"
	"github.com/zhangga/luban/core/schema"
	"github.com/zhangga/luban/pkg/logger"
	"path/filepath"
	"sync"
)

var _ schema.ISchemaCollector = (*DefaultSchemaCollector)(nil)

// DefaultSchemaCollector 默认SchemaCollector
type DefaultSchemaCollector struct {
	CollectorBase
	logger   logger.Logger
	pipeline pipeline.IPipeline
	config   *lubanconf.LubanConfig
}

func NewDefaultSchemaCollector(logger logger.Logger, pipeline pipeline.IPipeline) schema.ISchemaCollector {
	return &DefaultSchemaCollector{
		logger:   logger,
		pipeline: pipeline,
	}
}

func (s *DefaultSchemaCollector) Name() string {
	return "default"
}

// Load 加载所有的schema
func (s *DefaultSchemaCollector) Load(config *lubanconf.LubanConfig) {
	s.config = config
	// schema loader
	for _, importFile := range s.config.Imports {
		s.logger.Debugf("import schema file: %s, type: %s", importFile.FileName, importFile.Type)
		schemaMgr := manager.MustGet[*schema.Manager]()
		schemaLoader := schemaMgr.CreateSchemaLoader(filepath.Ext(importFile.FileName), importFile.Type, s)
		schemaLoader.Load(importFile.FileName)
	}

	s.loadTableValueTypeSchemasFromFile()
}

func (s *DefaultSchemaCollector) loadTableValueTypeSchemasFromFile() {
	beanSchemaLoaderName := s.pipeline.Args().GetOptionOrDefault(pipeline.SchemaCollectorFamily, "beanSchemaLoader", "default", true)
	var wg sync.WaitGroup
	for _, t := range s.rawTables {
		if !t.ReadSchemaFromFile {
			continue
		}

		wg.Add(1)
		go func() {
			defer wg.Done()

			fileName := t.InputFiles[0]
			schemaMgr := manager.MustGet[*schema.Manager]()
			beanLoader := schemaMgr.CreateBeanLoader(beanSchemaLoaderName, s)
			fullPath := filepath.Join(s.config.InputDataDir, fileName)
			bean := beanLoader.Load(fullPath, t.ValueType)
			s.AddBean(bean)
		}()
	}
	wg.Done()
}
