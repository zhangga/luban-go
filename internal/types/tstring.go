package types

import "github.com/zhangga/luban/core/refs"

var _ refs.TType = (*TString)(nil)

type TString struct {
	refs.EmbedTType
	refs.EmbedTString
}

func NewTString(isNullable bool, tags map[string]string, _ refs.IDefType, _ ...refs.TType) refs.TType {
	return &TString{
		EmbedTType: refs.NewEmbedTType(isNullable, tags),
	}
}

func (t *TString) TypeName() string {
	return "string"
}

func (t *TString) IsBean() bool {
	return false
}

func (t *TString) IsNullable() bool {
	return t.EmbedTType.IsNullable
}

func (t *TString) TryParseFrom(s string) bool {
	return true
}
