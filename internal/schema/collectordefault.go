package schema

import (
	"github.com/zhangga/luban/core/manager"
	"github.com/zhangga/luban/core/pctx"
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

func (s *DefaultSchemaCollector) Pipeline() pipeline.IPipeline {
	return s.pipeline
}

// Load 加载所有的schema
func (s *DefaultSchemaCollector) Load(ctx pctx.Context) {
	// schema loader
	for _, importFile := range ctx.Config().Imports {
		s.logger.Debugf("import schema file: %s, type: %s", importFile.FileName, importFile.Type)
		schemaMgr := manager.MustGetIface[manager.ISchemaManager]()
		schemaLoader := schemaMgr.CreateSchemaLoader(filepath.Ext(importFile.FileName), importFile.Type, s)
		schemaLoader.Load(ctx, importFile.FileName)
	}

	s.loadTableValueTypeSchemasFromFile(ctx)
}

func (s *DefaultSchemaCollector) loadTableValueTypeSchemasFromFile(ctx pctx.Context) {
	beanSchemaLoaderName := ctx.Args().GetOptionOrDefault(pipeline.SchemaCollectorFamily, "beanSchemaLoader", "default", true)
	var wg sync.WaitGroup
	for _, t := range s.rawTables {
		if !t.ReadSchemaFromFile {
			continue
		}

		wg.Add(1)
		go func() {
			defer wg.Done()

			fileName := t.InputFiles[0]
			schemaMgr := manager.MustGetIface[manager.ISchemaManager]()
			beanLoader := schemaMgr.CreateBeanLoader(beanSchemaLoaderName, s)
			fullPath := filepath.Join(ctx.Config().InputDataDir, fileName)
			bean := beanLoader.Load(fullPath, t.ValueType)
			s.AddBean(bean)
		}()
	}
	wg.Done()
}
