package schema

import "github.com/zhangga/luban-go/core/rawdefs"

// ISchemaCollector 负责收集不同数据源解析出的 RawDefs 元数据
type ISchemaCollector interface {
	Load(config *LubanConfig) error
	CreateRawAssembly() *rawdefs.RawAssembly

	AddTable(table *rawdefs.RawTable)
	AddBean(bean *rawdefs.RawBean)
	AddEnum(enum *rawdefs.RawEnum)
	AddRefGroup(refGroup *rawdefs.RawRefGroup)
	AddConstAlias(name string, alias string)
}

// ISchemaLoader 定义了特定文件类型（如 xml，excel）架构的加载行为
type ISchemaLoader interface {
	Type() string
	Collector() ISchemaCollector
	Load(fileName string) error
}
