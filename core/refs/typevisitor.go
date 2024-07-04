package refs

import (
	"reflect"
	"sync"
)

type ITypeVisitor interface {
	Name() string
}

func DispatchAccept0[R DType](visitor ITypeVisitor0[R], ttype TType) R {
	switch ttype.(type) {
	case mustEmbedTBool:
		return visitor.AcceptBool(ttype)
	case mustEmbedTByte:
		return visitor.AcceptByte(ttype)
	case mustEmbedTShort:
		return visitor.AcceptShort(ttype)
	case mustEmbedTInt:
		return visitor.AcceptInt(ttype)
	case mustEmbedTLong:
		return visitor.AcceptLong(ttype)
	case mustEmbedTFloat:
		return visitor.AcceptFloat(ttype)
	case mustEmbedTDouble:
		return visitor.AcceptDouble(ttype)
	case mustEmbedTEnum:
		return visitor.AcceptEnum(ttype)
	case mustEmbedTString:
		return visitor.AcceptString(ttype)
	case mustEmbedTDateTime:
		return visitor.AcceptDateTime(ttype)
	case mustEmbedTBean:
		return visitor.AcceptBean(ttype)
	case mustEmbedTArray:
		return visitor.AcceptArray(ttype)
	case mustEmbedTList:
		return visitor.AcceptList(ttype)
	case mustEmbedTSet:
		return visitor.AcceptSet(ttype)
	case mustEmbedTMap:
		return visitor.AcceptMap(ttype)
	default:
		panic("not support type: " + ttype.TypeName())
	}
}

type ITypeVisitor0[R DType] interface {
	ITypeVisitor
	AcceptBool(TType) R
	AcceptByte(TType) R
	AcceptShort(TType) R
	AcceptInt(TType) R
	AcceptLong(TType) R
	AcceptFloat(TType) R
	AcceptDouble(TType) R
	AcceptEnum(TType) R
	AcceptString(TType) R
	AcceptDateTime(TType) R
	AcceptBean(TType) R
	AcceptArray(TType) R
	AcceptList(TType) R
	AcceptSet(TType) R
	AcceptMap(TType) R
}

func DispatchAccept1[T any, R DType](visitor ITypeVisitor1[T, R], ttype TType, t T) R {
	switch ttype.(type) {
	case mustEmbedTBool:
		return visitor.AcceptBool(ttype, t)
	case mustEmbedTByte:
		return visitor.AcceptByte(ttype, t)
	case mustEmbedTShort:
		return visitor.AcceptShort(ttype, t)
	case mustEmbedTInt:
		return visitor.AcceptInt(ttype, t)
	case mustEmbedTLong:
		return visitor.AcceptLong(ttype, t)
	case mustEmbedTFloat:
		return visitor.AcceptFloat(ttype, t)
	case mustEmbedTDouble:
		return visitor.AcceptDouble(ttype, t)
	case mustEmbedTEnum:
		return visitor.AcceptEnum(ttype, t)
	case mustEmbedTString:
		return visitor.AcceptString(ttype, t)
	case mustEmbedTDateTime:
		return visitor.AcceptDateTime(ttype, t)
	case mustEmbedTBean:
		return visitor.AcceptBean(ttype, t)
	case mustEmbedTArray:
		return visitor.AcceptArray(ttype, t)
	case mustEmbedTList:
		return visitor.AcceptList(ttype, t)
	case mustEmbedTSet:
		return visitor.AcceptSet(ttype, t)
	case mustEmbedTMap:
		return visitor.AcceptMap(ttype, t)
	default:
		panic("not support type: " + ttype.TypeName())
	}
}

type ITypeVisitor1[T any, R DType] interface {
	ITypeVisitor
	AcceptBool(TType, T) R
	AcceptByte(TType, T) R
	AcceptShort(TType, T) R
	AcceptInt(TType, T) R
	AcceptLong(TType, T) R
	AcceptFloat(TType, T) R
	AcceptDouble(TType, T) R
	AcceptEnum(TType, T) R
	AcceptString(TType, T) R
	AcceptDateTime(TType, T) R
	AcceptBean(TType, T) R
	AcceptArray(TType, T) R
	AcceptList(TType, T) R
	AcceptSet(TType, T) R
	AcceptMap(TType, T) R
}

func DispatchAccept2[T, S any, R DType](visitor ITypeVisitor2[T, S, R], ttype TType, t T, s S) R {
	switch ttype.(type) {
	case mustEmbedTBool:
		return visitor.AcceptBool(ttype, t, s)
	case mustEmbedTByte:
		return visitor.AcceptByte(ttype, t, s)
	case mustEmbedTShort:
		return visitor.AcceptShort(ttype, t, s)
	case mustEmbedTInt:
		return visitor.AcceptInt(ttype, t, s)
	case mustEmbedTLong:
		return visitor.AcceptLong(ttype, t, s)
	case mustEmbedTFloat:
		return visitor.AcceptFloat(ttype, t, s)
	case mustEmbedTDouble:
		return visitor.AcceptDouble(ttype, t, s)
	case mustEmbedTEnum:
		return visitor.AcceptEnum(ttype, t, s)
	case mustEmbedTString:
		return visitor.AcceptString(ttype, t, s)
	case mustEmbedTDateTime:
		return visitor.AcceptDateTime(ttype, t, s)
	case mustEmbedTBean:
		return visitor.AcceptBean(ttype, t, s)
	case mustEmbedTArray:
		return visitor.AcceptArray(ttype, t, s)
	case mustEmbedTList:
		return visitor.AcceptList(ttype, t, s)
	case mustEmbedTSet:
		return visitor.AcceptSet(ttype, t, s)
	case mustEmbedTMap:
		return visitor.AcceptMap(ttype, t, s)
	default:
		panic("not support type: " + ttype.TypeName())
	}
}

type ITypeVisitor2[T, S any, R DType] interface {
	ITypeVisitor
	AcceptBool(TType, T, S) R
	AcceptByte(TType, T, S) R
	AcceptShort(TType, T, S) R
	AcceptInt(TType, T, S) R
	AcceptLong(TType, T, S) R
	AcceptFloat(TType, T, S) R
	AcceptDouble(TType, T, S) R
	AcceptEnum(TType, T, S) R
	AcceptString(TType, T, S) R
	AcceptDateTime(TType, T, S) R
	AcceptBean(TType, T, S) R
	AcceptArray(TType, T, S) R
	AcceptList(TType, T, S) R
	AcceptSet(TType, T, S) R
	AcceptMap(TType, T, S) R
}

func DispatchAccept3[T, S, U any, R DType](visitor ITypeVisitor3[T, S, U, R], ttype TType, t T, s S, u U) R {
	switch ttype.(type) {
	case mustEmbedTBool:
		return visitor.AcceptBool(ttype, t, s, u)
	case mustEmbedTByte:
		return visitor.AcceptByte(ttype, t, s, u)
	case mustEmbedTShort:
		return visitor.AcceptShort(ttype, t, s, u)
	case mustEmbedTInt:
		return visitor.AcceptInt(ttype, t, s, u)
	case mustEmbedTLong:
		return visitor.AcceptLong(ttype, t, s, u)
	case mustEmbedTFloat:
		return visitor.AcceptFloat(ttype, t, s, u)
	case mustEmbedTDouble:
		return visitor.AcceptDouble(ttype, t, s, u)
	case mustEmbedTEnum:
		return visitor.AcceptEnum(ttype, t, s, u)
	case mustEmbedTString:
		return visitor.AcceptString(ttype, t, s, u)
	case mustEmbedTDateTime:
		return visitor.AcceptDateTime(ttype, t, s, u)
	case mustEmbedTBean:
		return visitor.AcceptBean(ttype, t, s, u)
	case mustEmbedTArray:
		return visitor.AcceptArray(ttype, t, s, u)
	case mustEmbedTList:
		return visitor.AcceptList(ttype, t, s, u)
	case mustEmbedTSet:
		return visitor.AcceptSet(ttype, t, s, u)
	case mustEmbedTMap:
		return visitor.AcceptMap(ttype, t, s, u)
	default:
		panic("not support type: " + ttype.TypeName())
	}
}

type ITypeVisitor3[T, S, U any, R DType] interface {
	ITypeVisitor
	AcceptBool(TType, T, S, U) R
	AcceptByte(TType, T, S, U) R
	AcceptShort(TType, T, S, U) R
	AcceptInt(TType, T, S, U) R
	AcceptLong(TType, T, S, U) R
	AcceptFloat(TType, T, S, U) R
	AcceptDouble(TType, T, S, U) R
	AcceptEnum(TType, T, S, U) R
	AcceptString(TType, T, S, U) R
	AcceptDateTime(TType, T, S, U) R
	AcceptBean(TType, T, S, U) R
	AcceptArray(TType, T, S, U) R
	AcceptList(TType, T, S, U) R
	AcceptSet(TType, T, S, U) R
	AcceptMap(TType, T, S, U) R
}

func DispatchAccept4[T, S, U, V any, R DType](visitor ITypeVisitor4[T, S, U, V, R], ttype TType, t T, s S, u U, v V) R {
	switch ttype.(type) {
	case mustEmbedTBool:
		return visitor.AcceptBool(ttype, t, s, u, v)
	case mustEmbedTByte:
		return visitor.AcceptByte(ttype, t, s, u, v)
	case mustEmbedTShort:
		return visitor.AcceptShort(ttype, t, s, u, v)
	case mustEmbedTInt:
		return visitor.AcceptInt(ttype, t, s, u, v)
	case mustEmbedTLong:
		return visitor.AcceptLong(ttype, t, s, u, v)
	case mustEmbedTFloat:
		return visitor.AcceptFloat(ttype, t, s, u, v)
	case mustEmbedTDouble:
		return visitor.AcceptDouble(ttype, t, s, u, v)
	case mustEmbedTEnum:
		return visitor.AcceptEnum(ttype, t, s, u, v)
	case mustEmbedTString:
		return visitor.AcceptString(ttype, t, s, u, v)
	case mustEmbedTDateTime:
		return visitor.AcceptDateTime(ttype, t, s, u, v)
	case mustEmbedTBean:
		return visitor.AcceptBean(ttype, t, s, u, v)
	case mustEmbedTArray:
		return visitor.AcceptArray(ttype, t, s, u, v)
	case mustEmbedTList:
		return visitor.AcceptList(ttype, t, s, u, v)
	case mustEmbedTSet:
		return visitor.AcceptSet(ttype, t, s, u, v)
	case mustEmbedTMap:
		return visitor.AcceptMap(ttype, t, s, u, v)
	default:
		panic("not support type: " + ttype.TypeName())
	}
}

type ITypeVisitor4[T, S, U, V any, R DType] interface {
	ITypeVisitor
	AcceptBool(TType, T, S, U, V) R
	AcceptByte(TType, T, S, U, V) R
	AcceptShort(TType, T, S, U, V) R
	AcceptInt(TType, T, S, U, V) R
	AcceptLong(TType, T, S, U, V) R
	AcceptFloat(TType, T, S, U, V) R
	AcceptDouble(TType, T, S, U, V) R
	AcceptEnum(TType, T, S, U, V) R
	AcceptString(TType, T, S, U, V) R
	AcceptDateTime(TType, T, S, U, V) R
	AcceptBean(TType, T, S, U, V) R
	AcceptArray(TType, T, S, U, V) R
	AcceptList(TType, T, S, U, V) R
	AcceptSet(TType, T, S, U, V) R
	AcceptMap(TType, T, S, U, V) R
}

var (
	typeVisitorManager = map[string]ITypeVisitor{}
	locker             sync.RWMutex
)

func RegisterTypeVisitor[T ITypeVisitor]() {
	var t T
	typ := reflect.TypeOf(t).Elem()
	// 创建类型T的新实例
	instance := reflect.New(typ).Interface().(ITypeVisitor)

	locker.Lock()
	defer locker.Unlock()
	if _, ok := typeVisitorManager[instance.Name()]; ok {
		panic("type visitor: " + instance.Name() + " already registered")
	}
	typeVisitorManager[instance.Name()] = instance
}

func GetTypeVisitor(name string) ITypeVisitor {
	locker.RLock()
	defer locker.RUnlock()
	if visitor, ok := typeVisitorManager[name]; ok {
		return visitor
	}
	return nil
}

func MustGetTypeVisitor[T ITypeVisitor](name string) T {
	locker.RLock()
	defer locker.RUnlock()
	if visitor, ok := typeVisitorManager[name]; ok {
		result, ok := visitor.(T)
		if !ok {
			panic("type visitor: " + name + " not match")
		}
		return result
	} else {
		panic("type visitor: " + name + " not registered")
	}
}
