package pipeline

import (
	"context"
	"github.com/zhangga/luban/core/pipeline"
	"github.com/zhangga/luban/internal/defs"
	"github.com/zhangga/luban/internal/rawdefs"
)

type contextBuilder struct {
	assembly    *defs.DefAssembly
	includeTags []string
	excludeTags []string
	timeZone    string
}

func NewContextBuilder(assembly *defs.DefAssembly, includeTags, excludeTags []string, timeZone string) contextBuilder {
	return contextBuilder{
		assembly:    assembly,
		includeTags: includeTags,
		excludeTags: excludeTags,
		timeZone:    timeZone,
	}
}

func (b contextBuilder) Build(ctx context.Context) (*pContext, error) {
	pctx := &pContext{
		Context:     ctx,
		DefAssembly: b.assembly,
	}

	return pctx, nil
}

var _ pipeline.Context = (*pContext)(nil)

type pContext struct {
	context.Context
	*defs.DefAssembly
}

func (c *pContext) Target() *rawdefs.RawTarget {
	return c.DefAssembly.Target
}

func (c *pContext) TopModule() string {
	return c.Target().TopModule
}
