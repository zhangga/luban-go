package schema

import (
	"github.com/zhangga/luban/core/refs"
	"github.com/zhangga/luban/internal/rawdefs"
)

type CollectorBase struct {
	rawTables []*rawdefs.RawTable
	rawEnums  []*rawdefs.RawEnum
	rawBeans  []*rawdefs.RawBean
}

func (c *CollectorBase) AddTable(t refs.UnimplementedTable) {
	c.rawTables = append(c.rawTables, t.(*rawdefs.RawTable))
}

func (c *CollectorBase) AddEnum(e refs.UnimplementedEnum) {
	c.rawEnums = append(c.rawEnums, e.(*rawdefs.RawEnum))
}

func (c *CollectorBase) AddBean(b refs.UnimplementedBean) {
	c.rawBeans = append(c.rawBeans, b.(*rawdefs.RawBean))
}

func (c *CollectorBase) AddRefGroup(g refs.UnimplementedRefGroup) {
	//c.rawTables = append(c.rawTables, t.(*rawrefs.RawTable))
}
