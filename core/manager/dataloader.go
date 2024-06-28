package manager

import "github.com/modern-go/reflect2"

type IDataLoaderManager interface {
	IManager
	MustEmbedDataLoaderManager
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
