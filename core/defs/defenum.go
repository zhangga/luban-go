package defs

import "github.com/zhangga/luban-go/core/rawdefs"

type DefEnumImpl struct {
	Raw *rawdefs.RawEnum
}

func NewDefEnumImpl(raw *rawdefs.RawEnum) *DefEnumImpl {
	return &DefEnumImpl{Raw: raw}
}

func (e *DefEnumImpl) FullName() string {
	return e.Raw.FullName()
}

func (e *DefEnumImpl) Name() string {
	return e.Raw.Name
}

func (e *DefEnumImpl) Namespace() string {
	return e.Raw.Namespace
}
