package types

// TText 对应 luban 中的 text 类型，用于多语言本地化支持
type TText struct {
	*TTypeBase
}

func NewTText(isNullable bool, tags map[string]string) *TText {
	return &TText{
		TTypeBase: NewTTypeBase(isNullable, tags),
	}
}

func (t *TText) TypeName() string {
	return "text"
}

func (t *TText) TryParseFrom(s string) bool {
	return true
}

func (t *TText) Accept(visitor ITypeVisitor) interface{} {
	return visitor.VisitTText(t)
}
