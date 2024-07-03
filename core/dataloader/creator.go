package dataloader

import (
	"github.com/zhangga/luban/pkg/logger"
	"sync"
)

// Creator IDataLoader构造器
type Creator func(logger logger.Logger) IDataLoader

var (
	creators = make(map[string]Creator)
	locker   sync.RWMutex
)

func Register(creator Creator) {
	locker.Lock()
	defer locker.Unlock()
	t := creator(logger.Default())
	for _, name := range t.Names() {
		if _, ok := creators[name]; ok {
			panic("register duplicate IDataLoader creator, name: " + name)
		}
		creators[name] = creator
	}
}

func GetCreator(name string) Creator {
	locker.RLock()
	defer locker.RUnlock()
	if creator, ok := creators[name]; ok {
		return creator
	}
	return nil
}
