package refs

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
	PreCompile()
	Compile()
	PostCompile()
}

type IDefField interface {
	Name() string
	SetAutoId(id int)
	Compile()
	PostCompile()
}
