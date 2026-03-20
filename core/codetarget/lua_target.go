package codetarget

import (
	"strings"
)

// NewLuaCodeTarget 创建一个 Lua 代码生成器
func NewLuaCodeTarget(templateDir string) (*TemplateCodeTarget, error) {
	return NewTemplateCodeTarget("lua", "lua", templateDir, luaTypeMapper)
}

// luaTypeMapper 映射内部类型到 Lua 注释/类型 (Lua 是动态类型语言，这里主要用于注释或类型提示声明)
func luaTypeMapper(t string) string {
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
	
	if strings.HasPrefix(t, "list,") || strings.HasPrefix(t, "array,") || strings.HasPrefix(t, "map,") || strings.HasPrefix(t, "set,") {
		return "table"
	}

	return t
}
