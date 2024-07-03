package schema

import (
	"github.com/zhangga/luban/core/pctx"
	"github.com/zhangga/luban/core/pipeline"
	"github.com/zhangga/luban/core/refs"
)

type ISchemaCollector interface {
	Name() string
	Pipeline() pipeline.IPipeline
	Load(ctx pctx.Context)

	AddTable(t refs.UnimplementedTable)
	AddBean(b refs.UnimplementedBean)
	AddEnum(e refs.UnimplementedEnum)
	AddRefGroup(g refs.UnimplementedRefGroup)
}

type ISchemaLoader interface {
	Load(ctx pctx.Context, fileName string)
}

type IBeanSchemaLoader interface {
	Name() string
	Load(fileName, beanFullName string) refs.UnimplementedBean
}
