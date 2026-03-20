package codetarget

import (
	"strings"
)

// NewCppTarget 创建一个 C++ 代码生成器
func NewCppTarget(templateDir string) (*TemplateCodeTarget, error) {
	return NewTemplateCodeTarget("cpp", "h", templateDir, cppTypeMapper)
}

// cppTypeMapper 映射内部类型到 C++ 类型
func cppTypeMapper(t string) string {
	t = strings.TrimSpace(t)
	switch t {
	case "bool":
		return "bool"
	case "byte":
		return "uint8_t"
	case "short", "fshort":
		return "int16_t"
	case "int", "fint":
		return "int32_t"
	case "long", "flong":
		return "int64_t"
	case "float":
		return "float"
	case "double":
		return "double"
	case "string":
		return "std::string"
	case "text":
		return "std::string"
	}

	// 处理集合类型，例如 list,int
	if strings.HasPrefix(t, "list,") || strings.HasPrefix(t, "array,") {
		elemType := strings.TrimPrefix(t, "list,")
		if strings.HasPrefix(t, "array,") {
			elemType = strings.TrimPrefix(t, "array,")
		}
		return "std::vector<" + cppTypeMapper(elemType) + ">"
	}
	if strings.HasPrefix(t, "map,") {
		parts := strings.Split(strings.TrimPrefix(t, "map,"), ",")
		if len(parts) == 2 {
			return "std::map<" + cppTypeMapper(parts[0]) + ", " + cppTypeMapper(parts[1]) + ">"
		}
	}
	if strings.HasPrefix(t, "set,") {
		elemType := strings.TrimPrefix(t, "set,")
		return "std::set<" + cppTypeMapper(elemType) + ">"
	}

	// 默认假定为自定义类型（如枚举或 Bean）
	// 在 C++ 中，为了防止循环引用和多态，通常在容器中建议使用智能指针 std::shared_ptr<T>
	// 这里作为简单示例，假定传值。实际生产中可能需要根据是否为 Bean 来决定是否包 shared_ptr
	return t
}
