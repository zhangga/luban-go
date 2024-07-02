package manager

import (
	"github.com/modern-go/reflect2"
	"github.com/zhangga/luban/core/refs"
)

type IDataLoaderManager interface {
	IManager
	MustEmbedDataLoaderManager
	LoadTableFile(valueType refs.TType, file, subAssetName string, options map[string]string) ([]*refs.Record, error)
}

type MustEmbedDataLoaderManager interface {
	mustEmbedDataLoaderManager()
}

type EmbedDataLoaderManager struct {
}

func (EmbedDataLoaderManager) mustEmbedDataLoaderManager() {}

func (EmbedDataLoaderManager) key() uintptr {
	rtype := reflect2.RTypeOf((*IDataLoaderManager)(nil))
	return rtype
}
