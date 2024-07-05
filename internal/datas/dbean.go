package datas

import (
	"github.com/zhangga/luban/core/refs"
	"github.com/zhangga/luban/internal/defs"
	"github.com/zhangga/luban/internal/types"
)

var _ refs.DType = (*DBean)(nil)

type DBean struct {
	ttype    *types.TBean
	implType *defs.DefBean
	fields   []refs.DType
}

func NewDBean(defType *types.TBean, implType *defs.DefBean, fields []refs.DType) *DBean {
	return &DBean{
		ttype:    defType,
		implType: implType,
		fields:   fields,
	}
}

func (b *DBean) TypeName() string {
	return "bean"
}

func (b *DBean) CompareTo(other refs.DType) int {
	//TODO implement me
	panic("implement me")
}

func (b *DBean) String() string {
	//TODO implement me
	panic("implement me")
}
