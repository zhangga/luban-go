package manager

import (
	"github.com/modern-go/reflect2"
	"github.com/zhangga/luban/core/pipeline"
)

type IPipelineManager interface {
	IManager
	MustEmbedPipelineManager

	CreatePipeline(name string) pipeline.IPipeline
}

type MustEmbedPipelineManager interface {
	mustEmbedPipelineManager()
}

type EmbedPipelineManager struct {
}

func (EmbedPipelineManager) mustEmbedPipelineManager() {}

func (EmbedPipelineManager) key() uintptr {
	rtype := reflect2.RTypeOf((*IPipelineManager)(nil))
	return rtype
}
