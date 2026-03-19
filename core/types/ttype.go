package types

// TType 定义了 Luban 类型系统的基本接口
type TType interface {
	IsNullable() bool
	HasTag(attrName string) bool
	GetTag(attrName string) (string, bool)
	GetTagOrDefault(attrName, defaultValue string) string
	TypeName() string

	IsCollection() bool
	IsBean() bool
	IsEnum() bool
	ElementType() TType

	TryParseFrom(s string) bool

	// Accept 是访问者模式的入口
	Accept(visitor ITypeVisitor) interface{}
}

// TTypeBase 提供了 TType 接口的基础实现
type TTypeBase struct {
	isNullable bool
	tags       map[string]string
}

func NewTTypeBase(isNullable bool, tags map[string]string) *TTypeBase {
	if tags == nil {
		tags = make(map[string]string)
	}
	return &TTypeBase{
		isNullable: isNullable,
		tags:       tags,
	}
}

func (t *TTypeBase) IsNullable() bool {
	return t.isNullable
}

func (t *TTypeBase) Tags() map[string]string {
	return t.tags
}

func (t *TTypeBase) HasTag(attrName string) bool {
	if t.tags == nil {
		return false
	}
	_, ok := t.tags[attrName]
	return ok
}

func (t *TTypeBase) GetTag(attrName string) (string, bool) {
	if t.tags == nil {
		return "", false
	}
	val, ok := t.tags[attrName]
	return val, ok
}

func (t *TTypeBase) GetTagOrDefault(attrName, defaultValue string) string {
	if t.tags == nil {
		return defaultValue
	}
	if val, ok := t.tags[attrName]; ok {
		return val
	}
	return defaultValue
}

func (t *TTypeBase) IsCollection() bool {
	return false
}

func (t *TTypeBase) IsBean() bool {
	return false
}

func (t *TTypeBase) IsEnum() bool {
	return false
}

func (t *TTypeBase) ElementType() TType {
	return nil
}
