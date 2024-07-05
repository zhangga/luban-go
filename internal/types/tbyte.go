package types

import (
	"github.com/zhangga/luban/core/refs"
	"strconv"
)

var _ refs.TType = (*TByte)(nil)

type TByte struct {
	refs.EmbedTType
	refs.EmbedTByte
}

func NewTByte(isNullable bool, tags map[string]string, _ refs.IDefType, _ ...refs.TType) refs.TType {
	return &TByte{
		EmbedTType: refs.NewEmbedTType(isNullable, tags),
	}
}

func (b *TByte) TypeName() string {
	return "byte"
}

func (b *TByte) IsBean() bool {
	return false
}

func (b *TByte) IsNullable() bool {
	return b.EmbedTType.IsNullable
}

func (b *TByte) TryParseFrom(s string) bool {
	_, err := strconv.ParseInt(s, 10, 8)
	return err == nil
}
