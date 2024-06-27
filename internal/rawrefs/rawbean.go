package rawrefs

import "github.com/zhangga/luban/core/refs"

var _ refs.UnimplementedBean = (*RawBean)(nil)

type RawBean struct {
	Namespace   string
	Name        string
	FullName    string
	Parent      string
	IsValueType bool
	Comment     string
	Tags        map[string]string
	Alias       string
	Sep         string
	Groups      []string
	Fields      []RawField
	TypeMappers []TypeMapper
}

func (r RawBean) MustEmbedUnimplementedBean() {}
