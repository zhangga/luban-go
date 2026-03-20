package types

import (
	"testing"
)

type dummyVisitor struct{}

func (v *dummyVisitor) VisitTBool(t *TBool) interface{}         { return t.TypeName() }
func (v *dummyVisitor) VisitTByte(t *TByte) interface{}         { return t.TypeName() }
func (v *dummyVisitor) VisitTShort(t *TShort) interface{}       { return t.TypeName() }
func (v *dummyVisitor) VisitTInt(t *TInt) interface{}           { return t.TypeName() }
func (v *dummyVisitor) VisitTLong(t *TLong) interface{}         { return t.TypeName() }
func (v *dummyVisitor) VisitTFloat(t *TFloat) interface{}       { return t.TypeName() }
func (v *dummyVisitor) VisitTDouble(t *TDouble) interface{}     { return t.TypeName() }
func (v *dummyVisitor) VisitTString(t *TString) interface{}     { return t.TypeName() }
func (v *dummyVisitor) VisitTText(t *TText) interface{}         { return t.TypeName() }
func (v *dummyVisitor) VisitTEnum(t *TEnum) interface{}         { return t.TypeName() }
func (v *dummyVisitor) VisitTBean(t *TBean) interface{}         { return t.TypeName() }
func (v *dummyVisitor) VisitTArray(t *TArray) interface{}       { return t.TypeName() }
func (v *dummyVisitor) VisitTList(t *TList) interface{}         { return t.TypeName() }
func (v *dummyVisitor) VisitTSet(t *TSet) interface{}           { return t.TypeName() }
func (v *dummyVisitor) VisitTMap(t *TMap) interface{}           { return t.TypeName() }
func (v *dummyVisitor) VisitTDateTime(t *TDateTime) interface{} { return t.TypeName() }

func TestTypeNames(t *testing.T) {
	v := &dummyVisitor{}

	tb := NewTBool(false, nil)
	if tb.Accept(v) != "bool" {
		t.Errorf("Expected bool")
	}

	tint := NewTInt(false, nil)
	if tint.Accept(v) != "int" {
		t.Errorf("Expected int")
	}

	tlist := NewTList(false, tint, nil)
	if tlist.Accept(v) != "list" {
		t.Errorf("Expected list")
	}
	if tlist.ElementType().TypeName() != "int" {
		t.Errorf("Expected element int")
	}
}
