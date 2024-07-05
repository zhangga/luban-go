package types

import (
	"github.com/zhangga/luban/core/refs"
	"strings"
)

var _ refs.TType = (*TBool)(nil)

type TBool struct {
	refs.EmbedTType
	refs.EmbedTBool
}

func NewTBool(isNullable bool, tags map[string]string, _ refs.IDefType, _ ...refs.TType) refs.TType {
	return &TBool{
		EmbedTType: refs.NewEmbedTType(isNullable, tags),
	}
}

func (b *TBool) TypeName() string {
	return "bool"
}

func (b *TBool) IsBean() bool {
	return false
}

func (b *TBool) IsNullable() bool {
	return b.EmbedTType.IsNullable
}

func (b *TBool) TryParseFrom(s string) bool {
	s = strings.ToLower(s)
	return s == "true" || s == "false" || s == "1" || s == "0"
}
