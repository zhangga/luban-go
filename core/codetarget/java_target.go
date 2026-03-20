package codetarget

import (
	"strings"
)

// NewJavaTarget 创建一个 Java 代码生成器
func NewJavaTarget(templateDir string) (*TemplateCodeTarget, error) {
	return NewTemplateCodeTarget("java", "java", templateDir, javaTypeMapper)
}

// javaTypeMapper 映射内部类型到 Java 类型
func javaTypeMapper(t string) string {
	t = strings.TrimSpace(t)
	switch t {
	case "bool":
		return "boolean"
	case "byte":
		return "byte"
	case "short", "fshort":
		return "short"
	case "int", "fint":
		return "int"
	case "long", "flong":
		return "long"
	case "float":
		return "float"
	case "double":
		return "double"
	case "string":
		return "String"
	case "text":
		return "String"
	}

	// 泛型装箱类型（当作为集合元素时需要，这里简单起见，统一返回包装类如果在集合里，或者直接返回）
	// 为简单实现，在容器里暂时直接依赖 Java 的自动类型转换，如果需要精确匹配，需写一个辅助的 boxedTypeMapper
	// 处理集合类型，例如 list,int -> java.util.List<Integer>
	if strings.HasPrefix(t, "list,") || strings.HasPrefix(t, "array,") {
		elemType := strings.TrimPrefix(t, "list,")
		if strings.HasPrefix(t, "array,") {
			elemType = strings.TrimPrefix(t, "array,")
		}
		return "java.util.List<" + javaBoxedTypeMapper(elemType) + ">"
	}
	if strings.HasPrefix(t, "map,") {
		parts := strings.Split(strings.TrimPrefix(t, "map,"), ",")
		if len(parts) == 2 {
			return "java.util.Map<" + javaBoxedTypeMapper(parts[0]) + ", " + javaBoxedTypeMapper(parts[1]) + ">"
		}
	}
	if strings.HasPrefix(t, "set,") {
		elemType := strings.TrimPrefix(t, "set,")
		return "java.util.Set<" + javaBoxedTypeMapper(elemType) + ">"
	}

	// 默认假定为自定义类型（如枚举或 Bean）
	return t
}

func javaBoxedTypeMapper(t string) string {
	t = strings.TrimSpace(t)
	switch t {
	case "bool":
		return "Boolean"
	case "byte":
		return "Byte"
	case "short", "fshort":
		return "Short"
	case "int", "fint":
		return "Integer"
	case "long", "flong":
		return "Long"
	case "float":
		return "Float"
	case "double":
		return "Double"
	case "string":
		return "String"
	}
	return t
}
