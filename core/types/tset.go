package types

type TSet struct {
	*TTypeBase
	ElementType_ TType
}

func NewTSet(isNullable bool, elementType TType, tags map[string]string) *TSet {
	return &TSet{
		TTypeBase:    NewTTypeBase(isNullable, tags),
		ElementType_: elementType,
	}
}

func (t *TSet) TypeName() string {
	return "set"
}

func (t *TSet) IsCollection() bool {
	return true
}

func (t *TSet) ElementType() TType {
	return t.ElementType_
}

func (t *TSet) TryParseFrom(s string) bool {
	return false
}

func (t *TSet) Accept(visitor ITypeVisitor) interface{} {
	return visitor.VisitTSet(t)
}
