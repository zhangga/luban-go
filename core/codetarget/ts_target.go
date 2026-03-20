package codetarget

import (
	"strings"
)

// NewTSTarget 创建一个 TypeScript 代码生成器
func NewTSTarget(templateDir string) (*TemplateCodeTarget, error) {
	return NewTemplateCodeTarget("ts", "ts", templateDir, tsTypeMapper)
}

// tsTypeMapper 映射内部类型到 TypeScript 类型
func tsTypeMapper(t string) string {
	t = strings.TrimSpace(t)
	switch t {
	case "bool":
		return "boolean"
	case "byte", "short", "fshort", "int", "fint", "long", "flong", "float", "double":
		return "number"
	case "string":
		return "string"
	case "text":
		return "string"
	}
	
	// 处理集合类型，例如 list,int
	if strings.HasPrefix(t, "list,") {
		elemType := strings.TrimPrefix(t, "list,")
		return tsTypeMapper(elemType) + "[]"
	}
	if strings.HasPrefix(t, "array,") {
		elemType := strings.TrimPrefix(t, "array,")
		return tsTypeMapper(elemType) + "[]"
	}
	if strings.HasPrefix(t, "map,") {
		parts := strings.Split(strings.TrimPrefix(t, "map,"), ",")
		if len(parts) == 2 {
			return "Map<" + tsTypeMapper(parts[0]) + ", " + tsTypeMapper(parts[1]) + ">"
		}
	}
	if strings.HasPrefix(t, "set,") {
		elemType := strings.TrimPrefix(t, "set,")
		return "Set<" + tsTypeMapper(elemType) + ">"
	}

	// 默认假定为自定义类型（如枚举或 Bean）
	return t
}
