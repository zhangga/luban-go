package defs

import "github.com/zhangga/luban-go/core/rawdefs"

// DefAssembly 聚合了所有的 Def 定义对象
type DefAssembly struct {
	RawAssembly *rawdefs.RawAssembly
	Types       map[string]DefTypeBase
	Tables      map[string]*DefTable
	// ... 其他 Def 定义
}

func NewDefAssembly(raw *rawdefs.RawAssembly) *DefAssembly {
	return &DefAssembly{
		RawAssembly: raw,
		Types:       make(map[string]DefTypeBase),
		Tables:      make(map[string]*DefTable),
	}
}

func (a *DefAssembly) AddType(t DefTypeBase) {
	a.Types[t.FullName()] = t
}

func (a *DefAssembly) AddTable(t *DefTable) {
	a.Tables[t.FullName()] = t
}
