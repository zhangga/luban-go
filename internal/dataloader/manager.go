package dataloader

import (
	"github.com/zhangga/luban/core/manager"
	"github.com/zhangga/luban/pkg/logger"
)

var _ manager.IDataLoaderManager = (*Manager)(nil)

type Manager struct {
	manager.EmbedDataLoaderManager
	logger logger.Logger
}

func (m *Manager) Init(logger logger.Logger) {
	m.logger = logger
}

func (m *Manager) PostInit() {
}
