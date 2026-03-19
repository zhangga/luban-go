package defs

// DefTypeBase 定义了所有配置定义对象的基础接口
type DefTypeBase interface {
	FullName() string
	Name() string
	Namespace() string
}
