package datas

import (
	"fmt"
	"github.com/zhangga/luban-go/core/types"
)

type DEnum struct {
	Value    int32
	StrValue string
	Type     *types.TEnum
}

func NewDEnum(v int32, strValue string, t *types.TEnum) *DEnum {
	return &DEnum{
		Value:    v,
		StrValue: strValue,
		Type:     t,
	}
}

func (d *DEnum) TypeName() string {
	return "enum"
}

func (d *DEnum) Accept(visitor IDataVisitor) interface{} {
	return visitor.VisitDEnum(d)
}

func (d *DEnum) String() string {
	if d.StrValue != "" {
		return d.StrValue
	}
	return fmt.Sprintf("%d", d.Value)
}
