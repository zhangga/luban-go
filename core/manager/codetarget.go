package manager

import "github.com/modern-go/reflect2"

type ICodeTargetManager interface {
	IManager
	MustEmbedCodeTargetManager
}

type MustEmbedCodeTargetManager interface {
	mustEmbedCodeTargetManager()
}

type EmbedCodeTargetManager struct {
}

func (EmbedCodeTargetManager) mustEmbedCodeTargetManager() {}

func (EmbedCodeTargetManager) key() uintptr {
	rtype := reflect2.RTypeOf((*ICodeTargetManager)(nil))
	return rtype
}
