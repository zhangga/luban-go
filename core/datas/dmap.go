package datas

import (
	"github.com/zhangga/luban-go/core/types"
	"strings"
)

type DMap struct {
	Type  *types.TMap
	Datas map[DType]DType
}

func NewDMap(t *types.TMap, datas map[DType]DType) *DMap {
	return &DMap{Type: t, Datas: datas}
}

func (d *DMap) TypeName() string { return "map" }

func (d *DMap) Accept(visitor IDataVisitor) interface{} { return visitor.VisitDMap(d) }

func (d *DMap) String() string {
	var sb strings.Builder
	sb.WriteString("{")
	i := 0
	for k, v := range d.Datas {
		if i > 0 {
			sb.WriteString(", ")
		}
		sb.WriteString(k.String())
		sb.WriteString(":")
		if v != nil {
			sb.WriteString(v.String())
		} else {
			sb.WriteString("null")
		}
		i++
	}
	sb.WriteString("}")
	return sb.String()
}
