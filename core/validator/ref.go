package validator

import (
	"fmt"
	"strings"

	"github.com/zhangga/luban-go/core/datas"
	"github.com/zhangga/luban-go/core/defs"
)

// RefValidator 处理外键引用校验
type RefValidator struct{}

func (v *RefValidator) Name() string {
	return "ref"
}

func (v *RefValidator) Validate(ctx *ValidatorContext, data datas.DType, rule string) error {
	// rule example: "TbItem" (目标表名)
	targetTableName := strings.TrimSpace(rule)
	if targetTableName == "" {
		return nil
	}

	// 1. 查找目标表
	var targetTableFullName string
	for name, table := range ctx.Assembly.TablesByName {
		if name == targetTableName || table.Name() == targetTableName {
			targetTableFullName = table.FullName()
			break
		}
	}

	if targetTableFullName == "" {
		return fmt.Errorf("ref validation failed: target table %s not found", targetTableName)
	}

	// 2. 获取目标表的数据
	records, ok := ctx.TableRecordsMap[targetTableFullName]
	if !ok || len(records) == 0 {
		// 目标表没有数据，或者没有加载到数据，如果当前值不为默认空值则报错
		if !isDefaultOrEmpty(data) {
			return fmt.Errorf("ref validation failed: target table %s has no records but ref value is %v", targetTableName, data)
		}
		return nil
	}

	// 3. 构建目标表的主键集合 (这里假设我们已经知道目标表的主键是第一列，或者是某个特定字段)
	// 为了简化，我们遍历 records，提取每个 record 的主键值 (通常是第一列 Id)
	targetTableDef := ctx.Assembly.TablesByFullName[targetTableFullName]
	indexFieldName := targetTableDef.Raw.Index

	validKeys := make(map[string]bool)
	for _, rec := range records {
		if rec.Data != nil {
			dBean := rec.Data
			
			defBeanImpl, ok := dBean.ImplType.DefBean.(*defs.DefBeanImpl)
			if ok {
				// 找到主键字段对应的值
				for i, field := range defBeanImpl.HierarchyFields {
					if field.Raw.Name == indexFieldName {
						if i < len(dBean.Fields) {
							keyStr := fmt.Sprintf("%v", dBean.Fields[i])
							validKeys[keyStr] = true
						}
						break
					}
				}
			}
		}
	}

	// 4. 判断当前 data 的值是否在目标表的主键集合中
	valStr := fmt.Sprintf("%v", data)
	if _, exists := validKeys[valStr]; !exists {
		// 允许配置了 ref 但填空值的情况 (具体视业务而定，这里假设0或空字符串允许)
		if !isDefaultOrEmpty(data) {
			return fmt.Errorf("ref validation failed: value '%s' not found in table %s", valStr, targetTableName)
		}
	}

	return nil
}

func isDefaultOrEmpty(data datas.DType) bool {
	switch d := data.(type) {
	case *datas.DInt:
		return d.Value == 0
	case *datas.DLong:
		return d.Value == 0
	case *datas.DString:
		return d.Value == ""
	}
	return false
}
