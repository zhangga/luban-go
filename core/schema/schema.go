package schema

import "github.com/zhangga/luban/core/lubanconf"

type ISchemaCollector interface {
	Name() string
	Load(config *lubanconf.LubanConfig)
}

type ISchemaLoader interface {
	Load(fileName string)
}
