package rawdefs

import (
	"github.com/zhangga/luban/core/refs"
	"github.com/zhangga/luban/internal/utils"
)

var _ refs.UnimplementedEnum = (*RawEnum)(nil)

type RawEnum struct {
	Namespace      string
	Name           string
	IsFlags        bool
	IsUniqueItemId bool
	Comment        string
	Tags           map[string]string
	Items          []EnumItem
	Groups         []string
	TypeMappers    []TypeMapper
}

func (e *RawEnum) MustEmbedUnimplementedEnum() {}

func (e *RawEnum) FullName() string {
	return utils.MakeFullName(e.Namespace, e.Name)
}

type EnumItem struct {
	Name    string
	Alias   string
	Value   string
	Comment string
	Tags    map[string]string
}
