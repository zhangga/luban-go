package types

import (
	"github.com/zhangga/luban/core/refs"
	"github.com/zhangga/luban/internal/defs"
)

var _ refs.TType = (*TEnum)(nil)

type TEnum struct {
	refs.EmbedTType
	refs.EmbedTEnum
	DefEnum *defs.DefEnum
}

func NewTEnum(isNullable bool, tags map[string]string, def refs.IDefType, _ ...refs.TType) refs.TType {
	enum := &TEnum{
		EmbedTType: refs.NewEmbedTType(isNullable, tags),
	}
	if def != nil {
		enum.DefEnum = def.(*defs.DefEnum)
	}
	return enum
}

func (e *TEnum) TypeName() string {
	return "enum"
}

func (e *TEnum) IsBean() bool {
	return false
}

func (e *TEnum) IsNullable() bool {
	return e.EmbedTType.IsNullable
}

func (e *TEnum) TryParseFrom(s string) bool {
	panic("not supported")
}
