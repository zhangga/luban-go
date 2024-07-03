package refs

import (
	"github.com/zhangga/luban/core/pctx"
)

type IDefAssembly interface {
	mustEmbedDefAssembly()
}

type EmbedDefAssembly struct {
}

func (EmbedDefAssembly) mustEmbedDefAssembly() {}

type IDefType interface {
	Namespace() string
	FullName() string
	SetAssembly(assembly IDefAssembly)
	PreCompile(ctx pctx.Context)
	Compile(ctx pctx.Context)
	PostCompile(ctx pctx.Context)
}

type IDefField interface {
	Name() string
	SetAutoId(id int)
	Compile(ctx pctx.Context)
	PostCompile(ctx pctx.Context)
}
