package types

import (
	"github.com/zhangga/luban/core/refs"
	"strconv"
)

var _ refs.TType = (*TShort)(nil)

type TShort struct {
	refs.EmbedTType
	refs.EmbedTShort
}

func NewTShort(isNullable bool, tags map[string]string, _ refs.IDefType, _ ...refs.TType) refs.TType {
	return &TShort{
		EmbedTType: refs.NewEmbedTType(isNullable, tags),
	}
}

func (t *TShort) TypeName() string {
	return "short"
}

func (t *TShort) IsBean() bool {
	return false
}

func (t *TShort) IsNullable() bool {
	return t.EmbedTType.IsNullable
}

func (t *TShort) TryParseFrom(s string) bool {
	_, err := strconv.ParseInt(s, 10, 16)
	return err == nil
}
