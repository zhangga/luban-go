package defs

import (
	"fmt"
	"github.com/zhangga/luban/core/pipeline"
	"github.com/zhangga/luban/core/refs"
	"github.com/zhangga/luban/internal/rawdefs"
	"github.com/zhangga/luban/internal/utils"
)

var _ refs.IDefType = (*DefBean)(nil)

// DefBean 对象定义
type DefBean struct {
	DefTypeBase
	Id                           int64
	AutoId                       int // for protobuf
	Alias                        string
	Parent                       string
	ParentDefType                *DefBean
	IsMultiRow                   bool
	Sep                          string
	IsValueType                  bool
	Children                     []*DefBean
	HierarchyNotAbstractChildren []*DefBean
	HierarchyFields              []*DefField
	Fields                       []*DefField
	hierarchyExportFields        []*DefField
	exportFields                 []*DefField
}

func NewDefBean(rawBean rawdefs.RawBean) *DefBean {
	bean := &DefBean{
		DefTypeBase: DefTypeBase{
			Name:        rawBean.Name,
			namespace:   rawBean.Namespace,
			Comment:     rawBean.Comment,
			Tags:        rawBean.Tags,
			Groups:      rawBean.Groups,
			TypeMappers: rawBean.TypeMappers,
		},
		Parent:      rawBean.Parent,
		Id:          utils.ComputeCfgHashIdByName(rawBean.FullName()),
		Alias:       rawBean.Alias,
		Sep:         rawBean.Sep,
		IsValueType: rawBean.IsValueType,
	}
	for _, field := range rawBean.Fields {
		bean.Fields = append(bean.Fields, bean.CreateField(field, 0))
	}
	return bean
}

func (b *DefBean) CreateField(field *rawdefs.RawField, idOffset int) *DefField {
	return NewDefField(b, field, idOffset)
}

func (b *DefBean) GetField(name string) *DefField {
	for _, field := range b.HierarchyFields {
		if field.Name() == name {
			return field
		}
	}
	return nil
}

func (b *DefBean) TryGetField(name string) (*DefField, int) {
	for i, field := range b.HierarchyFields {
		if field.Name() == name {
			return field, i
		}
	}
	return nil, -1
}

func (b *DefBean) RootDefType() *DefBean {
	if b.ParentDefType == nil {
		return b
	}
	return b.ParentDefType.RootDefType()
}

func (b *DefBean) IsAbstractType() bool {
	return len(b.Children) > 0
}

func (b *DefBean) IsAssignableFrom(parent *DefBean) bool {
	for {
		if parent == nil {
			return false
		}
		if parent == b {
			return true
		}
		parent = parent.ParentDefType
	}
}

func (b *DefBean) GetHierarchyChildren() []*DefBean {
	var children []*DefBean
	// 将自身添加到列表
	children = append(children, b)
	// 如果有子类，将子类添加到列表
	for _, child := range b.Children {
		children = append(children, child.GetHierarchyChildren()...)
	}
	return children
}

func (b *DefBean) HierarchyExportFields() []*DefField {
	var fields []*DefField
	if len(b.hierarchyExportFields) > 0 {
		for _, field := range b.HierarchyFields {
			if field.NeedExport() {
				fields = append(fields, field)
			}
		}
	}
	return fields
}

func (b *DefBean) ExportFields() []*DefField {
	var fields []*DefField
	if len(b.exportFields) > 0 {
		for _, field := range b.Fields {
			if field.NeedExport() {
				fields = append(fields, field)
			}
		}
	}
	return fields
}

func (b *DefBean) CollectHierarchyNotAbstractChildren(children []*DefBean) {
	if b.IsAbstractType() {
		for _, child := range b.Children {
			child.CollectHierarchyNotAbstractChildren(children)
		}
	} else {
		children = append(children, b)
	}
}

func (b *DefBean) CollectHierarchyFields(fields []*DefField) {
	if b.ParentDefType != nil {
		b.ParentDefType.CollectHierarchyFields(fields)
	}
	fields = append(fields, b.Fields...)
}

func (b *DefBean) SetupParentRecursively() {
	if b.ParentDefType != nil || len(b.Parent) == 0 {
		return
	}

	parent := b.Assembly.GetDefType(b.namespace, b.Parent)
	if parent == nil {
		panic(fmt.Errorf("bean: %s, parent: %s not found", b.FullName(), b.Parent))
	}
	parentType, ok := parent.(*DefBean)
	if !ok {
		panic(fmt.Errorf("bean: %s, parent: %s is not a bean", b.FullName(), b.Parent))
	}
	b.ParentDefType = parentType
	b.ParentDefType.Children = append(b.ParentDefType.Children, b)
	b.ParentDefType.SetupParentRecursively()
}

func (b *DefBean) PreCompile(pipeline pipeline.IPipeline) {
	b.DefTypeBase.PreCompile(pipeline)
	b.SetupParentRecursively()
	b.CollectHierarchyFields(b.HierarchyFields)
}

func (b *DefBean) Compile(pipeline pipeline.IPipeline) {
	var cs []*DefBean
	if len(b.Children) > 0 {
		b.CollectHierarchyNotAbstractChildren(cs)
	}
	b.HierarchyNotAbstractChildren = cs
	// 检查别名是否重复
	nameOrAliasNames := map[string]struct{}{}
	for _, child := range b.HierarchyNotAbstractChildren {
		nameOrAliasNames[child.Name] = struct{}{}
	}
	for _, child := range b.HierarchyNotAbstractChildren {
		if len(child.Alias) > 0 {
			if _, ok := nameOrAliasNames[child.Alias]; ok {
				panic(fmt.Errorf("bean: %s, alias: %s 重复", child.FullName(), child.Alias))
			}
			nameOrAliasNames[child.Alias] = struct{}{}
		}
	}
	CompileFields(pipeline, &b.DefTypeBase, b.HierarchyFields)
}

func (b *DefBean) PostCompile(pipeline pipeline.IPipeline) {
	for _, field := range b.HierarchyFields {
		field.PostCompile(pipeline)
	}
	if b.IsAbstractType() && b.ParentDefType == nil {
		autoId := 1
		for _, child := range b.HierarchyNotAbstractChildren {
			child.AutoId = autoId
			autoId++
		}
	}
}
