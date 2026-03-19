package types

type DefBean interface {
	FullName() string
	Name() string
	Namespace() string
	HierarchyNotAbstractChildren() []interface{}
}

type TBean struct {
	*TTypeBase
	DefBean DefBean
}

func NewTBean(isNullable bool, defBean DefBean, tags map[string]string) *TBean {
	return &TBean{
		TTypeBase: NewTTypeBase(isNullable, tags),
		DefBean:   defBean,
	}
}

func (t *TBean) TypeName() string {
	return "bean"
}

func (t *TBean) IsBean() bool {
	return true
}

// IsDynamic 是否为多态基类（即是否存在非抽象子类且包含自己与否）
func (t *TBean) IsDynamic() bool {
	if t.DefBean == nil {
		return false
	}
	children := t.DefBean.HierarchyNotAbstractChildren()
	return len(children) > 1
}

func (t *TBean) TryParseFrom(s string) bool {
	return false
}

func (t *TBean) Accept(visitor ITypeVisitor) interface{} {
	return visitor.VisitTBean(t)
}
