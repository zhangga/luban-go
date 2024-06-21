package core

import (
	"fmt"
	"github.com/zhangga/luban/core/manager"
	"github.com/zhangga/luban/core/options"
	"github.com/zhangga/luban/core/pipeline"
	"github.com/zhangga/luban/pkg/logger"
	"strings"
)

// SimpleLauncher 启动器
type SimpleLauncher struct {
	logger  logger.Logger
	options map[string]string // xargs参数
}

func NewSimpleLauncher(logger logger.Logger) *SimpleLauncher {
	return &SimpleLauncher{
		logger: logger,
	}
}

func (s *SimpleLauncher) Start(opts options.CommandOptions) {
	s.options = parseOptions(opts.Xargs...)
	s.initManagers()

	pipeMgr, ok := manager.Get[*pipeline.Manager]()
	if !ok {
		panic("pipeline manager not found")
	}
	pipe := pipeMgr.CreatePipeline(opts.Pipeline)
	pipe.Run(pipeline.CreateArguments(opts))
}

// initManagers 初始化管理器
func (s *SimpleLauncher) initManagers() {
	manager.Traverse(func(mgr manager.IManager) {
		mgr.Init(s.logger)
	})
	//s.schemaMgr = schema.NewManager(s.logger)
	//s.schemaMgr.Init()
	//s.pipelineMgr = pipeline.NewManager(s.logger)
	//s.pipelineMgr.Init()
}

func parseOptions(xargs ...string) map[string]string {
	options := make(map[string]string)
	for _, arg := range xargs {
		kv := strings.Split(arg, "=")
		if len(kv) != 2 {
			err := fmt.Errorf("invalid xargs: %s", arg)
			panic(err)
		}
		if _, ok := options[kv[0]]; ok {
			err := fmt.Errorf("duplicate xargs: %s", arg)
			panic(err)
		}
		options[kv[0]] = kv[1]
	}
	return options
}
