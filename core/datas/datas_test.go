package datas

import (
	"github.com/zhangga/luban-go/core/types"
	"testing"
)

type dummyDataVisitor struct{}

func (v *dummyDataVisitor) VisitDBool(d *DBool) interface{}         { return d.TypeName() }
func (v *dummyDataVisitor) VisitDByte(d *DByte) interface{}         { return d.TypeName() }
func (v *dummyDataVisitor) VisitDShort(d *DShort) interface{}       { return d.TypeName() }
func (v *dummyDataVisitor) VisitDInt(d *DInt) interface{}           { return d.TypeName() }
func (v *dummyDataVisitor) VisitDLong(d *DLong) interface{}         { return d.TypeName() }
func (v *dummyDataVisitor) VisitDFloat(d *DFloat) interface{}       { return d.TypeName() }
func (v *dummyDataVisitor) VisitDDouble(d *DDouble) interface{}     { return d.TypeName() }
func (v *dummyDataVisitor) VisitDString(d *DString) interface{}     { return d.TypeName() }
func (v *dummyDataVisitor) VisitDText(d *DText) interface{}         { return d.TypeName() }
func (v *dummyDataVisitor) VisitDEnum(d *DEnum) interface{}         { return d.TypeName() }
func (v *dummyDataVisitor) VisitDBean(d *DBean) interface{}         { return d.TypeName() }
func (v *dummyDataVisitor) VisitDArray(d *DArray) interface{}       { return d.TypeName() }
func (v *dummyDataVisitor) VisitDList(d *DList) interface{}         { return d.TypeName() }
func (v *dummyDataVisitor) VisitDSet(d *DSet) interface{}           { return d.TypeName() }
func (v *dummyDataVisitor) VisitDMap(d *DMap) interface{}           { return d.TypeName() }
func (v *dummyDataVisitor) VisitDDateTime(d *DDateTime) interface{} { return d.TypeName() }

func TestDataNames(t *testing.T) {
	v := &dummyDataVisitor{}

	dInt := NewDInt(123)
	if dInt.Accept(v) != "int" {
		t.Errorf("Expected int")
	}

	if dInt.String() != "123" {
		t.Errorf("Expected 123, got %s", dInt.String())
	}

	dList := NewDList(types.NewTList(false, types.NewTInt(false, nil), nil), []DType{NewDInt(1), NewDInt(2)})
	if dList.Accept(v) != "list" {
		t.Errorf("Expected list")
	}

	if dList.String() != "[1, 2]" {
		t.Errorf("Expected [1, 2], got %s", dList.String())
	}
}
