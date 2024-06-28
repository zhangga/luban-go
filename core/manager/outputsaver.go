package manager

import "github.com/modern-go/reflect2"

type IOutputSaverManager interface {
	IManager
	MustEmbedOutputSaverManager
}

type MustEmbedOutputSaverManager interface {
	mustEmbedOutputSaverManager()
}

type EmbedOutputSaverManager struct {
}

func (EmbedOutputSaverManager) mustEmbedOutputSaverManager() {}

func (EmbedOutputSaverManager) key() uintptr {
	rtype := reflect2.RTypeOf((*IOutputSaverManager)(nil))
	return rtype
}
