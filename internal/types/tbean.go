package types

import (
	"github.com/zhangga/luban/core/refs"
	"github.com/zhangga/luban/internal/defs"
)

var _ refs.TType = (*TBean)(nil)

type TBean struct {
	refs.EmbedTType
	DefBean *defs.DefBean
}

func NewTBean(defBean *defs.DefBean, isNullable bool, tags map[string]string) *TBean {
	return &TBean{
		EmbedTType: refs.NewEmbedTType(isNullable, tags),
		DefBean:    defBean,
	}
}

func (b *TBean) TypeName() string {
	return "bean"
}

func (b *TBean) IsBean() bool {
	return true
}

func (b *TBean) IsNullable() bool {
	//TODO implement me
	panic("implement me")
}

func (b *TBean) TryParseFrom(s string) bool {
	panic("not supported")
}
