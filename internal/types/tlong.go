package types

import (
	"github.com/zhangga/luban/core/refs"
	"strconv"
)

var _ refs.TType = (*TLong)(nil)

type TLong struct {
	refs.EmbedTType
	refs.EmbedTLong
}

func NewTLong(isNullable bool, tags map[string]string, _ refs.IDefType, _ ...refs.TType) refs.TType {
	return &TLong{
		EmbedTType: refs.NewEmbedTType(isNullable, tags),
	}
}

func (l *TLong) TypeName() string {
	return "long"
}

func (l *TLong) IsBean() bool {
	return false
}

func (l *TLong) IsNullable() bool {
	return l.EmbedTType.IsNullable
}

func (l *TLong) TryParseFrom(s string) bool {
	_, err := strconv.ParseInt(s, 10, 64)
	return err == nil
}
