package types

type TMap struct {
	*TTypeBase
	KeyType_   TType
	ValueType_ TType
}

func NewTMap(isNullable bool, keyType, valueType TType, tags map[string]string) *TMap {
	return &TMap{
		TTypeBase:  NewTTypeBase(isNullable, tags),
		KeyType_:   keyType,
		ValueType_: valueType,
	}
}

func (t *TMap) TypeName() string {
	return "map"
}

func (t *TMap) IsCollection() bool {
	return true
}

func (t *TMap) KeyType() TType {
	return t.KeyType_
}

func (t *TMap) ValueType() TType {
	return t.ValueType_
}

func (t *TMap) TryParseFrom(s string) bool {
	return false
}

func (t *TMap) Accept(visitor ITypeVisitor) interface{} {
	return visitor.VisitTMap(t)
}
