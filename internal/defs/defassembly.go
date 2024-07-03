package defs

import (
	"fmt"
	"github.com/zhangga/luban/core/pctx"
	"github.com/zhangga/luban/core/refs"
	"github.com/zhangga/luban/internal/rawdefs"
	"github.com/zhangga/luban/internal/utils"
	"github.com/zhangga/luban/pkg/logger"
	"strings"
)

type DefAssembly struct {
	refs.EmbedDefAssembly
	logger                 logger.Logger
	TypeMap                map[string]refs.IDefType
	TypeList               []refs.IDefType
	notCaseSenseTypes      map[string]refs.IDefType // 不区分大小写的类型
	namespaces             map[string]struct{}      // 命名空间去重
	notCaseSenseNamespaces map[string]refs.IDefType // 不区分大小写的命名空间
	targets                []*rawdefs.RawTarget
	Target                 *rawdefs.RawTarget
	refGroups              map[string]*DefRefGroup
	TablesByName           map[string]*DefTable
	TablesByFullName       map[string]*DefTable
	exportTables           []*DefTable
}

func NewDefAssembly(ctx pctx.Context, logger logger.Logger, rawAssembly *rawdefs.RawAssembly, target string, outputTables []string) *DefAssembly {
	assembly := &DefAssembly{
		logger:  logger,
		targets: rawAssembly.Targets,
	}

	assembly.Target = assembly.GetTarget(target)
	if assembly.Target == nil {
		panic("target not found: " + target)
	}

	for _, g := range rawAssembly.RefGroups {
		assembly.AddRefGroup(g)
	}
	for _, e := range rawAssembly.Enums {
		assembly.AddType(NewDefEnum(e))
	}
	for _, b := range rawAssembly.Beans {
		assembly.AddType(NewDefBean(*b))
	}
	for _, t := range rawAssembly.Tables {
		table := NewDefTable(t)
		assembly.AddType(table)
		assembly.AddCfgTable(table)
	}

	if len(outputTables) == 0 {
		originTables := assembly.GetAllTables()
		for _, table := range originTables {
			if assembly.NeedExport(table.Groups) {
				assembly.exportTables = append(assembly.exportTables, table)
			}
		}
	} else {
		for _, tableName := range outputTables {
			if table := assembly.GetCfgTable(tableName); table != nil {
				assembly.exportTables = append(assembly.exportTables, table)
			} else {
				panic(fmt.Errorf("outputTable: %s, 未找到", tableName))
			}
		}
	}

	for _, table := range assembly.exportTables {
		table.IsExported = true
	}
	for _, t := range assembly.TypeList {
		t.SetAssembly(assembly)
	}
	for _, t := range assembly.TypeList {
		t.PreCompile(ctx)
	}
	for _, t := range assembly.TypeList {
		t.Compile(ctx)
	}
	for _, t := range assembly.TypeList {
		t.PostCompile(ctx)
	}
	return assembly
}

func (a *DefAssembly) AddRefGroup(group *rawdefs.RawRefGroup) {
	if _, ok := a.refGroups[group.Name]; ok {
		panic("duplicate ref group: " + group.Name)
	}
	a.refGroups[group.Name] = NewDefRefGroup(group)
}

func (a *DefAssembly) GetRefGroup(name string) *DefRefGroup {
	return a.refGroups[name]
}

func (a *DefAssembly) AddType(defType refs.IDefType) {
	fullName := defType.FullName()
	if _, ok := a.TypeMap[fullName]; ok {
		panic("duplicate type: " + fullName)
	}
	if exist, ok := a.notCaseSenseTypes[strings.ToLower(fullName)]; ok {
		panic(fmt.Errorf("type: %s, type: %s, 类名小写重复. 在win平台有问题", fullName, exist.FullName()))
	} else {
		a.notCaseSenseTypes[strings.ToLower(fullName)] = defType
	}

	namespace := defType.Namespace()
	if exist, ok := a.notCaseSenseNamespaces[strings.ToLower(namespace)]; ok {
		panic(fmt.Errorf("type: %s, type: %s, 命名空间小写重复. 在win平台有问题，请修改定义并删除生成的代码目录后再重新生成", fullName, exist.FullName()))
	} else {
		a.notCaseSenseNamespaces[strings.ToLower(namespace)] = defType
	}

	a.namespaces[namespace] = struct{}{}
	a.TypeMap[fullName] = defType
	a.TypeList = append(a.TypeList, defType)
}

func (a *DefAssembly) GetDefTypeByName(fullName string) refs.IDefType {
	return a.TypeMap[fullName]
}

func (a *DefAssembly) AddCfgTable(table *DefTable) {
	if _, ok := a.TablesByFullName[table.FullName()]; ok {
		panic("duplicate table: " + table.FullName())
	}
	a.TablesByFullName[table.FullName()] = table

	if exist, ok := a.TablesByName[table.Name]; ok {
		panic(fmt.Errorf("table: %s, table: %s, 的表名重复(不同模块下也不允许定义同名表，将来可能会放开限制)", table.FullName(), exist.FullName()))
	}
}

func (a *DefAssembly) GetCfgTable(fullName string) *DefTable {
	return a.TablesByFullName[fullName]
}

func (a *DefAssembly) GetAllTables() []*DefTable {
	tables := make([]*DefTable, 0, len(a.TypeList))
	for _, t := range a.TypeList {
		if table, ok := t.(*DefTable); ok {
			tables = append(tables, table)
		}
	}
	return tables
}

func (a *DefAssembly) GetTarget(name string) *rawdefs.RawTarget {
	for _, target := range a.targets {
		if target.Name == name {
			return target
		}
	}
	return nil
}

func (a *DefAssembly) NeedExport(groups []string) bool {
	if len(groups) == 0 {
		return true
	}
	for _, g := range groups {
		if utils.Contain[string](a.Target.Groups, g) {
			return true
		}
	}
	return false
}

func (a *DefAssembly) GetDefType(moduleName, typeName string) refs.IDefType {
	fullName := utils.MakeFullName(moduleName, typeName)
	if defType, ok := a.TypeMap[fullName]; ok {
		return defType
	}
	if defType, ok := a.TypeMap[typeName]; ok {
		return defType
	}
	return nil
}
