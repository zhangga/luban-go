package datas

import (
	"github.com/zhangga/luban-go/core/types"
	"strings"
)

type DSet struct {
	Type     *types.TSet
	Elements []DType
}

func NewDSet(t *types.TSet, elements []DType) *DSet {
	return &DSet{Type: t, Elements: elements}
}

func (d *DSet) TypeName() string { return "set" }

func (d *DSet) Accept(visitor IDataVisitor) interface{} { return visitor.VisitDSet(d) }

func (d *DSet) String() string {
	var sb strings.Builder
	sb.WriteString("{")
	for i, e := range d.Elements {
		if i > 0 {
			sb.WriteString(", ")
		}
		if e != nil {
			sb.WriteString(e.String())
		} else {
			sb.WriteString("null")
		}
	}
	sb.WriteString("}")
	return sb.String()
}
