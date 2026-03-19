package types

type DefEnum interface {
	FullName() string
	Name() string
	Namespace() string
}

type TEnum struct {
	*TTypeBase
	DefEnum DefEnum
}

func NewTEnum(isNullable bool, defEnum DefEnum, tags map[string]string) *TEnum {
	// 在原C#代码中合并了 defEnum 的 tags，这里暂时只用传入的 tags
	return &TEnum{
		TTypeBase: NewTTypeBase(isNullable, tags),
		DefEnum:   defEnum,
	}
}

func (t *TEnum) TypeName() string {
	return "enum"
}

func (t *TEnum) IsEnum() bool {
	return true
}

func (t *TEnum) TryParseFrom(s string) bool {
	// 枚举类型的解析通常需要依赖 defEnum 来判断字符串或数字是否有效
	return false
}

func (t *TEnum) Accept(visitor ITypeVisitor) interface{} {
	return visitor.VisitTEnum(t)
}
