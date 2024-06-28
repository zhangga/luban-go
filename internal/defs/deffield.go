package defs

import "github.com/zhangga/luban/internal/types"

type DefField struct {
	Assembly             *DefAssembly
	HostType             *DefBean
	Name                 string
	Type                 string
	CType                types.TType
	Comment              string
	AutoId               int
	Tags                 map[string]string
	IgnoreNameValidation bool
}

func (f *DefField) IsNullable() bool {
	return f.CType.IsNullable()
}
