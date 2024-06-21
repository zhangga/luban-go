package schema

import "github.com/zhangga/luban/pkg/logger"

type Manager struct {
	logger logger.Logger
}

func NewManager(logger logger.Logger) *Manager {
	return &Manager{
		logger: logger,
	}
}

func (m *Manager) Init() {

}
