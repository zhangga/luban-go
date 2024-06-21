package pipeline

import "github.com/zhangga/luban/core/pipeline"

var _ pipeline.IPipeline = (*DefaultPipeline)(nil)

type DefaultPipeline struct {
}

func NewDefaultPipeline() pipeline.IPipeline {
	return &DefaultPipeline{}
}

func (p *DefaultPipeline) Name() string {
	return "default"
}

func (p *DefaultPipeline) Run(args pipeline.Arguments) error {
	//TODO implement me
	panic("implement me")
}
