package datas

import (
	"github.com/zhangga/luban-go/core/types"
	"strings"
)

type DList struct {
	Type     *types.TList
	Elements []DType
}

func NewDList(t *types.TList, elements []DType) *DList {
	return &DList{Type: t, Elements: elements}
}

func (d *DList) TypeName() string { return "list" }

func (d *DList) Accept(visitor IDataVisitor) interface{} { return visitor.VisitDList(d) }

func (d *DList) String() string {
	var sb strings.Builder
	sb.WriteString("[")
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
	sb.WriteString("]")
	return sb.String()
}
