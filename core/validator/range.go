package validator

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/zhangga/luban-go/core/datas"
)

// RangeValidator 处理数值大小和字符串长度的范围校验
type RangeValidator struct{}

func (v *RangeValidator) Name() string {
	return "range"
}

func (v *RangeValidator) Validate(ctx *ValidatorContext, data datas.DType, rule string) error {
	// rule example: "[1, 10]" or "(1, 10)"
	rule = strings.TrimSpace(rule)
	if rule == "" {
		return nil
	}

	leftInclude := strings.HasPrefix(rule, "[")
	rightInclude := strings.HasSuffix(rule, "]")
	
	ruleContent := rule[1 : len(rule)-1]
	parts := strings.Split(ruleContent, ",")
	if len(parts) != 2 {
		return fmt.Errorf("invalid range rule: %s", rule)
	}

	minStr, maxStr := strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1])

	// 针对不同数据类型进行范围校验
	switch d := data.(type) {
	case *datas.DInt:
		return checkIntRange(int64(d.Value), minStr, maxStr, leftInclude, rightInclude, rule)
	case *datas.DLong:
		return checkIntRange(d.Value, minStr, maxStr, leftInclude, rightInclude, rule)
	case *datas.DShort:
		return checkIntRange(int64(d.Value), minStr, maxStr, leftInclude, rightInclude, rule)
	case *datas.DFloat:
		return checkFloatRange(float64(d.Value), minStr, maxStr, leftInclude, rightInclude, rule)
	case *datas.DDouble:
		return checkFloatRange(d.Value, minStr, maxStr, leftInclude, rightInclude, rule)
	case *datas.DString:
		length := int64(len(d.Value))
		return checkIntRange(length, minStr, maxStr, leftInclude, rightInclude, rule)
	}

	return nil
}

func checkIntRange(val int64, minStr, maxStr string, leftInc, rightInc bool, rule string) error {
	if minStr != "" {
		min, err := strconv.ParseInt(minStr, 10, 64)
		if err == nil {
			if (leftInc && val < min) || (!leftInc && val <= min) {
				return fmt.Errorf("value %d is out of range, violated rule %s", val, rule)
			}
		}
	}
	if maxStr != "" {
		max, err := strconv.ParseInt(maxStr, 10, 64)
		if err == nil {
			if (rightInc && val > max) || (!rightInc && val >= max) {
				return fmt.Errorf("value %d is out of range, violated rule %s", val, rule)
			}
		}
	}
	return nil
}

func checkFloatRange(val float64, minStr, maxStr string, leftInc, rightInc bool, rule string) error {
	if minStr != "" {
		min, err := strconv.ParseFloat(minStr, 64)
		if err == nil {
			if (leftInc && val < min) || (!leftInc && val <= min) {
				return fmt.Errorf("value %f is out of range, violated rule %s", val, rule)
			}
		}
	}
	if maxStr != "" {
		max, err := strconv.ParseFloat(maxStr, 64)
		if err == nil {
			if (rightInc && val > max) || (!rightInc && val >= max) {
				return fmt.Errorf("value %f is out of range, violated rule %s", val, rule)
			}
		}
	}
	return nil
}
