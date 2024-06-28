package manager

import (
	"fmt"
	"github.com/modern-go/reflect2"
	"github.com/zhangga/luban/pkg/logger"
	"reflect"
	"sync"
)

// IManager 所有管理器的接口
type IManager interface {
	key() uintptr

	Init(logger logger.Logger)
	PostInit()
}

var (
	// managers 注册的管理器
	managers = make(map[uintptr]IManager)
	locker   sync.RWMutex
)

func Register[T IManager]() {
	var t T
	typ := reflect.TypeOf(t).Elem()
	// 创建类型T的新实例
	instance := reflect.New(typ).Interface().(IManager)

	locker.Lock()
	defer locker.Unlock()
	if _, ok := managers[instance.key()]; ok {
		panic(fmt.Errorf("manager: %s already registered", typ.String()))
	}
	managers[instance.key()] = instance
}

func Deregister[T IManager]() {
	// 获取T的类型信息
	var t T
	typ := reflect.TypeOf(t).Elem()
	// 创建类型T的新实例
	instance := reflect.New(typ).Interface().(IManager)

	locker.Lock()
	defer locker.Unlock()
	delete(managers, instance.key())
}

func GetIface[T IManager]() (T, bool) {
	// 获取T的类型信息
	rtype := reflect2.RTypeOf((*T)(nil))

	locker.RLock()
	defer locker.RUnlock()
	if manager, ok := managers[rtype]; ok {
		return manager.(T), true
	}

	var instance T
	return instance, false
}

func MustGetIface[T IManager]() T {
	manager, ok := GetIface[T]()
	if !ok {
		panic(fmt.Sprintf("manager: %s not found", reflect.TypeOf(manager).Elem().String()))
	}
	return manager
}

func Traverse(f func(m IManager)) {
	locker.RLock()
	defer locker.RUnlock()
	for _, manager := range managers {
		f(manager)
	}
}
