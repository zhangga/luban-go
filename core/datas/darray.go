package datas

import (
	"github.com/zhangga/luban-go/core/types"
	"strings"
)

type DArray struct {
	Type     *types.TArray
	Elements []DType
}

func NewDArray(t *types.TArray, elements []DType) *DArray {
	return &DArray{
		Type:     t,
		Elements: elements,
	}
}

func (d *DArray) TypeName() string {
	return "array"
}

func (d *DArray) Accept(visitor IDataVisitor) interface{} {
	return visitor.VisitDArray(d)
}

func (d *DArray) String() string {
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
