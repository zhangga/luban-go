package schema

import (
	"github.com/zhangga/luban/core/pipeline"
	"github.com/zhangga/luban/pkg/logger"
	"sync"
)

// CollectorCreator ISchemaCollector构造器
type CollectorCreator func(logger logger.Logger, pipeline pipeline.IPipeline) ISchemaCollector

var (
	creators = make(map[string]CollectorCreator)
	locker   sync.RWMutex
)

func RegisterCollector(creator CollectorCreator) {
	locker.Lock()
	defer locker.Unlock()
	t := creator(logger.Default(), nil)
	if _, ok := creators[t.Name()]; ok {
		panic("register duplicate ISchemaCollector creator, name: " + t.Name())
	}
	creators[t.Name()] = creator
}

func GetCollectorCreator(name string) CollectorCreator {
	locker.RLock()
	defer locker.RUnlock()
	if creator, ok := creators[name]; ok {
		return creator
	}
	return nil
}
