package rawdefs

import (
	"github.com/zhangga/luban/core/refs"
	"github.com/zhangga/luban/internal/utils"
)

var _ refs.UnimplementedBean = (*RawBean)(nil)

type RawBean struct {
	Namespace   string
	Name        string
	Parent      string
	IsValueType bool
	Comment     string
	Tags        map[string]string
	Alias       string
	Sep         string
	Groups      []string
	Fields      []*RawField
	TypeMappers []*TypeMapper
}

func (r *RawBean) MustEmbedUnimplementedBean() {}

func (r *RawBean) FullName() string {
	return utils.MakeFullName(r.Namespace, r.Name)
}
