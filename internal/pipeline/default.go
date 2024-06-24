package pipeline

import (
	"errors"
	"github.com/zhangga/luban/core/manager"
	"github.com/zhangga/luban/core/pipeline"
	"github.com/zhangga/luban/core/schema"
	"github.com/zhangga/luban/pkg/logger"
)

var _ pipeline.IPipeline = (*DefaultPipeline)(nil)

// DefaultPipeline default pipeline
type DefaultPipeline struct {
	logger logger.Logger
	args   pipeline.Arguments
}

func NewDefaultPipeline(logger logger.Logger) pipeline.IPipeline {
	return &DefaultPipeline{
		logger: logger,
	}
}

func (p *DefaultPipeline) Name() string {
	return "default"
}

func (p *DefaultPipeline) Run(args pipeline.Arguments) error {
	p.args = args

	if err := p.loadSchema(); err != nil {
		return err
	}
	return nil
}

func (p *DefaultPipeline) loadSchema() error {
	schemaMgr, ok := manager.Get[*schema.Manager]()
	if !ok {
		return errors.New("schema manager not found")
	}

	schemaCollector := schemaMgr.CreateSchemaCollector(p.args.SchemaCollector)
	schemaCollector.Load()
	return nil
}
