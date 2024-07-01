package pipeline

import "github.com/zhangga/luban/core/lubanconf"

type IPipeline interface {
	Name() string                   // pipeline名称
	Args() Arguments                // 启动参数
	Config() *lubanconf.LubanConfig // luban配置
	Context() Context               // 上下文
	Run(args Arguments) error
}
