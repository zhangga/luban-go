package dataloader

import (
	"github.com/zhangga/luban/core/manager"
	"github.com/zhangga/luban/core/refs"
	"github.com/zhangga/luban/pkg/logger"
	"path/filepath"
)

const loaderKey = "loader"

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

func (m *Manager) LoadTableFile(valueType refs.TType, file, subAssetName string, options map[string]string) ([]*refs.Record, error) {
	loaderName, ok := options[loaderKey]
	if !ok {
		loaderName = filepath.Ext(file)
	}

}
