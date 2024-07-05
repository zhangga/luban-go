package types

import (
	"github.com/zhangga/luban/core/refs"
	"github.com/zhangga/luban/internal/defs"
)

var _ refs.TType = (*TBean)(nil)

type TBean struct {
	refs.EmbedTType
	refs.EmbedTBean
	DefBean *defs.DefBean
}

func NewTBean(isNullable bool, tags map[string]string, def refs.IDefType, typ ...refs.TType) refs.TType {
	bean := &TBean{
		EmbedTType: refs.NewEmbedTType(isNullable, tags),
	}
	if def != nil {
		bean.DefBean = def.(*defs.DefBean)
	}
	return bean
}

func (b *TBean) TypeName() string {
	return "bean"
}

func (b *TBean) IsBean() bool {
	return true
}

func (b *TBean) IsNullable() bool {
	return b.EmbedTType.IsNullable
}

func (b *TBean) TryParseFrom(s string) bool {
	panic("not supported")
}
