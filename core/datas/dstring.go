package datas

type DString struct {
	Value string
}

func NewDString(v string) *DString {
	return &DString{Value: v}
}

func (d *DString) TypeName() string {
	return "string"
}

func (d *DString) Accept(visitor IDataVisitor) interface{} {
	return visitor.VisitDString(d)
}

func (d *DString) String() string {
	return d.Value
}
