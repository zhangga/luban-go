package types

type TArray struct {
	*TTypeBase
	ElementType_ TType
}

func NewTArray(isNullable bool, elementType TType, tags map[string]string) *TArray {
	return &TArray{
		TTypeBase:    NewTTypeBase(isNullable, tags),
		ElementType_: elementType,
	}
}

func (t *TArray) TypeName() string {
	return "array"
}

func (t *TArray) IsCollection() bool {
	return true
}

func (t *TArray) ElementType() TType {
	return t.ElementType_
}

func (t *TArray) TryParseFrom(s string) bool {
	return false
}

func (t *TArray) Accept(visitor ITypeVisitor) interface{} {
	return visitor.VisitTArray(t)
}
