package schema

import (
	"github.com/zhangga/luban/pkg/logger"
	"sync"
)

// BeanLoaderCreator IBeanSchemaLoader 构造器
type BeanLoaderCreator func(logger logger.Logger, collector ISchemaCollector) IBeanSchemaLoader

var (
	beanLoaderCreators = make(map[string]BeanLoaderCreator)
	eanLoaderLocker    sync.RWMutex
)

func RegisterBeanLoaderCreator(creator BeanLoaderCreator) {
	eanLoaderLocker.Lock()
	defer eanLoaderLocker.Unlock()
	t := creator(logger.Default(), nil)
	if _, ok := beanLoaderCreators[t.Name()]; ok {
		panic("register duplicate IBeanSchemaLoader creator, name: " + t.Name())
	}
	beanLoaderCreators[t.Name()] = creator
}

func getBeanLoaderCreator(name string) BeanLoaderCreator {
	eanLoaderLocker.RLock()
	defer eanLoaderLocker.RUnlock()
	if creator, ok := beanLoaderCreators[name]; ok {
		return creator
	}
	return nil
}
