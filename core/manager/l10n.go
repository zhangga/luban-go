package manager

import "github.com/modern-go/reflect2"

type IL10nManager interface {
	IManager
	MustEmbedL10nManager
}

type MustEmbedL10nManager interface {
	mustEmbedL10nManager()
}

type EmbedL10nManager struct {
}

func (EmbedL10nManager) mustEmbedL10nManager() {}

func (EmbedL10nManager) key() uintptr {
	rtype := reflect2.RTypeOf((*IL10nManager)(nil))
	return rtype
}
