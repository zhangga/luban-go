package manager

import "github.com/modern-go/reflect2"

type IPostProcessManager interface {
	IManager
	MustEmbedPostProcessManager
}

type MustEmbedPostProcessManager interface {
	mustEmbedPostProcessManager()
}

type EmbedPostProcessManager struct {
}

func (EmbedPostProcessManager) mustEmbedPostProcessManager() {}

func (EmbedPostProcessManager) key() uintptr {
	rtype := reflect2.RTypeOf((*IPostProcessManager)(nil))
	return rtype
}
