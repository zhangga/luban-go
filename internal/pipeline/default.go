package pipeline

import (
	"errors"
	"github.com/zhangga/luban/core/lubanconf"
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
	config *lubanconf.Conf
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
	var err error
	confLoader := lubanconf.NewGlobalConfigLoader(p.logger)
	if p.config, err = confLoader.Load(p.args.ConfFile); err != nil {
		p.logger.Errorf("load config file %s, failed: %s", p.args.ConfFile, err)
		return err
	}

	schemaMgr, ok := manager.Get[*schema.Manager]()
	if !ok {
		return errors.New("schema manager not found")
	}

	_ = schemaMgr
	//schemaCollector := schemaMgr.CreateSchemaCollector(p.args.SchemaCollector)
	//schemaCollector.Load()
	return nil
}
