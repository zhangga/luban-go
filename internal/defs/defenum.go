package defs

import (
	"github.com/zhangga/luban/core/refs"
	"github.com/zhangga/luban/internal/rawdefs"
)

var _ refs.IDefType = (*DefEnum)(nil)

type Item struct {
	Name     string
	Value    string
	Alias    string
	IntValue int
	Comment  string
	Tags     map[string]string
}

func (t *Item) AliasOrName() string {
	if len(t.Alias) == 0 {
		return t.Name
	}
	return t.Alias
}

func (t *Item) CommentOrAlias() string {
	if len(t.Comment) == 0 {
		return t.Alias
	}
	return t.Comment
}

func (t *Item) HasTag(attrName string) bool {
	if len(t.Tags) == 0 {
		return false
	}
	_, ok := t.Tags[attrName]
	return ok
}

func (t *Item) GetTag(attrName string) (string, bool) {
	if len(t.Tags) == 0 {
		return "", false
	}
	v, ok := t.Tags[attrName]
	return v, ok
}

// DefEnum 枚举定义
type DefEnum struct {
	DefTypeBase
	IsFlags        bool
	IsUniqueItemId bool
	Items          []Item
}

func NewDefEnum(enum *rawdefs.RawEnum) *DefEnum {
	defEnum := &DefEnum{
		DefTypeBase: DefTypeBase{
			Name:        enum.Name,
			namespace:   enum.Namespace,
			Comment:     enum.Comment,
			Tags:        enum.Tags,
			Groups:      enum.Groups,
			TypeMappers: enum.TypeMappers,
		},
		IsFlags:        enum.IsFlags,
		IsUniqueItemId: enum.IsUniqueItemId,
	}
	for _, item := range enum.Items {
		comment := item.Comment
		if len(comment) == 0 {
			comment = item.Alias
		}
		defEnum.Items = append(defEnum.Items, Item{
			Name:    item.Name,
			Alias:   item.Alias,
			Value:   item.Value,
			Tags:    item.Tags,
			Comment: comment,
		})
	}
	return defEnum
}

func (e *DefEnum) PreCompile() {
	//TODO implement me
	panic("implement me")
}

func (e *DefEnum) Compile() {
	//TODO implement me
	panic("implement me")
}

func (e *DefEnum) PostCompile() {
	//TODO implement me
	panic("implement me")
}
