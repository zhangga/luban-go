package pctx

import (
	"github.com/zhangga/luban/core/lubanconf"
	"github.com/zhangga/luban/core/options"
)

type Context interface {
	TopModule() string
	Config() *lubanconf.LubanConfig
	Args() options.Arguments
}
