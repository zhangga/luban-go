package defs

import (
	"fmt"
	"strings"

	"github.com/zhangga/luban-go/core/types"
)

// TTypeFactory 模拟一个根据字符串创建 TType 的工厂
type TTypeFactory struct {
	assembly *DefAssemblyImpl
}

func NewTTypeFactory(assembly *DefAssemblyImpl) *TTypeFactory {
	return &TTypeFactory{assembly: assembly}
}

// CreateType 解析类型字符串，如 "int", "string", "array,int", "map,int,string" 或自定义类型 "demo.Item"
func (f *TTypeFactory) CreateType(typeStr string) (types.TType, error) {
	parts := strings.Split(typeStr, ",")
	for i := range parts {
		parts[i] = strings.TrimSpace(parts[i])
	}

	if len(parts) == 0 {
		return nil, fmt.Errorf("empty type string")
	}

	mainType := parts[0]

	switch mainType {
	case "bool":
		return types.NewTBool(false, nil), nil
	case "byte":
		return types.NewTByte(false, nil), nil
	case "short":
		return types.NewTShort(false, nil), nil
	case "int":
		return types.NewTInt(false, nil), nil
	case "long":
		return types.NewTLong(false, nil), nil
	case "float":
		return types.NewTFloat(false, nil), nil
	case "double":
		return types.NewTDouble(false, nil), nil
	case "string":
		return types.NewTString(false, nil), nil
	case "datetime":
		return types.NewTDateTime(false, nil), nil
	case "array":
		if len(parts) < 2 {
			return nil, fmt.Errorf("array missing element type")
		}
		elemType, err := f.CreateType(parts[1])
		if err != nil {
			return nil, err
		}
		return types.NewTArray(false, elemType, nil), nil
	case "list":
		if len(parts) < 2 {
			return nil, fmt.Errorf("list missing element type")
		}
		elemType, err := f.CreateType(parts[1])
		if err != nil {
			return nil, err
		}
		return types.NewTList(false, elemType, nil), nil
	case "set":
		if len(parts) < 2 {
			return nil, fmt.Errorf("set missing element type")
		}
		elemType, err := f.CreateType(parts[1])
		if err != nil {
			return nil, err
		}
		return types.NewTSet(false, elemType, nil), nil
	case "map":
		if len(parts) < 3 {
			return nil, fmt.Errorf("map missing key or value type")
		}
		keyType, err := f.CreateType(parts[1])
		if err != nil {
			return nil, err
		}
		valType, err := f.CreateType(parts[2])
		if err != nil {
			return nil, err
		}
		return types.NewTMap(false, keyType, valType, nil), nil
	default:
		// 尝试在 Assembly 中查找自定义类型（Bean或Enum）
		// 如果 mainType 不是全名，可能需要尝试加上当前的 namespace（或者 Assembly 中的 namespace）
		defType := f.assembly.GetType(mainType)
		if defType == nil {
			// 在测试例子里，类名为 demo.Item, 但传进来可能是 Item
			// 尝试模糊匹配或在所有类型里找以 .mainType 结尾的
			for fullName, t := range f.assembly.Types {
				if fullName == mainType || strings.HasSuffix(fullName, "."+mainType) {
					defType = t
					break
				}
			}
		}

		if defType != nil {
			switch dt := defType.(type) {
			case *DefBeanImpl:
				return types.NewTBean(false, dt, nil), nil
			case *DefEnumImpl:
				return types.NewTEnum(false, dt, nil), nil
			}
		}
		return nil, fmt.Errorf("unknown type: %s", mainType)
	}
}
