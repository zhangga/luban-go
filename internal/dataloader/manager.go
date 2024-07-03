package dataloader

import (
	"fmt"
	"github.com/zhangga/luban/core/dataloader"
	"github.com/zhangga/luban/core/manager"
	"github.com/zhangga/luban/core/refs"
	"github.com/zhangga/luban/internal/utils"
	"github.com/zhangga/luban/pkg/logger"
	"strings"
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
	// loader名称
	loaderName, ok := options[loaderKey]
	if !ok {
		loaderName = utils.FileExtWithoutDot(file)
	}
	// 创建loader
	creator := dataloader.GetCreator(loaderName)
	if creator == nil {
		return nil, fmt.Errorf("loader %s not found", loaderName)
	}
	loader := creator(m.logger)
	if err := loader.Load(file, subAssetName); err != nil {
		return nil, err
	}

	if isMultiRecordFile(file, subAssetName) {
		return loader.ReadMulti(valueType), nil
	}
	return []*refs.Record{loader.ReadOne(valueType)}, nil
}

func isMultiRecordFile(file, sheetOrFieldName string) bool {
	return utils.IsExcelFile(file) || isMultiRecordField(sheetOrFieldName)
}

func isMultiRecordField(sheet string) bool {
	return len(sheet) > 0 && strings.HasPrefix(sheet, "*")
}
