package defs

import (
	"fmt"
	"github.com/zhangga/luban/core/schema"
	"github.com/zhangga/luban/internal/rawdefs"
	"github.com/zhangga/luban/internal/utils"
)

type IDefType interface {
	PreCompile()
	Compile()
	PostCompile()
}

// DefTypeBase 类型定义的基类
type DefTypeBase struct {
	Assembly    *DefAssembly
	Name        string
	Namespace   string
	Groups      []string
	Comment     string
	Tags        map[string]string
	TypeMappers []rawdefs.TypeMapper
}

func (t *DefTypeBase) FullName() string {
	return utils.MakeFullName(t.Namespace, t.Name)
}

func (t *DefTypeBase) NamespaceWithTopModule() string {

}

func (t *DefTypeBase) FullNameWithTopModule() string {

}

func (t *DefTypeBase) HasTag(attrName string) bool {
	if len(t.Tags) == 0 {
		return false
	}
	_, ok := t.Tags[attrName]
	return ok
}

func (t *DefTypeBase) GetTag(attrName string) (string, bool) {
	if len(t.Tags) == 0 {
		return "", false
	}
	v, ok := t.Tags[attrName]
	return v, ok
}

func (t *DefTypeBase) PreCompile(collector schema.ISchemaCollector) {
	if len(t.Groups) == 0 {
		return
	}

	config := collector.Pipeline().Config()
	if utils.Contain[string](t.Groups, "*") {
		t.Groups = t.Groups[:0]
		for _, g := range config.Groups {
			t.Groups = append(t.Groups, g.Names...)
		}
	} else {
		defGroups := map[string]struct{}{}
		for _, g := range config.Groups {
			for _, name := range g.Names {
				defGroups[name] = struct{}{}
			}
		}
		for _, g := range t.Groups {
			if _, ok := defGroups[g]; !ok {
				panic(fmt.Errorf("type: %s, group: %s not found", t.FullName(), g))
			}
		}
	}
}
