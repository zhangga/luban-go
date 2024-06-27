package schema

import (
	"github.com/zhangga/luban/pkg/logger"
	"sync"
)

type LoaderCreator func(logger logger.Logger, dataType string, collector ISchemaCollector) ISchemaLoader

type LoaderInfo struct {
	DataType string
	ExtName  string
	Priority int
	Creator  LoaderCreator
}

var (
	loaderInfos  = make(map[loaderKey]*LoaderInfo)
	loaderLocker sync.RWMutex
)

func RegisterSchemaLoader(creator LoaderCreator, priority int, dataType string, extNames ...string) {
	loaderLocker.Lock()
	defer loaderLocker.Unlock()
	for _, extName := range extNames {
		key := getLoaderKey(dataType, extName)
		if oldInfo, ok := loaderInfos[key]; ok {
			if oldInfo.Priority >= priority {
				logger.Errorf("register ISchemaLoader creator priority is lower, extName: %s, dataType: %s", extName, dataType)
				continue
			}
		}
		logger.Debugf("register ISchemaLoader creator, extName: %s, dataType: %s, priority: %d", extName, dataType, priority)
		loaderInfos[key] = &LoaderInfo{DataType: dataType, ExtName: extName, Priority: priority, Creator: creator}
	}
}

func getSchemaLoaderInfo(dataType, extName string) *LoaderInfo {
	locker.RLock()
	defer locker.RUnlock()
	key := getLoaderKey(dataType, extName)
	if info, ok := loaderInfos[key]; ok {
		return info
	}
	return nil
}

type loaderKey struct {
	dataType string
	extName  string
}

func getLoaderKey(dataType, extName string) loaderKey {
	return loaderKey{dataType: dataType, extName: extName}
}
