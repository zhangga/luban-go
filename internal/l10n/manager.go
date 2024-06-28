package l10n

import (
	"github.com/zhangga/luban/core/manager"
	"github.com/zhangga/luban/pkg/logger"
)

var _ manager.IL10nManager = (*Manager)(nil)

type Manager struct {
	manager.EmbedL10nManager
	logger logger.Logger
}

func (m *Manager) Init(logger logger.Logger) {
	m.logger = logger
}

func (m *Manager) PostInit() {
}
