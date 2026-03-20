package datatarget

import (
	"fmt"
	"strings"

	"github.com/zhangga/luban-go/core/datas"
	"github.com/zhangga/luban-go/core/defs"
)

// ToLuaVisitor 是一个实现 IDataVisitor 接口的访问者，将 DType 转换为 Lua Table 字符串
type ToLuaVisitor struct {
	TargetGroups []string
}

func NewToLuaVisitor(targetGroups []string) *ToLuaVisitor {
	return &ToLuaVisitor{
		TargetGroups: targetGroups,
	}
}

func (v *ToLuaVisitor) VisitDBool(d *datas.DBool) interface{} {
	if d.Value {
		return "true"
	}
	return "false"
}

func (v *ToLuaVisitor) VisitDByte(d *datas.DByte) interface{} {
	return fmt.Sprintf("%d", d.Value)
}

func (v *ToLuaVisitor) VisitDShort(d *datas.DShort) interface{} {
	return fmt.Sprintf("%d", d.Value)
}

func (v *ToLuaVisitor) VisitDInt(d *datas.DInt) interface{} {
	return fmt.Sprintf("%d", d.Value)
}

func (v *ToLuaVisitor) VisitDLong(d *datas.DLong) interface{} {
	return fmt.Sprintf("%d", d.Value)
}

func (v *ToLuaVisitor) VisitDFloat(d *datas.DFloat) interface{} {
	return fmt.Sprintf("%v", d.Value)
}

func (v *ToLuaVisitor) VisitDDouble(d *datas.DDouble) interface{} {
	return fmt.Sprintf("%v", d.Value)
}

func (v *ToLuaVisitor) VisitDString(d *datas.DString) interface{} {
	// 简单转义双引号和换行
	str := strings.ReplaceAll(d.Value, "\\", "\\\\")
	str = strings.ReplaceAll(str, "\"", "\\\"")
	str = strings.ReplaceAll(str, "\n", "\\n")
	return fmt.Sprintf("\"%s\"", str)
}

func (v *ToLuaVisitor) VisitDText(d *datas.DText) interface{} {
	// 在 Lua 中，text 类型可以导出为带有 key 和 text 的 table，或者直接导出为字符串
	// 这里简单处理为只导出字符串或者包含 key 的 table
	str := strings.ReplaceAll(d.RawText, "\\", "\\\\")
	str = strings.ReplaceAll(str, "\"", "\\\"")
	str = strings.ReplaceAll(str, "\n", "\\n")

	if d.Key != "" {
		keyStr := strings.ReplaceAll(d.Key, "\\", "\\\\")
		keyStr = strings.ReplaceAll(keyStr, "\"", "\\\"")
		keyStr = strings.ReplaceAll(keyStr, "\n", "\\n")
		return fmt.Sprintf("{key=\"%s\", text=\"%s\"}", keyStr, str)
	}
	return fmt.Sprintf("\"%s\"", str)
}

func (v *ToLuaVisitor) VisitDEnum(d *datas.DEnum) interface{} {
	return fmt.Sprintf("%d", d.Value)
}

func (v *ToLuaVisitor) VisitDBean(d *datas.DBean) interface{} {
	var sb strings.Builder
	sb.WriteString("{")

	implBean := d.Type
	if d.ImplType != nil {
		implBean = d.ImplType
	}

	fieldCount := 0
	if implBean != nil && implBean.DefBean != nil {
		if beanImpl, ok := implBean.DefBean.(*defs.DefBeanImpl); ok {

			// 处理多态标识 ($type) -> 转化为 _type 或者不处理（可选，这里用 _type 标识）
			if d.Type.IsDynamic() && d.ImplType != nil {
				sb.WriteString(fmt.Sprintf("_type=\"%s\"", d.ImplType.DefBean.Name()))
				fieldCount++
			}

			for i, field := range beanImpl.HierarchyFields {
				if !field.NeedExport(v.TargetGroups) {
					continue
				}

				if i < len(d.Fields) {
					if d.Fields[i] != nil {
						if fieldCount > 0 {
							sb.WriteString(", ")
						}

						// Lua 的 key 如果是合法标识符可以直接写 key=value
						sb.WriteString(field.Raw.Name)
						sb.WriteString("=")
						sb.WriteString(fmt.Sprintf("%v", d.Fields[i].Accept(v)))
						fieldCount++
					}
				}
			}
		}
	}
	sb.WriteString("}")
	return sb.String()
}

func (v *ToLuaVisitor) formatList(elements []datas.DType) string {
	var sb strings.Builder
	sb.WriteString("{")
	for i, e := range elements {
		if i > 0 {
			sb.WriteString(", ")
		}
		if e != nil {
			sb.WriteString(fmt.Sprintf("%v", e.Accept(v)))
		} else {
			sb.WriteString("nil")
		}
	}
	sb.WriteString("}")
	return sb.String()
}

func (v *ToLuaVisitor) VisitDArray(d *datas.DArray) interface{} {
	return v.formatList(d.Elements)
}

func (v *ToLuaVisitor) VisitDList(d *datas.DList) interface{} {
	return v.formatList(d.Elements)
}

func (v *ToLuaVisitor) VisitDSet(d *datas.DSet) interface{} {
	return v.formatList(d.Elements)
}

func (v *ToLuaVisitor) VisitDMap(d *datas.DMap) interface{} {
	var sb strings.Builder
	sb.WriteString("{")
	i := 0
	for k, val := range d.Datas {
		if i > 0 {
			sb.WriteString(", ")
		}
		keyStr := fmt.Sprintf("%v", k.Accept(v))
		// 如果 key 是字符串，需要加 [] 括起来。因为上面 String 访问器已经包裹了 ""，这里直接包 [] 即可。
		// 如果 key 是数字，也应该包裹 []，这是 lua map 的标准格式
		sb.WriteString(fmt.Sprintf("[%s]=", keyStr))

		if val != nil {
			sb.WriteString(fmt.Sprintf("%v", val.Accept(v)))
		} else {
			sb.WriteString("nil")
		}
		i++
	}
	sb.WriteString("}")
	return sb.String()
}

func (v *ToLuaVisitor) VisitDDateTime(d *datas.DDateTime) interface{} {
	// Lua 中通常存时间戳或者字符串，这里直接转字符串
	return fmt.Sprintf("\"%s\"", d.String())
}
