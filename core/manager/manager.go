package manager

import (
	"fmt"
	"github.com/modern-go/reflect2"
	"github.com/zhangga/luban/pkg/logger"
	"reflect"
	"sync"
)

type IManager interface {
	Init(logger logger.Logger)
	PostInit()
}

var (
	managers = make(map[uintptr]IManager)
	locker   sync.RWMutex
)

func Register[T IManager]() {
	// 获取T的类型信息
	rtype := reflect2.RTypeOf((*T)(nil))
	var t T
	typ := reflect.TypeOf(t).Elem()
	// 创建类型T的新实例
	instance := reflect.New(typ).Interface().(IManager)

	locker.Lock()
	defer locker.Unlock()
	if _, ok := managers[rtype]; ok {
		panic(fmt.Errorf("manager: %s already registered", typ.String()))
	}
	managers[rtype] = instance
}

func Deregister[T IManager]() {
	// 获取T的类型信息
	rtype := reflect2.RTypeOf((*T)(nil))

	locker.Lock()
	defer locker.Unlock()
	delete(managers, rtype)
}

func Get[T IManager]() (T, bool) {
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

func Traverse(f func(m IManager)) {
	locker.RLock()
	defer locker.RUnlock()
	for _, manager := range managers {
		f(manager)
	}
}
