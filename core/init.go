package core

import (
	"github.com/zhangga/luban/core/codeformat"
	"github.com/zhangga/luban/core/codetarget"
	"github.com/zhangga/luban/core/dataloader"
	"github.com/zhangga/luban/core/datatarget"
	"github.com/zhangga/luban/core/l10n"
	"github.com/zhangga/luban/core/manager"
	"github.com/zhangga/luban/core/outputsaver"
	"github.com/zhangga/luban/core/pipeline"
	"github.com/zhangga/luban/core/postprocess"
	"github.com/zhangga/luban/core/schema"
	"github.com/zhangga/luban/core/template"
	"github.com/zhangga/luban/core/validator"
	pipeline2 "github.com/zhangga/luban/internal/pipeline"
)

// init 注册管理器
func init() {
	manager.Register[*schema.Manager]()
	manager.Register[*template.Manager]()
	manager.Register[*codeformat.Manager]()
	manager.Register[*codetarget.Manager]()
	manager.Register[*postprocess.Manager]()
	manager.Register[*outputsaver.Manager]()
	manager.Register[*dataloader.Manager]()
	manager.Register[*validator.Manager]()
	manager.Register[*datatarget.Manager]()
	manager.Register[*pipeline.Manager]()
	manager.Register[*l10n.Manager]()
}

// init 注册管道
func init() {
	pipeline.Register(pipeline2.NewDefaultPipeline)
}
