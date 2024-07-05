package types

import "github.com/zhangga/luban/core/refs"

var _ refs.TType = (*TDateTime)(nil)

type TDateTime struct {
	refs.EmbedTType
	refs.EmbedTDateTime
}

func NewTDateTime(isNullable bool, tags map[string]string, _ refs.IDefType, _ ...refs.TType) refs.TType {
	return &TDateTime{
		EmbedTType: refs.NewEmbedTType(isNullable, tags),
	}
}

func (t *TDateTime) TypeName() string {
	return "datetime"
}

func (t *TDateTime) IsBean() bool {
	return false
}

func (t *TDateTime) IsNullable() bool {
	return t.EmbedTType.IsNullable
}

func (t *TDateTime) TryParseFrom(s string) bool {
	panic("not supported")
}
