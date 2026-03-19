package core

import (
	"fmt"
	"github.com/zhangga/luban-go/core/pipeline"
)

type SimpleLauncher struct {
	// 还可以存放配置、日志等信息
}

func NewSimpleLauncher() *SimpleLauncher {
	return &SimpleLauncher{}
}

func (s *SimpleLauncher) Start(args *pipeline.PipelineArguments) {
	fmt.Println("Luban Go initialized.")

	// TODO: InitManagers 注册各种 SchemaLoader、CodeTarget、DataTarget 等
	s.InitManagers()

	// 启动默认管线
	pipe := pipeline.NewDefaultPipeline()
	if err := pipe.Process(args); err != nil {
		fmt.Printf("Pipeline execution failed: %v\n", err)
	}

	fmt.Println("Luban Go execution completed.")
}

func (s *SimpleLauncher) InitManagers() {
	// SchemaManager.Ins.Init()
	// CodeTargetManager.Ins.Init()
	// DataTargetManager.Ins.Init()
	// DataLoaderManager.Ins.Init()
}
