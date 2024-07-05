package types

import "github.com/zhangga/luban/core/refs"

var _ refs.TType = (*TSet)(nil)

type TSet struct {
	refs.EmbedTType
	refs.EmbedTSet
}

func NewTSet(isNullable bool, tags map[string]string, _ refs.IDefType, typ ...refs.TType) refs.TType {
	et := refs.NewEmbedTType(isNullable, tags)
	et.IsCollection = true
	return &TSet{
		EmbedTType: et,
	}
}

func (t *TSet) TypeName() string {
	return "set"
}

func (t *TSet) IsBean() bool {
	return false
}

func (t *TSet) IsNullable() bool {
	return t.EmbedTType.IsNullable
}

func (t *TSet) TryParseFrom(s string) bool {
	panic("not supported")
}
