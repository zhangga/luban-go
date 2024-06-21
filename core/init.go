package core

import (
	"github.com/zhangga/luban/core/manager"
	"github.com/zhangga/luban/core/pipeline"
	pipeline2 "github.com/zhangga/luban/internal/pipeline"
)

// init 注册管理器
func init() {
	manager.Register[*pipeline.Manager]()
}

// init 注册管道
func init() {
	pipeline.Register(pipeline2.NewDefaultPipeline)
}
