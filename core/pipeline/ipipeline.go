package pipeline

import (
	"github.com/zhangga/luban-go/core/codetarget"
	"github.com/zhangga/luban-go/core/defs"
)

type PipelineArguments struct {
	Target        string
	CodeTargets   []string
	DataTargets   []string
	InputDataDir  string
	OutputCodeDir string
	OutputDataDir string
	TemplateDir   string
	L10NProvider  string
	Timezone      string
	// ...可以加入更多参数
}

type IPipeline interface {
	Process(args *PipelineArguments) error
}

// 模拟上下文结构，供整个管线生命周期传递状态
type GenerationContext struct {
	Args     *PipelineArguments
	Assembly *defs.DefAssemblyImpl
	Manifest *codetarget.OutputFileManifest
}

func NewGenerationContext(args *PipelineArguments, assembly *defs.DefAssemblyImpl) *GenerationContext {
	return &GenerationContext{
		Args:     args,
		Assembly: assembly,
		Manifest: codetarget.NewOutputFileManifest(),
	}
}
