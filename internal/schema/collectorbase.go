package schema

import (
	"github.com/zhangga/luban/core/refs"
	"github.com/zhangga/luban/internal/rawrefs"
)

type CollectorBase struct {
	rawTables []*rawrefs.RawTable
	rawEnums  []*rawrefs.RawEnum
	rawBeans  []*rawrefs.RawBean
}

func (c *CollectorBase) AddTable(t refs.UnimplementedTable) {
	c.rawTables = append(c.rawTables, t.(*rawrefs.RawTable))
}

func (c *CollectorBase) AddEnum(e refs.UnimplementedEnum) {
	c.rawEnums = append(c.rawEnums, e.(*rawrefs.RawEnum))
}

func (c *CollectorBase) AddBean(b refs.UnimplementedBean) {
	c.rawBeans = append(c.rawBeans, b.(*rawrefs.RawBean))
}

func (c *CollectorBase) AddRefGroup(g refs.UnimplementedRefGroup) {
	//c.rawTables = append(c.rawTables, t.(*rawrefs.RawTable))
}
