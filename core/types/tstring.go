package types

type TString struct {
	*TTypeBase
}

func NewTString(isNullable bool, tags map[string]string) *TString {
	return &TString{
		TTypeBase: NewTTypeBase(isNullable, tags),
	}
}

func (t *TString) TypeName() string {
	return "string"
}

func (t *TString) TryParseFrom(s string) bool {
	return true
}

func (t *TString) Accept(visitor ITypeVisitor) interface{} {
	return visitor.VisitTString(t)
}
