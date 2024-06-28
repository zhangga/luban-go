package core

import (
	"github.com/zhangga/luban/core/manager"
	"github.com/zhangga/luban/core/options"
	"github.com/zhangga/luban/core/pipeline"
	"github.com/zhangga/luban/pkg/logger"
)

// SimpleLauncher 启动器
type SimpleLauncher struct {
	logger logger.Logger
}

func NewSimpleLauncher(logger logger.Logger) *SimpleLauncher {
	return &SimpleLauncher{
		logger: logger,
	}
}

func (s *SimpleLauncher) Start(opts options.CommandOptions) {
	s.initManagers()

	pipeMgr, ok := manager.GetIface[manager.IPipelineManager]()
	if !ok {
		panic("pipeline manager not found")
	}
	pipe := pipeMgr.CreatePipeline(opts.Pipeline)
	if err := pipe.Run(pipeline.CreateArguments(opts)); err != nil {
		panic(err)
	}

	s.logger.Info("bye~")
}

// initManagers 初始化管理器
func (s *SimpleLauncher) initManagers() {
	manager.Traverse(func(mgr manager.IManager) {
		mgr.Init(s.logger)
	})
	manager.Traverse(func(mgr manager.IManager) {
		mgr.PostInit()
	})
}
