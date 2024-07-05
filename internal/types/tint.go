package types

import (
	"github.com/zhangga/luban/core/refs"
	"strconv"
)

var _ refs.TType = (*TInt)(nil)

type TInt struct {
	refs.EmbedTType
	refs.EmbedTInt
}

func NewTInt(isNullable bool, tags map[string]string, _ refs.IDefType, _ ...refs.TType) refs.TType {
	return &TInt{
		EmbedTType: refs.NewEmbedTType(isNullable, tags),
	}
}

func (i *TInt) TypeName() string {
	return "int"
}

func (i *TInt) IsBean() bool {
	return false
}

func (i *TInt) IsNullable() bool {
	return i.EmbedTType.IsNullable
}

func (i *TInt) TryParseFrom(s string) bool {
	_, err := strconv.ParseInt(s, 10, 32)
	return err == nil
}
