package pipeline

import (
	"fmt"
	"github.com/zhangga/luban/core/manager"
	"github.com/zhangga/luban/pkg/logger"
)

var _ manager.IManager = (*Manager)(nil)

// Manager is the manager of the pipeline
type Manager struct {
	logger logger.Logger
}

func (m *Manager) Init(logger logger.Logger) {
	m.logger = logger
}

func (m *Manager) PostInit() {
}

func (m *Manager) CreatePipeline(name string) IPipeline {
	creator := getCreator(name)
	if creator == nil {
		panic(fmt.Errorf("pipeline %s not found", name))
	}
	return creator(m.logger)
}
