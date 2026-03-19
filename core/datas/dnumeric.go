package datas

import "fmt"

type DByte struct{ Value byte }

func NewDByte(v byte) *DByte                             { return &DByte{Value: v} }
func (d *DByte) TypeName() string                        { return "byte" }
func (d *DByte) Accept(visitor IDataVisitor) interface{} { return visitor.VisitDByte(d) }
func (d *DByte) String() string                          { return fmt.Sprintf("%v", d.Value) }

type DShort struct{ Value int16 }

func NewDShort(v int16) *DShort                           { return &DShort{Value: v} }
func (d *DShort) TypeName() string                        { return "short" }
func (d *DShort) Accept(visitor IDataVisitor) interface{} { return visitor.VisitDShort(d) }
func (d *DShort) String() string                          { return fmt.Sprintf("%v", d.Value) }

type DInt struct{ Value int32 }

func NewDInt(v int32) *DInt                             { return &DInt{Value: v} }
func (d *DInt) TypeName() string                        { return "int" }
func (d *DInt) Accept(visitor IDataVisitor) interface{} { return visitor.VisitDInt(d) }
func (d *DInt) String() string                          { return fmt.Sprintf("%v", d.Value) }

type DLong struct{ Value int64 }

func NewDLong(v int64) *DLong                            { return &DLong{Value: v} }
func (d *DLong) TypeName() string                        { return "long" }
func (d *DLong) Accept(visitor IDataVisitor) interface{} { return visitor.VisitDLong(d) }
func (d *DLong) String() string                          { return fmt.Sprintf("%v", d.Value) }

type DFloat struct{ Value float32 }

func NewDFloat(v float32) *DFloat                         { return &DFloat{Value: v} }
func (d *DFloat) TypeName() string                        { return "float" }
func (d *DFloat) Accept(visitor IDataVisitor) interface{} { return visitor.VisitDFloat(d) }
func (d *DFloat) String() string                          { return fmt.Sprintf("%v", d.Value) }

type DDouble struct{ Value float64 }

func NewDDouble(v float64) *DDouble                        { return &DDouble{Value: v} }
func (d *DDouble) TypeName() string                        { return "double" }
func (d *DDouble) Accept(visitor IDataVisitor) interface{} { return visitor.VisitDDouble(d) }
func (d *DDouble) String() string                          { return fmt.Sprintf("%v", d.Value) }
