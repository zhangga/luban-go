package defs

import "github.com/zhangga/luban-go/core/rawdefs"

type DefBeanImpl struct {
	Raw           *rawdefs.RawBean
	Fields        []*DefField
	Parent        string
	ParentDefType *DefBeanImpl
	Children      []*DefBeanImpl

	HierarchyFields []*DefField
	IsAbstractType  bool
}

func NewDefBeanImpl(raw *rawdefs.RawBean) *DefBeanImpl {
	b := &DefBeanImpl{
		Raw:             raw,
		Fields:          make([]*DefField, 0),
		Children:        make([]*DefBeanImpl, 0),
		HierarchyFields: make([]*DefField, 0),
	}
	if raw != nil {
		b.Parent = raw.Parent
		if raw.Fields != nil {
			for _, f := range raw.Fields {
				b.Fields = append(b.Fields, NewDefField(f))
			}
		}
	}
	return b
}

func (b *DefBeanImpl) FullName() string {
	return b.Raw.FullName()
}

func (b *DefBeanImpl) Name() string {
	return b.Raw.Name
}

func (b *DefBeanImpl) Namespace() string {
	return b.Raw.Namespace
}

// HierarchyNotAbstractChildren 收集所有非抽象的子类
func (b *DefBeanImpl) HierarchyNotAbstractChildren() []interface{} {
	var result []interface{}
	if !b.IsAbstractType {
		result = append(result, b)
	}
	for _, c := range b.Children {
		result = append(result, c.HierarchyNotAbstractChildren()...)
	}
	return result
}

// TryGetChild 尝试根据子类型名或别名获取对应的子类定义
func (b *DefBeanImpl) TryGetChild(typeStr string) *DefBeanImpl {
	if b.Name() == typeStr || b.FullName() == typeStr || b.Raw.Alias == typeStr {
		return b
	}
	for _, c := range b.Children {
		if res := c.TryGetChild(typeStr); res != nil {
			return res
		}
	}
	return nil
}

type DefField struct {
	Raw *rawdefs.RawField
}

func NewDefField(raw *rawdefs.RawField) *DefField {
	return &DefField{Raw: raw}
}
