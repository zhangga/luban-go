package types

import "strconv"

type TByte struct{ *TTypeBase }

func NewTByte(isNullable bool, tags map[string]string) *TByte {
	return &TByte{TTypeBase: NewTTypeBase(isNullable, tags)}
}
func (t *TByte) TypeName() string                        { return "byte" }
func (t *TByte) TryParseFrom(s string) bool              { _, err := strconv.ParseUint(s, 10, 8); return err == nil }
func (t *TByte) Accept(visitor ITypeVisitor) interface{} { return visitor.VisitTByte(t) }

type TShort struct{ *TTypeBase }

func NewTShort(isNullable bool, tags map[string]string) *TShort {
	return &TShort{TTypeBase: NewTTypeBase(isNullable, tags)}
}
func (t *TShort) TypeName() string { return "short" }
func (t *TShort) TryParseFrom(s string) bool {
	_, err := strconv.ParseInt(s, 10, 16)
	return err == nil
}
func (t *TShort) Accept(visitor ITypeVisitor) interface{} { return visitor.VisitTShort(t) }

type TInt struct{ *TTypeBase }

func NewTInt(isNullable bool, tags map[string]string) *TInt {
	return &TInt{TTypeBase: NewTTypeBase(isNullable, tags)}
}
func (t *TInt) TypeName() string                        { return "int" }
func (t *TInt) TryParseFrom(s string) bool              { _, err := strconv.ParseInt(s, 10, 32); return err == nil }
func (t *TInt) Accept(visitor ITypeVisitor) interface{} { return visitor.VisitTInt(t) }

type TLong struct{ *TTypeBase }

func NewTLong(isNullable bool, tags map[string]string) *TLong {
	return &TLong{TTypeBase: NewTTypeBase(isNullable, tags)}
}
func (t *TLong) TypeName() string                        { return "long" }
func (t *TLong) TryParseFrom(s string) bool              { _, err := strconv.ParseInt(s, 10, 64); return err == nil }
func (t *TLong) Accept(visitor ITypeVisitor) interface{} { return visitor.VisitTLong(t) }

type TFloat struct{ *TTypeBase }

func NewTFloat(isNullable bool, tags map[string]string) *TFloat {
	return &TFloat{TTypeBase: NewTTypeBase(isNullable, tags)}
}
func (t *TFloat) TypeName() string                        { return "float" }
func (t *TFloat) TryParseFrom(s string) bool              { _, err := strconv.ParseFloat(s, 32); return err == nil }
func (t *TFloat) Accept(visitor ITypeVisitor) interface{} { return visitor.VisitTFloat(t) }

type TDouble struct{ *TTypeBase }

func NewTDouble(isNullable bool, tags map[string]string) *TDouble {
	return &TDouble{TTypeBase: NewTTypeBase(isNullable, tags)}
}
func (t *TDouble) TypeName() string                        { return "double" }
func (t *TDouble) TryParseFrom(s string) bool              { _, err := strconv.ParseFloat(s, 64); return err == nil }
func (t *TDouble) Accept(visitor ITypeVisitor) interface{} { return visitor.VisitTDouble(t) }
