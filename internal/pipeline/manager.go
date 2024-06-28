package pipeline

import (
	"github.com/zhangga/luban/core/manager"
	"github.com/zhangga/luban/core/pipeline"
	"github.com/zhangga/luban/pkg/logger"
)

var _ manager.IPipelineManager = (*Manager)(nil)

type Manager struct {
	manager.EmbedPipelineManager
	logger logger.Logger
}

func (m *Manager) Init(logger logger.Logger) {
	m.logger = logger
}

func (m *Manager) PostInit() {
}

func (m *Manager) CreatePipeline(name string) pipeline.IPipeline {
	creator := pipeline.GetCreator(name)
	if creator == nil {
		panic("pipeline " + name + " not found")
	}
	return creator(m.logger)
}
