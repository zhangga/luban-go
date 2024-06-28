package manager

import "github.com/modern-go/reflect2"

type IDataTargetManager interface {
	IManager
	MustEmbedDataTargetManager
}

type MustEmbedDataTargetManager interface {
	mustEmbedDataTargetManager()
}

type EmbedDataTargetManager struct {
}

func (EmbedDataTargetManager) mustEmbedDataTargetManager() {}

func (EmbedDataTargetManager) key() uintptr {
	rtype := reflect2.RTypeOf((*IDataTargetManager)(nil))
	return rtype
}
