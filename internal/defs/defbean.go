package defs

import (
	"github.com/zhangga/luban/internal/rawdefs"
	"github.com/zhangga/luban/internal/utils"
)

type DefBean struct {
	DefTypeBase
	Id                           int64
	AutoId                       int // for protobuf
	Parent                       string
	ParentDefType                *DefBean
	Children                     []*DefBean
	HierarchyNotAbstractChildren []*DefBean
}

func NewDefBean(rawBean rawdefs.RawBean) *DefBean {
	bean := &DefBean{
		DefTypeBase: DefTypeBase{
			Name:        rawBean.Name,
			Namespace:   rawBean.Namespace,
			Comment:     rawBean.Comment,
			Tags:        rawBean.Tags,
			TypeMappers: rawBean.TypeMappers,
		},
		Parent: rawBean.Parent,
		Id:     utils.ComputeCfgHashIdByName(rawBean.FullName()),
	}
	for _, field := range rawBean.Fields {
		bean.Fields = append(bean.Fields, bean.CreateField(field, 0))
	}
	return bean
}

func (b *DefBean) CreateField(field rawdefs.RawField, idOffset int) DefField {

}

func (b *DefBean) RootDefType() *DefBean {
	if b.ParentDefType == nil {
		return b
	}
	return b.ParentDefType.RootDefType()
}
