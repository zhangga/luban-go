package codetarget

import (
	"strings"
)

// NewCSTarget 创建一个 C# 代码生成器
func NewCSTarget(templateDir string) (*TemplateCodeTarget, error) {
	return NewTemplateCodeTarget("cs", "cs", templateDir, csTypeMapper)
}

// csTypeMapper 映射内部类型到 C# 类型
func csTypeMapper(t string) string {
	t = strings.TrimSpace(t)
	switch t {
	case "bool":
		return "bool"
	case "byte":
		return "byte"
	case "short":
		return "short"
	case "fshort":
		return "short"
	case "int":
		return "int"
	case "fint":
		return "int"
	case "long":
		return "long"
	case "flong":
		return "long"
	case "float":
		return "float"
	case "double":
		return "double"
	case "string":
		return "string"
	}

	// 处理集合类型，例如 list,int
	if strings.HasPrefix(t, "list,") {
		elemType := strings.TrimPrefix(t, "list,")
		return "System.Collections.Generic.List<" + csTypeMapper(elemType) + ">"
	}
	if strings.HasPrefix(t, "array,") {
		elemType := strings.TrimPrefix(t, "array,")
		return csTypeMapper(elemType) + "[]"
	}
	if strings.HasPrefix(t, "map,") {
		parts := strings.Split(strings.TrimPrefix(t, "map,"), ",")
		if len(parts) == 2 {
			return "System.Collections.Generic.Dictionary<" + csTypeMapper(parts[0]) + ", " + csTypeMapper(parts[1]) + ">"
		}
	}
	if strings.HasPrefix(t, "set,") {
		elemType := strings.TrimPrefix(t, "set,")
		return "System.Collections.Generic.HashSet<" + csTypeMapper(elemType) + ">"
	}

	// 默认假定为自定义类型（如枚举或 Bean）
	return t
}
