package manager

import "github.com/modern-go/reflect2"

type ITemplateManager interface {
	IManager
	MustEmbedTemplateManager
}

type MustEmbedTemplateManager interface {
	mustEmbedTemplateManager()
}

type EmbedTemplateManager struct {
}

func (EmbedTemplateManager) mustEmbedTemplateManager() {}

func (EmbedTemplateManager) key() uintptr {
	rtype := reflect2.RTypeOf((*ITemplateManager)(nil))
	return rtype
}
