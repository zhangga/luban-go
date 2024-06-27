package schema

import (
	"github.com/zhangga/luban/core/lubanconf"
	"github.com/zhangga/luban/core/refs"
)

type ISchemaCollector interface {
	Name() string
	Load(config *lubanconf.LubanConfig)
	AddTable(t refs.UnimplementedTable)
	AddBean(b refs.UnimplementedBean)
	AddEnum(e refs.UnimplementedEnum)
	AddRefGroup(g refs.UnimplementedRefGroup)
}

type ISchemaLoader interface {
	Load(fileName string)
}

type IBeanSchemaLoader interface {
	Name() string
	Load(fileName, beanFullName string) refs.UnimplementedBean
}
