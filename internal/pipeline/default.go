package pipeline

import (
	"context"
	"errors"
	"fmt"
	"github.com/zhangga/luban/core/lubanconf"
	"github.com/zhangga/luban/core/manager"
	"github.com/zhangga/luban/core/options"
	"github.com/zhangga/luban/core/pctx"
	"github.com/zhangga/luban/core/pipeline"
	"github.com/zhangga/luban/pkg/logger"
)

var _ pipeline.IPipeline = (*DefaultPipeline)(nil)

// DefaultPipeline default pipeline
type DefaultPipeline struct {
	logger logger.Logger
	ctx    *pContext
}

func NewDefaultPipeline(logger logger.Logger) pipeline.IPipeline {
	return &DefaultPipeline{
		logger: logger,
	}
}

func (p *DefaultPipeline) Name() string {
	return "default"
}

func (p *DefaultPipeline) Context() pctx.Context {
	return p.ctx
}

func (p *DefaultPipeline) Run(opts options.CommandOptions) error {
	// 初始化上下文
	if err := p.initContext(opts); err != nil {
		return err
	}

	// 加载schema
	if err := p.loadSchema(); err != nil {
		return err
	}
	return nil
}

func (p *DefaultPipeline) initContext(opts options.CommandOptions) error {
	p.logger.Debugf("prepare pipeline context")

	// 运行参数
	args := options.CreateArguments(opts)

	// 加载luban配置文件
	config, err := p.loadLubanConfig(args.ConfFile)
	if err != nil {
		return err
	}

	ctx := &pContext{
		Context: context.Background(),
		config:  config,
		args:    args,
	}

	p.ctx = ctx
	return nil
}

// loadLubanConfig 加载luban配置文件
func (p *DefaultPipeline) loadLubanConfig(confFile string) (*lubanconf.LubanConfig, error) {
	confLoader := lubanconf.NewGlobalConfigLoader(p.logger)
	// 加载luban配置文件
	if config, err := confLoader.Load(confFile); err != nil {
		return nil, fmt.Errorf("load config file %s, failed: %w", confFile, err)
	} else {
		return config, nil
	}
}

func (p *DefaultPipeline) loadSchema() error {
	schemaMgr, ok := manager.GetIface[manager.ISchemaManager]()
	if !ok {
		return errors.New("schema manager not found")
	}

	p.logger.Infof("load schema.collector: %s, path: %s", p.ctx.args.SchemaCollector, p.ctx.args.ConfFile)
	schemaCollector := schemaMgr.CreateSchemaCollector(p.ctx.args.SchemaCollector, p)
	schemaCollector.Load(p.ctx)
	return nil
}
