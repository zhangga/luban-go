package schema

import (
	"github.com/zhangga/luban/core/lubanconf"
	"github.com/zhangga/luban/core/manager"
	"github.com/zhangga/luban/core/pipeline"
	"github.com/zhangga/luban/core/schema"
	"github.com/zhangga/luban/pkg/logger"
	"path/filepath"
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

func (s *DefaultSchemaCollector) Load(config *lubanconf.LubanConfig) {
	s.config = config
	for _, importFile := range s.config.Imports {
		s.logger.Debugf("import schema file: %s, type: %s", importFile.FileName, importFile.Type)
		schemaMgr := manager.MustGet[*schema.Manager]()
		schemaLoader := schemaMgr.CreateSchemaLoader(filepath.Ext(importFile.FileName), importFile.Type, s)
		schemaLoader.Load(importFile.FileName)
	}
}
