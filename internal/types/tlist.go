package types

import "github.com/zhangga/luban/core/refs"

var _ refs.TType = (*TList)(nil)

type TList struct {
	refs.EmbedTType
	refs.EmbedTList
}

func NewTList(isNullable bool, tags map[string]string, _ refs.IDefType, typ ...refs.TType) refs.TType {
	et := refs.NewEmbedTType(isNullable, tags)
	et.IsCollection = true
	return &TList{
		EmbedTType: et,
	}
}

func (l *TList) TypeName() string {
	return "list"
}

func (l *TList) IsBean() bool {
	return false
}

func (l *TList) IsNullable() bool {
	return l.EmbedTType.IsNullable
}

func (l *TList) TryParseFrom(s string) bool {
	panic("not supported")
}
