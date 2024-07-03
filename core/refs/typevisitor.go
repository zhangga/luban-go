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
		return visitor.AcceptBool()
	case mustEmbedTByte:
		return visitor.AcceptByte()
	case mustEmbedTShort:
		return visitor.AcceptShort()
	case mustEmbedTInt:
		return visitor.AcceptInt()
	case mustEmbedTLong:
		return visitor.AcceptLong()
	case mustEmbedTFloat:
		return visitor.AcceptFloat()
	case mustEmbedTDouble:
		return visitor.AcceptDouble()
	case mustEmbedTEnum:
		return visitor.AcceptEnum()
	case mustEmbedTString:
		return visitor.AcceptString()
	case mustEmbedTDateTime:
		return visitor.AcceptDateTime()
	case mustEmbedTBean:
		return visitor.AcceptBean()
	case mustEmbedTArray:
		return visitor.AcceptArray()
	case mustEmbedTList:
		return visitor.AcceptList()
	case mustEmbedTSet:
		return visitor.AcceptSet()
	case mustEmbedTMap:
		return visitor.AcceptMap()
	default:
		panic("not support type: " + ttype.TypeName())
	}
}

type ITypeVisitor0[R DType] interface {
	ITypeVisitor
	AcceptBool() R
	AcceptByte() R
	AcceptShort() R
	AcceptInt() R
	AcceptLong() R
	AcceptFloat() R
	AcceptDouble() R
	AcceptEnum() R
	AcceptString() R
	AcceptDateTime() R
	AcceptBean() R
	AcceptArray() R
	AcceptList() R
	AcceptSet() R
	AcceptMap() R
}

func DispatchAccept1[T any, R DType](visitor ITypeVisitor1[T, R], ttype TType, t T) R {
	switch ttype.(type) {
	case mustEmbedTBool:
		return visitor.AcceptBool(t)
	case mustEmbedTByte:
		return visitor.AcceptByte(t)
	case mustEmbedTShort:
		return visitor.AcceptShort(t)
	case mustEmbedTInt:
		return visitor.AcceptInt(t)
	case mustEmbedTLong:
		return visitor.AcceptLong(t)
	case mustEmbedTFloat:
		return visitor.AcceptFloat(t)
	case mustEmbedTDouble:
		return visitor.AcceptDouble(t)
	case mustEmbedTEnum:
		return visitor.AcceptEnum(t)
	case mustEmbedTString:
		return visitor.AcceptString(t)
	case mustEmbedTDateTime:
		return visitor.AcceptDateTime(t)
	case mustEmbedTBean:
		return visitor.AcceptBean(t)
	case mustEmbedTArray:
		return visitor.AcceptArray(t)
	case mustEmbedTList:
		return visitor.AcceptList(t)
	case mustEmbedTSet:
		return visitor.AcceptSet(t)
	case mustEmbedTMap:
		return visitor.AcceptMap(t)
	default:
		panic("not support type: " + ttype.TypeName())
	}
}

type ITypeVisitor1[T any, R DType] interface {
	ITypeVisitor
	AcceptBool(T) R
	AcceptByte(T) R
	AcceptShort(T) R
	AcceptInt(T) R
	AcceptLong(T) R
	AcceptFloat(T) R
	AcceptDouble(T) R
	AcceptEnum(T) R
	AcceptString(T) R
	AcceptDateTime(T) R
	AcceptBean(T) R
	AcceptArray(T) R
	AcceptList(T) R
	AcceptSet(T) R
	AcceptMap(T) R
}

func DispatchAccept2[T, S any, R DType](visitor ITypeVisitor2[T, S, R], ttype TType, t T, s S) R {
	switch ttype.(type) {
	case mustEmbedTBool:
		return visitor.AcceptBool(t, s)
	case mustEmbedTByte:
		return visitor.AcceptByte(t, s)
	case mustEmbedTShort:
		return visitor.AcceptShort(t, s)
	case mustEmbedTInt:
		return visitor.AcceptInt(t, s)
	case mustEmbedTLong:
		return visitor.AcceptLong(t, s)
	case mustEmbedTFloat:
		return visitor.AcceptFloat(t, s)
	case mustEmbedTDouble:
		return visitor.AcceptDouble(t, s)
	case mustEmbedTEnum:
		return visitor.AcceptEnum(t, s)
	case mustEmbedTString:
		return visitor.AcceptString(t, s)
	case mustEmbedTDateTime:
		return visitor.AcceptDateTime(t, s)
	case mustEmbedTBean:
		return visitor.AcceptBean(t, s)
	case mustEmbedTArray:
		return visitor.AcceptArray(t, s)
	case mustEmbedTList:
		return visitor.AcceptList(t, s)
	case mustEmbedTSet:
		return visitor.AcceptSet(t, s)
	case mustEmbedTMap:
		return visitor.AcceptMap(t, s)
	default:
		panic("not support type: " + ttype.TypeName())
	}
}

type ITypeVisitor2[T, S any, R DType] interface {
	ITypeVisitor
	AcceptBool(T, S) R
	AcceptByte(T, S) R
	AcceptShort(T, S) R
	AcceptInt(T, S) R
	AcceptLong(T, S) R
	AcceptFloat(T, S) R
	AcceptDouble(T, S) R
	AcceptEnum(T, S) R
	AcceptString(T, S) R
	AcceptDateTime(T, S) R
	AcceptBean(T, S) R
	AcceptArray(T, S) R
	AcceptList(T, S) R
	AcceptSet(T, S) R
	AcceptMap(T, S) R
}

func DispatchAccept3[T, S, U any, R DType](visitor ITypeVisitor3[T, S, U, R], ttype TType, t T, s S, u U) R {
	switch ttype.(type) {
	case mustEmbedTBool:
		return visitor.AcceptBool(t, s, u)
	case mustEmbedTByte:
		return visitor.AcceptByte(t, s, u)
	case mustEmbedTShort:
		return visitor.AcceptShort(t, s, u)
	case mustEmbedTInt:
		return visitor.AcceptInt(t, s, u)
	case mustEmbedTLong:
		return visitor.AcceptLong(t, s, u)
	case mustEmbedTFloat:
		return visitor.AcceptFloat(t, s, u)
	case mustEmbedTDouble:
		return visitor.AcceptDouble(t, s, u)
	case mustEmbedTEnum:
		return visitor.AcceptEnum(t, s, u)
	case mustEmbedTString:
		return visitor.AcceptString(t, s, u)
	case mustEmbedTDateTime:
		return visitor.AcceptDateTime(t, s, u)
	case mustEmbedTBean:
		return visitor.AcceptBean(t, s, u)
	case mustEmbedTArray:
		return visitor.AcceptArray(t, s, u)
	case mustEmbedTList:
		return visitor.AcceptList(t, s, u)
	case mustEmbedTSet:
		return visitor.AcceptSet(t, s, u)
	case mustEmbedTMap:
		return visitor.AcceptMap(t, s, u)
	default:
		panic("not support type: " + ttype.TypeName())
	}
}

type ITypeVisitor3[T, S, U any, R DType] interface {
	ITypeVisitor
	AcceptBool(T, S, U) R
	AcceptByte(T, S, U) R
	AcceptShort(T, S, U) R
	AcceptInt(T, S, U) R
	AcceptLong(T, S, U) R
	AcceptFloat(T, S, U) R
	AcceptDouble(T, S, U) R
	AcceptEnum(T, S, U) R
	AcceptString(T, S, U) R
	AcceptDateTime(T, S, U) R
	AcceptBean(T, S, U) R
	AcceptArray(T, S, U) R
	AcceptList(T, S, U) R
	AcceptSet(T, S, U) R
	AcceptMap(T, S, U) R
}

func DispatchAccept4[T, S, U, V any, R DType](visitor ITypeVisitor4[T, S, U, V, R], ttype TType, t T, s S, u U, v V) R {
	switch ttype.(type) {
	case mustEmbedTBool:
		return visitor.AcceptBool(t, s, u, v)
	case mustEmbedTByte:
		return visitor.AcceptByte(t, s, u, v)
	case mustEmbedTShort:
		return visitor.AcceptShort(t, s, u, v)
	case mustEmbedTInt:
		return visitor.AcceptInt(t, s, u, v)
	case mustEmbedTLong:
		return visitor.AcceptLong(t, s, u, v)
	case mustEmbedTFloat:
		return visitor.AcceptFloat(t, s, u, v)
	case mustEmbedTDouble:
		return visitor.AcceptDouble(t, s, u, v)
	case mustEmbedTEnum:
		return visitor.AcceptEnum(t, s, u, v)
	case mustEmbedTString:
		return visitor.AcceptString(t, s, u, v)
	case mustEmbedTDateTime:
		return visitor.AcceptDateTime(t, s, u, v)
	case mustEmbedTBean:
		return visitor.AcceptBean(t, s, u, v)
	case mustEmbedTArray:
		return visitor.AcceptArray(t, s, u, v)
	case mustEmbedTList:
		return visitor.AcceptList(t, s, u, v)
	case mustEmbedTSet:
		return visitor.AcceptSet(t, s, u, v)
	case mustEmbedTMap:
		return visitor.AcceptMap(t, s, u, v)
	default:
		panic("not support type: " + ttype.TypeName())
	}
}

type ITypeVisitor4[T, S, U, V any, R DType] interface {
	ITypeVisitor
	AcceptBool(T, S, U, V) R
	AcceptByte(T, S, U, V) R
	AcceptShort(T, S, U, V) R
	AcceptInt(T, S, U, V) R
	AcceptLong(T, S, U, V) R
	AcceptFloat(T, S, U, V) R
	AcceptDouble(T, S, U, V) R
	AcceptEnum(T, S, U, V) R
	AcceptString(T, S, U, V) R
	AcceptDateTime(T, S, U, V) R
	AcceptBean(T, S, U, V) R
	AcceptArray(T, S, U, V) R
	AcceptList(T, S, U, V) R
	AcceptSet(T, S, U, V) R
	AcceptMap(T, S, U, V) R
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
