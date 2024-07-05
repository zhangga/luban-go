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
	cacheDefTTypes         map[cacheDefKey]refs.TType
}

type cacheDefKey struct {
	def      refs.IDefType
	nullable bool
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

func (a *DefAssembly) GetTType(moduleName, typeName string, nullable bool, tags map[string]string) refs.TType {
	defType := a.GetDefType(moduleName, typeName)
	switch d := defType.(type) {
	case *DefBean:
		return a.GetOrCreateTBean(d, nullable, tags)
	case *DefEnum:
		return a.GetOrCreateTEnum(d, nullable, tags)
	}
	return nil
}

func (a *DefAssembly) GetOrCreateTBean(bean *DefBean, nullable bool, tags map[string]string) refs.TType {
	if len(tags) == 0 {
		if t, ok := a.cacheDefTTypes[cacheDefKey{def: bean, nullable: nullable}]; ok {
			return t
		}
		tbean := refs.GetTypeCreator("bean")(nullable, tags, bean)
		a.cacheDefTTypes[cacheDefKey{def: bean, nullable: nullable}] = tbean
		return tbean
	} else {
		refs.GetTypeCreator("bean")(nullable, tags, bean)
	}
	return nil
}

func (a *DefAssembly) GetOrCreateTEnum(enum *DefEnum, nullable bool, tags map[string]string) refs.TType {
	if len(tags) == 0 {
		if t, ok := a.cacheDefTTypes[cacheDefKey{def: enum, nullable: nullable}]; ok {
			return t
		}
		tenum := refs.GetTypeCreator("enum")(nullable, tags, enum)
		a.cacheDefTTypes[cacheDefKey{def: enum, nullable: nullable}] = tenum
		return tenum
	} else {
		refs.GetTypeCreator("enum")(nullable, tags, enum)
	}
	return nil
}

func (a *DefAssembly) CreateType(module, name string, containerElementType bool) refs.TType {
	name = utils.TrimBracePairs(name)
	if sepIndex := utils.IndexOfBaseTypeEnd(name); sepIndex > 0 {
		containerTypeAndTags := utils.TrimBracePairs(name[:sepIndex])
		elementTypeAndTags := strings.TrimSpace(name[sepIndex+1:])
		containerType, containerTags := utils.ParseTypeAndValidAttrs(containerTypeAndTags)
		return a.CreateContainerType(module, containerType, containerTags, elementTypeAndTags)
	} else {
		return a.CreateNotContainerType(module, name, containerElementType)
	}
}

// CreateNotContainerType 创建非容器类型
func (a *DefAssembly) CreateNotContainerType(module, name string, containerElementType bool) refs.TType {
	defaultable, nullable := true, false
	// 去掉rawType两侧匹配的()
	rawType := utils.TrimBracePairs(name)
	typ, tags := utils.ParseTypeAndValidAttrs(rawType)
	for {
		if strings.HasSuffix(typ, "?") {
			if containerElementType {
				panic(fmt.Errorf("container element type can't be nullable type: %s.%s", module, typ))
			}
			nullable = true
			typ = typ[:len(typ)-1]
			continue
		}
		if strings.HasSuffix(typ, "!") {
			defaultable = false
			typ = typ[:len(typ)-1]
			continue
		}
		break
	}

	if !defaultable {
		if _, ok := tags["not-default"]; !ok {
			tags["not-default"] = "1"
		}
	}

	switch typ {
	case "bool":
		return refs.GetTypeCreator("bool")(nullable, tags, nil)
	case "int8", "byte":
		return refs.GetTypeCreator("byte")(nullable, tags, nil)
	case "int16", "short":
		return refs.GetTypeCreator("short")(nullable, tags, nil)
	case "int32", "int":
		return refs.GetTypeCreator("int")(nullable, tags, nil)
	case "int64", "long":
		return refs.GetTypeCreator("long")(nullable, tags, nil)
	case "float32", "float":
		return refs.GetTypeCreator("float")(nullable, tags, nil)
	case "float64", "double":
		return refs.GetTypeCreator("double")(nullable, tags, nil)
	case "string":
		return refs.GetTypeCreator("string")(nullable, tags, nil)
	case "text":
		tags["text"] = "1"
		return refs.GetTypeCreator("string")(nullable, tags, nil)
	case "time", "datetime":
		return refs.GetTypeCreator("datetime")(nullable, tags, nil)
	default:
		if dtype := a.GetTType(module, typ, nullable, tags); dtype != nil {
			return dtype
		} else {
			panic(fmt.Errorf("invalid type, module: %s, type: %s", module, typ))
		}
	}
}

// CreateContainerType 创建容器类型
func (a *DefAssembly) CreateContainerType(module, containerType string, containerTags map[string]string, elementType string) refs.TType {
	panic("implement me")
}
