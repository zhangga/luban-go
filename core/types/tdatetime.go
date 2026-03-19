package types

type TDateTime struct {
	*TTypeBase
}

func NewTDateTime(isNullable bool, tags map[string]string) *TDateTime {
	return &TDateTime{
		TTypeBase: NewTTypeBase(isNullable, tags),
	}
}

func (t *TDateTime) TypeName() string {
	return "datetime"
}

func (t *TDateTime) TryParseFrom(s string) bool {
	// 这里简单的 DateTime 解析留作后期完善（可以借助外部库或 time.Parse）
	// 目前仅返回 true 作为占位
	return true
}

func (t *TDateTime) Accept(visitor ITypeVisitor) interface{} {
	return visitor.VisitTDateTime(t)
}
