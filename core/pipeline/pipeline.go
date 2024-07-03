package pipeline

import (
	"github.com/zhangga/luban/core/options"
	"github.com/zhangga/luban/core/pctx"
)

type IPipeline interface {
	Name() string          // pipeline名称
	Context() pctx.Context // 上下文
	Run(opts options.CommandOptions) error
}
