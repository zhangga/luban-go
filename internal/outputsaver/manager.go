package outputsaver

import (
	"github.com/zhangga/luban/core/manager"
	"github.com/zhangga/luban/pkg/logger"
)

var _ manager.IOutputSaverManager = (*Manager)(nil)

type Manager struct {
	manager.EmbedOutputSaverManager
	logger logger.Logger
}

func (m *Manager) Init(logger logger.Logger) {
	m.logger = logger
}

func (m *Manager) PostInit() {
}
