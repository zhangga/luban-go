package defs

import (
	"fmt"
	"github.com/zhangga/luban-go/core/rawdefs"
)

type DefAssemblyImpl struct {
	RawAssembly *rawdefs.RawAssembly
	Target      *rawdefs.RawTarget

	Types            map[string]DefTypeBase
	TypeList         []DefTypeBase
	TablesByName     map[string]*DefTable
	TablesByFullName map[string]*DefTable

	ExportTables []*DefTable
}

func NewDefAssemblyImpl(raw *rawdefs.RawAssembly, targetName string, outputTables []string) (*DefAssemblyImpl, error) {
	a := &DefAssemblyImpl{
		RawAssembly:      raw,
		Types:            make(map[string]DefTypeBase),
		TypeList:         make([]DefTypeBase, 0),
		TablesByName:     make(map[string]*DefTable),
		TablesByFullName: make(map[string]*DefTable),
		ExportTables:     make([]*DefTable, 0),
	}

	// 查找目标
	for _, t := range raw.Targets {
		if t.Name == targetName {
			a.Target = t
			break
		}
	}
	if a.Target == nil {
		// return nil, fmt.Errorf("target:%s is invalid", targetName)
		// 允许目标为空的默认逻辑
		a.Target = &rawdefs.RawTarget{Name: targetName}
	}

	for _, e := range raw.Enums {
		defEnum := NewDefEnumImpl(e)
		a.AddType(defEnum)
	}

	for _, b := range raw.Beans {
		defBean := NewDefBeanImpl(b)
		a.AddType(defBean)
	}

	// 编译 Bean，建立父子依赖关系和层级字段 (模拟 Compile 阶段)
	a.compileBeans()

	for _, tb := range raw.Tables {
		defTable := NewDefTable(tb)
		a.AddType(defTable)
		if err := a.AddCfgTable(defTable); err != nil {
			return nil, err
		}
	}

	// 初始化 ExportTables
	if len(outputTables) == 0 {
		a.ExportTables = a.GetAllTables()
	} else {
		for _, tableName := range outputTables {
			if t, ok := a.TablesByFullName[tableName]; ok {
				a.ExportTables = append(a.ExportTables, t)
			} else {
				return nil, fmt.Errorf("outputTable:%s not found", tableName)
			}
		}
	}

	// TODO: Compile, PostCompile 链接与验证

	return a, nil
}

func (a *DefAssemblyImpl) AddType(t DefTypeBase) {
	a.Types[t.FullName()] = t
	a.TypeList = append(a.TypeList, t)
}

func (a *DefAssemblyImpl) AddCfgTable(t *DefTable) error {
	if _, ok := a.TablesByFullName[t.FullName()]; ok {
		return fmt.Errorf("table:'%s' duplicated", t.FullName())
	}
	a.TablesByFullName[t.FullName()] = t

	if existing, ok := a.TablesByName[t.Name()]; ok {
		return fmt.Errorf("table:'%s' 与 table:'%s' 的表名重复", t.FullName(), existing.FullName())
	}
	a.TablesByName[t.Name()] = t

	return nil
}

func (a *DefAssemblyImpl) GetAllTables() []*DefTable {
	tables := make([]*DefTable, 0)
	for _, t := range a.TypeList {
		if table, ok := t.(*DefTable); ok {
			tables = append(tables, table)
		}
	}
	return tables
}

func (a *DefAssemblyImpl) GetType(fullName string) DefTypeBase {
	return a.Types[fullName]
}

func (a *DefAssemblyImpl) compileBeans() {
	// 1. 建立父子链接
	for _, t := range a.TypeList {
		if bean, ok := t.(*DefBeanImpl); ok {
			if bean.Parent != "" {
				parentFullName := bean.Parent
				if pType := a.GetType(parentFullName); pType != nil {
					if parentBean, ok := pType.(*DefBeanImpl); ok {
						bean.ParentDefType = parentBean
						parentBean.Children = append(parentBean.Children, bean)
						parentBean.IsAbstractType = true
						// fmt.Printf("DEBUG: Link Parent %s -> Child %s\n", parentBean.FullName(), bean.FullName())
					}
				} else {
					// 尝试跨命名空间查找
					for fullName, t := range a.Types {
						if fullName == parentFullName || t.Name() == parentFullName {
							if parentBean, ok := t.(*DefBeanImpl); ok {
								bean.ParentDefType = parentBean
								parentBean.Children = append(parentBean.Children, bean)
								parentBean.IsAbstractType = true
								// fmt.Printf("DEBUG: Link Parent %s -> Child %s (via fallback)\n", parentBean.FullName(), bean.FullName())
							}
							break
						}
					}
				}
			}
		}
	}

	// 2. 收集继承字段 (需要从树的根向下或者递归处理)
	for _, t := range a.TypeList {
		if bean, ok := t.(*DefBeanImpl); ok {
			a.buildHierarchyFields(bean)
		}
	}
}

func (a *DefAssemblyImpl) buildHierarchyFields(bean *DefBeanImpl) {
	if len(bean.HierarchyFields) > 0 || len(bean.Fields) == 0 && bean.Parent == "" {
		return // 已经构建过或空
	}

	if bean.ParentDefType != nil {
		a.buildHierarchyFields(bean.ParentDefType)
		bean.HierarchyFields = append(bean.HierarchyFields, bean.ParentDefType.HierarchyFields...)
	}

	bean.HierarchyFields = append(bean.HierarchyFields, bean.Fields...)
}
