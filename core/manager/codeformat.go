package manager

import (
	"github.com/modern-go/reflect2"
)

type ICodeFormatManager interface {
	IManager
	MustEmbedCodeFormatManager
}

type MustEmbedCodeFormatManager interface {
	mustEmbedCodeFormatManager()
}

type EmbedCodeFormatManager struct {
}

func (EmbedCodeFormatManager) mustEmbedCodeFormatManager() {}

func (EmbedCodeFormatManager) key() uintptr {
	rtype := reflect2.RTypeOf((*ICodeFormatManager)(nil))
	return rtype
}
