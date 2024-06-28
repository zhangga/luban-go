package pipeline

import (
	"errors"
	"fmt"
	"github.com/zhangga/luban/core/lubanconf"
	"github.com/zhangga/luban/core/manager"
	"github.com/zhangga/luban/core/pipeline"
	"github.com/zhangga/luban/pkg/logger"
)

var _ pipeline.IPipeline = (*DefaultPipeline)(nil)

// DefaultPipeline default pipeline
type DefaultPipeline struct {
	logger logger.Logger
	args   pipeline.Arguments
	config *lubanconf.LubanConfig
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
	p.loadLubanConfig()

	if err := p.loadSchema(); err != nil {
		return err
	}
	return nil
}

func (p *DefaultPipeline) Args() pipeline.Arguments {
	return p.args
}

func (p *DefaultPipeline) Config() *lubanconf.LubanConfig {
	return p.config
}

func (p *DefaultPipeline) loadLubanConfig() {
	confLoader := lubanconf.NewGlobalConfigLoader(p.logger)
	// 加载luban配置文件
	if config, err := confLoader.Load(p.args.ConfFile); err != nil {
		panic(fmt.Errorf("load config file %s, failed: %w", p.args.ConfFile, err))
	} else {
		p.config = config
	}
}

func (p *DefaultPipeline) loadSchema() error {
	schemaMgr, ok := manager.GetIface[manager.ISchemaManager]()
	if !ok {
		return errors.New("schema manager not found")
	}

	p.logger.Infof("load schema.collector: %s, path: %s", p.args.SchemaCollector, p.args.ConfFile)
	schemaCollector := schemaMgr.CreateSchemaCollector(p.args.SchemaCollector, p)
	schemaCollector.Load()
	return nil
}
