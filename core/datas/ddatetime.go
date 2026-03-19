package datas

import "time"

type DDateTime struct {
	Value time.Time
}

func NewDDateTime(v time.Time) *DDateTime {
	return &DDateTime{Value: v}
}

func (d *DDateTime) TypeName() string {
	return "datetime"
}

func (d *DDateTime) Accept(visitor IDataVisitor) interface{} {
	return visitor.VisitDDateTime(d)
}

func (d *DDateTime) String() string {
	return d.Value.Format("2006-01-02 15:04:05")
}
