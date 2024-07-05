package types

import (
	"github.com/zhangga/luban/core/refs"
	"strconv"
)

var _ refs.TType = (*TFloat)(nil)

type TFloat struct {
	refs.EmbedTType
	refs.EmbedTFloat
}

func NewTFloat(isNullable bool, tags map[string]string, _ refs.IDefType, _ ...refs.TType) refs.TType {
	return &TFloat{
		EmbedTType: refs.NewEmbedTType(isNullable, tags),
	}
}

func (f *TFloat) TypeName() string {
	return "float"
}

func (f *TFloat) IsBean() bool {
	return false
}

func (f *TFloat) IsNullable() bool {
	return f.EmbedTType.IsNullable
}

func (f *TFloat) TryParseFrom(s string) bool {
	f32, err := strconv.ParseFloat(s, 32)
	_ = f32
	return err == nil
}
