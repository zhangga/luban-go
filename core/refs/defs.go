package refs

import "github.com/zhangga/luban/core/pipeline"

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
	PreCompile(pipeline pipeline.IPipeline)
	Compile(pipeline pipeline.IPipeline)
	PostCompile(pipeline pipeline.IPipeline)
}

type IDefField interface {
	Name() string
	SetAutoId(id int)
	Compile(pipeline pipeline.IPipeline)
	PostCompile(pipeline pipeline.IPipeline)
}
