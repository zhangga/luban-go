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

	// 编译 Bean，建立父子依赖关系和层级字段
	if err := a.compileBeans(); err != nil {
		return nil, err
	}

	for _, tb := range raw.Tables {
		defTable := NewDefTable(tb)
		a.AddType(defTable)
		if err := a.AddCfgTable(defTable); err != nil {
			return nil, err
		}
	}

	// 执行完整校验 (Compile Phase)
	if err := a.compileAll(); err != nil {
		return nil, err
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

func (a *DefAssemblyImpl) compileBeans() error {
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
	return nil
}

func (a *DefAssemblyImpl) compileAll() error {
	// 先校验所有的 Bean (字段重名等)
	for _, t := range a.TypeList {
		if bean, ok := t.(*DefBeanImpl); ok {
			if err := a.validateBean(bean); err != nil {
				return err
			}
		}
	}

	for _, t := range a.TypeList {
		switch def := t.(type) {
		case *DefEnumImpl:
			if err := a.validateEnum(def); err != nil {
				return err
			}
		case *DefTable:
			if err := a.validateTable(def); err != nil {
				return err
			}
		}
	}
	return nil
}

func (a *DefAssemblyImpl) validateEnum(enum *DefEnumImpl) error {
	if len(enum.Raw.Items) == 0 {
		return fmt.Errorf("enum '%s' has no items", enum.FullName())
	}

	itemNames := make(map[string]bool)
	itemValues := make(map[string]bool)

	for _, item := range enum.Raw.Items {
		if itemNames[item.Name] {
			return fmt.Errorf("enum '%s' has duplicated item name: '%s'", enum.FullName(), item.Name)
		}
		itemNames[item.Name] = true

		// 检查别名 (如果存在且不等于名称)
		if item.Alias != "" && item.Alias != item.Name {
			if itemNames[item.Alias] {
				return fmt.Errorf("enum '%s' item '%s' alias '%s' conflicts with existing items", enum.FullName(), item.Name, item.Alias)
			}
			itemNames[item.Alias] = true
		}

		if item.Value != "" {
			if itemValues[item.Value] {
				// 在某些原版规则中允许相同值，但通常定义应唯一。这里为了严谨可给个警告，我们暂定报错
				return fmt.Errorf("enum '%s' has duplicated item value: %s (item: %s)", enum.FullName(), item.Value, item.Name)
			}
			itemValues[item.Value] = true
		}
	}

	return nil
}

func (a *DefAssemblyImpl) validateTable(table *DefTable) error {
	// 检查 ValueType 是否存在且必须是 Bean
	valType := a.GetType(table.Raw.ValueType)
	if valType == nil {
		return fmt.Errorf("table '%s' value type '%s' not found", table.FullName(), table.Raw.ValueType)
	}

	bean, ok := valType.(*DefBeanImpl)
	if !ok {
		return fmt.Errorf("table '%s' value type '%s' is not a Bean", table.FullName(), table.Raw.ValueType)
	}

	// 如果指定了 index，检查它是否存在于 Bean 的 HierarchyFields 中
	if table.Raw.Index != "" {
		found := false
		for _, f := range bean.HierarchyFields {
			if f.Raw.Name == table.Raw.Index {
				found = true
				break
			}
		}
		if !found {
			return fmt.Errorf("table '%s' index field '%s' not found in Bean '%s'", table.FullName(), table.Raw.Index, bean.FullName())
		}
	} else if len(bean.HierarchyFields) == 0 {
		return fmt.Errorf("table '%s' value type Bean '%s' has no fields, cannot be used as a table record", table.FullName(), bean.FullName())
	}

	return nil
}

func (a *DefAssemblyImpl) validateBean(bean *DefBeanImpl) error {
	// 1. 检查字段重名 (包含继承的字段)
	fieldNames := make(map[string]bool)
	for _, f := range bean.HierarchyFields {
		if fieldNames[f.Raw.Name] {
			return fmt.Errorf("bean '%s' has duplicated field name: '%s'", bean.FullName(), f.Raw.Name)
		}
		fieldNames[f.Raw.Name] = true
	}

	// 2. 如果是多态根节点，确保没有普通实例被错误标记为多态等
	// (更复杂的原版校验可逐步在这里补全)
	return nil
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
