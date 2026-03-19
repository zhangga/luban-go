package types

import "strconv"

type TBool struct {
	*TTypeBase
}

func NewTBool(isNullable bool, tags map[string]string) *TBool {
	return &TBool{
		TTypeBase: NewTTypeBase(isNullable, tags),
	}
}

func (t *TBool) TypeName() string {
	return "bool"
}

func (t *TBool) TryParseFrom(s string) bool {
	_, err := strconv.ParseBool(s)
	return err == nil
}

func (t *TBool) Accept(visitor ITypeVisitor) interface{} {
	return visitor.VisitTBool(t)
}
