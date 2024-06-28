package rawdefs

import "github.com/zhangga/luban/core/refs"

var _ refs.UnimplementedRefGroup = (*RawRefGroup)(nil)

type RawRefGroup struct {
	Name string
	Refs []string
}

func (r RawRefGroup) MustEmbedUnimplementedRefGroup() {}
