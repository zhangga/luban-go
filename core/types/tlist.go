package types

type TList struct {
	*TTypeBase
	ElementType_ TType
}

func NewTList(isNullable bool, elementType TType, tags map[string]string) *TList {
	return &TList{
		TTypeBase:    NewTTypeBase(isNullable, tags),
		ElementType_: elementType,
	}
}

func (t *TList) TypeName() string {
	return "list"
}

func (t *TList) IsCollection() bool {
	return true
}

func (t *TList) ElementType() TType {
	return t.ElementType_
}

func (t *TList) TryParseFrom(s string) bool {
	return false
}

func (t *TList) Accept(visitor ITypeVisitor) interface{} {
	return visitor.VisitTList(t)
}
