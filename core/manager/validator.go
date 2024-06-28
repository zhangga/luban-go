package manager

import "github.com/modern-go/reflect2"

type IValidatorManager interface {
	IManager
	MustEmbedValidatorManager
}

type MustEmbedValidatorManager interface {
	mustEmbedValidatorManager()
}

type EmbedValidatorManager struct {
}

func (EmbedValidatorManager) mustEmbedValidatorManager() {}

func (EmbedValidatorManager) key() uintptr {
	rtype := reflect2.RTypeOf((*IValidatorManager)(nil))
	return rtype
}
