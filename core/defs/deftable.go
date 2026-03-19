package defs

import "github.com/zhangga/luban-go/core/rawdefs"

type DefTable struct {
	Raw *rawdefs.RawTable
}

func NewDefTable(raw *rawdefs.RawTable) *DefTable {
	return &DefTable{Raw: raw}
}

func (t *DefTable) FullName() string {
	return t.Raw.FullName()
}

func (t *DefTable) Name() string {
	return t.Raw.Name
}

func (t *DefTable) Namespace() string {
	return t.Raw.Namespace
}
