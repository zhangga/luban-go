package types

import (
	"github.com/zhangga/luban/core/refs"
	"strconv"
)

var _ refs.TType = (*TDouble)(nil)

type TDouble struct {
	refs.EmbedTType
	refs.EmbedTDouble
}

func NewTDouble(isNullable bool, tags map[string]string, _ refs.IDefType, _ ...refs.TType) refs.TType {
	return &TDouble{
		EmbedTType: refs.NewEmbedTType(isNullable, tags),
	}
}

func (d *TDouble) TypeName() string {
	return "double"
}

func (d *TDouble) IsBean() bool {
	return false
}

func (d *TDouble) IsNullable() bool {
	return d.EmbedTType.IsNullable
}

func (d *TDouble) TryParseFrom(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}
