package pipeline

import (
	"github.com/zhangga/luban/pkg/logger"
	"sync"
)

// Creator IPipeline构造器
type Creator func(logger logger.Logger) IPipeline

var (
	creators = make(map[string]Creator)
	locker   sync.RWMutex
)

func Register(creator Creator) {
	locker.Lock()
	defer locker.Unlock()
	t := creator(logger.Default())
	if _, ok := creators[t.Name()]; ok {
		panic("register duplicate IPipeline creator, name: " + t.Name())
	}
	creators[t.Name()] = creator
}

func GetCreator(name string) Creator {
	locker.RLock()
	defer locker.RUnlock()
	if creator, ok := creators[name]; ok {
		return creator
	}
	return nil
}
