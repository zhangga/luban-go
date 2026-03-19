package datas

type DBool struct {
	Value bool
}

func NewDBool(v bool) *DBool {
	return &DBool{Value: v}
}

func (d *DBool) TypeName() string {
	return "bool"
}

func (d *DBool) Accept(visitor IDataVisitor) interface{} {
	return visitor.VisitDBool(d)
}

func (d *DBool) String() string {
	if d.Value {
		return "true"
	}
	return "false"
}
