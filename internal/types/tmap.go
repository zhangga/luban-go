package types

import "github.com/zhangga/luban/core/refs"

var _ refs.TType = (*TMap)(nil)

type TMap struct {
	refs.EmbedTType
	refs.EmbedTMap
	KeyType   refs.TType
	ValueType refs.TType
}

func NewTMap(isNullable bool, tags map[string]string, _ refs.IDefType, typ ...refs.TType) refs.TType {
	et := refs.NewEmbedTType(isNullable, tags)
	et.IsCollection = true
	m := &TMap{
		EmbedTType: et,
	}
	if len(typ) > 1 {
		m.KeyType = typ[0]
		m.ValueType = typ[1]
	}
	return m
}

func (m *TMap) TypeName() string {
	return "map"
}

func (m *TMap) IsBean() bool {
	return false
}

func (m *TMap) IsNullable() bool {
	return m.EmbedTType.IsNullable
}

func (m *TMap) TryParseFrom(s string) bool {
	panic("not supported")
}
