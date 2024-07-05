package types

import (
	"github.com/zhangga/luban/core/refs"
)

func init() {
	refs.RegisterTypeCreator(NewTBean)
	refs.RegisterTypeCreator(NewTBool)
	refs.RegisterTypeCreator(NewTByte)
	refs.RegisterTypeCreator(NewTDateTime)
	refs.RegisterTypeCreator(NewTDouble)
	refs.RegisterTypeCreator(NewTEnum)
	refs.RegisterTypeCreator(NewTFloat)
	refs.RegisterTypeCreator(NewTInt)
	refs.RegisterTypeCreator(NewTList)
	refs.RegisterTypeCreator(NewTLong)
	refs.RegisterTypeCreator(NewTMap)
	refs.RegisterTypeCreator(NewTSet)
	refs.RegisterTypeCreator(NewTShort)
	refs.RegisterTypeCreator(NewTString)

}
