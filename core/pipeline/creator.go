package pipeline

import "sync"

// Creator IPipeline构造器
type Creator func() IPipeline

使用type注册，参考manager
var (
	creators = make(map[string]Creator)
	locker   sync.RWMutex
)

func Register(creator Creator) {
	locker.Lock()
	defer locker.Unlock()
	t := creator()
	if _, ok := creators[t.Name()]; ok {
		panic("register duplicate pipeline creator, name: " + t.Name())
	}
	creators[t.Name()] = creator
}

func getCreator(name string) Creator {
	locker.RLock()
	defer locker.RUnlock()
	if creator, ok := creators[name]; ok {
		return creator
	}
	return nil
}
