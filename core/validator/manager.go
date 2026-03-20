package validator

import (
	"github.com/zhangga/luban-go/core/datas"
	"github.com/zhangga/luban-go/core/defs"
)

// ValidatorContext 提供校验时的上下文信息，包含整个装配体和所有记录数据
type ValidatorContext struct {
	Assembly        *defs.DefAssemblyImpl
	TableRecordsMap map[string][]*defs.Record
}

// NewValidatorContext 创建校验上下文
func NewValidatorContext(assembly *defs.DefAssemblyImpl, tableRecordsMap map[string][]*defs.Record) *ValidatorContext {
	return &ValidatorContext{
		Assembly:        assembly,
		TableRecordsMap: tableRecordsMap,
	}
}

// IValidator 定义了校验器的接口
type IValidator interface {
	Validate(ctx *ValidatorContext, data datas.DType, rule string) error
	Name() string
}

// DataValidatorManager 负责管理和执行所有校验器
type DataValidatorManager struct {
	validators map[string]IValidator
}

// NewDataValidatorManager 创建校验管理器
func NewDataValidatorManager() *DataValidatorManager {
	return &DataValidatorManager{
		validators: make(map[string]IValidator),
	}
}

// Register 注册校验器
func (m *DataValidatorManager) Register(v IValidator) {
	m.validators[v.Name()] = v
}

// Validate 执行校验逻辑
func (m *DataValidatorManager) Validate(ctx *ValidatorContext) []error {
	var errs []error
	
	// 遍历所有表和数据
	for _, table := range ctx.Assembly.ExportTables {
		records := ctx.TableRecordsMap[table.FullName()]
		for _, record := range records {
			// record.Data 已经是 *datas.DBean
			if record.Data != nil {
				beanErrs := m.validateBean(ctx, record.Data)
				errs = append(errs, beanErrs...)
			}
		}
	}
	return errs
}

func (m *DataValidatorManager) validateBean(ctx *ValidatorContext, bean *datas.DBean) []error {
	var errs []error
	
	// 在 defs 中查找具体的实现类型，以此来获取 HierarchyFields
	defBeanImpl, ok := bean.ImplType.DefBean.(*defs.DefBeanImpl)
	if !ok {
		return errs
	}

	fields := defBeanImpl.HierarchyFields
	if len(bean.Fields) != len(fields) {
		return errs // 结构不匹配则跳过，在其他地方处理
	}

	for i, field := range fields {
		dType := bean.Fields[i]
		
		// 检查字段上是否有 tags (校验规则)
		// 例如 ref="TbItem" 或 range="[1,100]"
		if len(field.Raw.Tags) > 0 {
			for tagName, tagValue := range field.Raw.Tags {
				if validator, ok := m.validators[tagName]; ok {
					if err := validator.Validate(ctx, dType, tagValue); err != nil {
						errs = append(errs, err)
					}
				}
			}
		}
		
		// 如果字段本身是 Bean，递归校验
		if childBean, ok := dType.(*datas.DBean); ok {
			errs = append(errs, m.validateBean(ctx, childBean)...)
		}
	}

	return errs
}
