package codetarget

import (
	"fmt"
	"strings"
)

// NewPbTarget 创建一个 Protobuf 代码生成器
func NewPbTarget(templateDir string) (*TemplateCodeTarget, error) {
	t, err := NewTemplateCodeTarget("pb", "proto", templateDir, pbTypeMapper)
	if err != nil {
		return nil, err
	}
	
	// 为 pb 模板添加特殊函数，主要是为了给字段分配递增的 tag 序号
	// 原版 template_target 已经有个基础功能，但这里需要 add 函数
	t.tmpl.Funcs(map[string]interface{}{
		"add": func(a, b int) int {
			return a + b
		},
	})
	
	// 重新加载模板使其识别新的 Funcs
	if err := t.loadTemplates(); err != nil {
		return nil, fmt.Errorf("failed to reload pb templates: %v", err)
	}

	return t, nil
}

// pbTypeMapper 映射内部类型到 Protobuf (proto3) 类型
func pbTypeMapper(t string) string {
	t = strings.TrimSpace(t)
	switch t {
	case "bool":
		return "bool"
	case "byte", "short", "fshort", "int", "fint":
		return "int32"
	case "long", "flong":
		return "int64"
	case "float":
		return "float"
	case "double":
		return "double"
	case "string", "text":
		return "string"
	}
	
	// 处理集合类型，例如 list,int -> repeated int32
	if strings.HasPrefix(t, "list,") || strings.HasPrefix(t, "array,") || strings.HasPrefix(t, "set,") {
		// 取出元素类型
		parts := strings.SplitN(t, ",", 2)
		if len(parts) == 2 {
			elemType := strings.TrimSpace(parts[1])
			return "repeated " + pbTypeMapper(elemType)
		}
	}
	
	if strings.HasPrefix(t, "map,") {
		parts := strings.Split(strings.TrimPrefix(t, "map,"), ",")
		if len(parts) == 2 {
			kType := pbTypeMapper(strings.TrimSpace(parts[0]))
			vType := pbTypeMapper(strings.TrimSpace(parts[1]))
			return fmt.Sprintf("map<%s, %s>", kType, vType)
		}
	}

	// 默认假定为自定义类型（如枚举或 Bean）
	// Protobuf 不允许类型名包含 "."，如果是带 namespace 的可能需要替换成 "_"
	return strings.ReplaceAll(t, ".", "_")
}