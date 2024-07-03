package pipeline

import (
	"context"
	"github.com/zhangga/luban/core/lubanconf"
	"github.com/zhangga/luban/core/options"
	"github.com/zhangga/luban/core/pctx"
	"github.com/zhangga/luban/internal/defs"
	"github.com/zhangga/luban/internal/rawdefs"
)

var _ pctx.Context = (*pContext)(nil)

type pContext struct {
	context.Context
	*defs.DefAssembly
	config *lubanconf.LubanConfig
	args   options.Arguments
}

func (c *pContext) Config() *lubanconf.LubanConfig {
	return c.config
}

func (c *pContext) Args() options.Arguments {
	return c.args
}

func (c *pContext) Target() *rawdefs.RawTarget {
	return c.DefAssembly.Target
}

func (c *pContext) TopModule() string {
	return c.Target().TopModule
}
