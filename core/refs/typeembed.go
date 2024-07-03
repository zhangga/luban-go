package refs

var (
	_ mustEmbedTBool     = EmbedTBool{}
	_ mustEmbedTByte     = EmbedTByte{}
	_ mustEmbedTShort    = EmbedTShort{}
	_ mustEmbedTInt      = EmbedTInt{}
	_ mustEmbedTLong     = EmbedTLong{}
	_ mustEmbedTFloat    = EmbedTFloat{}
	_ mustEmbedTDouble   = EmbedTDouble{}
	_ mustEmbedTEnum     = EmbedTEnum{}
	_ mustEmbedTString   = EmbedTString{}
	_ mustEmbedTDateTime = EmbedTDateTime{}
	_ mustEmbedTBean     = EmbedTBean{}
	_ mustEmbedTArray    = EmbedTArray{}
	_ mustEmbedTList     = EmbedTList{}
	_ mustEmbedTSet      = EmbedTSet{}
	_ mustEmbedTMap      = EmbedTMap{}
)

type mustEmbedTBool interface {
	mustEmbedTBool()
}

type EmbedTBool struct{}

func (EmbedTBool) mustEmbedTBool() {}

type mustEmbedTByte interface {
	mustEmbedTByte()
}

type EmbedTByte struct{}

func (EmbedTByte) mustEmbedTByte() {}

type mustEmbedTShort interface {
	mustEmbedTShort()
}

type EmbedTShort struct{}

func (EmbedTShort) mustEmbedTShort() {}

type mustEmbedTInt interface {
	mustEmbedTInt()
}

type EmbedTInt struct{}

func (EmbedTInt) mustEmbedTInt() {}

type mustEmbedTLong interface {
	mustEmbedTLong()
}

type EmbedTLong struct{}

func (EmbedTLong) mustEmbedTLong() {}

type mustEmbedTFloat interface {
	mustEmbedTFloat()
}

type EmbedTFloat struct{}

func (EmbedTFloat) mustEmbedTFloat() {}

type mustEmbedTDouble interface {
	mustEmbedTDouble()
}

type EmbedTDouble struct{}

func (EmbedTDouble) mustEmbedTDouble() {}

type mustEmbedTEnum interface {
	mustEmbedTEnum()
}

type EmbedTEnum struct{}

func (EmbedTEnum) mustEmbedTEnum() {}

type mustEmbedTString interface {
	mustEmbedTString()
}

type EmbedTString struct{}

func (EmbedTString) mustEmbedTString() {}

type mustEmbedTDateTime interface {
	mustEmbedTDateTime()
}

type EmbedTDateTime struct{}

func (EmbedTDateTime) mustEmbedTDateTime() {}

type mustEmbedTBean interface {
	mustEmbedTBean()
}

type EmbedTBean struct{}

func (EmbedTBean) mustEmbedTBean() {}

type mustEmbedTArray interface {
	mustEmbedTArray()
}

type EmbedTArray struct{}

func (EmbedTArray) mustEmbedTArray() {}

type mustEmbedTList interface {
	mustEmbedTList()
}

type EmbedTList struct{}

func (EmbedTList) mustEmbedTList() {}

type mustEmbedTSet interface {
	mustEmbedTSet()
}

type EmbedTSet struct{}

func (EmbedTSet) mustEmbedTSet() {}

type mustEmbedTMap interface {
	mustEmbedTMap()
}

type EmbedTMap struct{}

func (EmbedTMap) mustEmbedTMap() {}
