package datas

import (
	"github.com/zhangga/luban-go/core/types"
	"strings"
)

type DBean struct {
	Type     *types.TBean
	ImplType *types.TBean // 如果是多态的话，这个字段表示具体的子类型
	Fields   []DType
}

func NewDBean(t *types.TBean, implType *types.TBean, fields []DType) *DBean {
	return &DBean{
		Type:     t,
		ImplType: implType,
		Fields:   fields,
	}
}

func (d *DBean) TypeName() string {
	return "bean"
}

func (d *DBean) Accept(visitor IDataVisitor) interface{} {
	return visitor.VisitDBean(d)
}

func (d *DBean) String() string {
	var sb strings.Builder
	sb.WriteString("{")
	for i, f := range d.Fields {
		if i > 0 {
			sb.WriteString(", ")
		}
		if f != nil {
			sb.WriteString(f.String())
		} else {
			sb.WriteString("null")
		}
	}
	sb.WriteString("}")
	return sb.String()
}
