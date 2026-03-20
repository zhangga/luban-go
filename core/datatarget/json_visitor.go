package datatarget

import (
	"fmt"
	"github.com/zhangga/luban-go/core/datas"
	"github.com/zhangga/luban-go/core/defs"
)

// ToJsonVisitor 是一个实现 IDataVisitor 接口的访问者，将 DType 转换为可以被 json.Marshal 的对象
type ToJsonVisitor struct {
	TargetGroups []string
}

func NewToJsonVisitor(targetGroups []string) *ToJsonVisitor {
	return &ToJsonVisitor{
		TargetGroups: targetGroups,
	}
}

func (v *ToJsonVisitor) VisitDBool(d *datas.DBool) interface{} {
	return d.Value
}

func (v *ToJsonVisitor) VisitDByte(d *datas.DByte) interface{} {
	return d.Value
}

func (v *ToJsonVisitor) VisitDShort(d *datas.DShort) interface{} {
	return d.Value
}

func (v *ToJsonVisitor) VisitDInt(d *datas.DInt) interface{} {
	return d.Value
}

func (v *ToJsonVisitor) VisitDLong(d *datas.DLong) interface{} {
	return d.Value
}

func (v *ToJsonVisitor) VisitDFloat(d *datas.DFloat) interface{} {
	return d.Value
}

func (v *ToJsonVisitor) VisitDDouble(d *datas.DDouble) interface{} {
	return d.Value
}

func (v *ToJsonVisitor) VisitDString(d *datas.DString) interface{} {
	return d.Value
}

func (v *ToJsonVisitor) VisitDText(d *datas.DText) interface{} {
	// 在 JSON 中导出，通常 text 可以导出为一个包含 key 的对象，或者直接导成文本 (取决于设置)
	// 原版 Luban 中可以开启 l10n.textValueFieldName。这里我们简化处理，将其输出为对象或字符串
	// 如果 Key 不为空，可以用 map 包含 key 和 text，否则只输出 text
	if d.Key != "" {
		m := make(map[string]interface{})
		m["key"] = d.Key
		m["text"] = d.RawText
		return m
	}
	return d.RawText
}

func (v *ToJsonVisitor) VisitDEnum(d *datas.DEnum) interface{} {
	if d.StrValue != "" {
		return d.StrValue
	}
	return d.Value
}

func (v *ToJsonVisitor) VisitDBean(d *datas.DBean) interface{} {
	m := make(map[string]interface{})

	// 处理多态标识 ($type)
	if d.Type.IsDynamic() && d.ImplType != nil {
		m["$type"] = d.ImplType.DefBean.Name()
	}

	implBean := d.Type
	if d.ImplType != nil {
		implBean = d.ImplType
	}

	if implBean != nil && implBean.DefBean != nil {
		if beanImpl, ok := implBean.DefBean.(*defs.DefBeanImpl); ok {
			for i, field := range beanImpl.HierarchyFields {
				// 添加 group 过滤
				if !field.NeedExport(v.TargetGroups) {
					continue
				}

				if i < len(d.Fields) {
					if d.Fields[i] != nil {
						m[field.Raw.Name] = d.Fields[i].Accept(v)
					}
				}
			}
		}
	}
	return m
}

func (v *ToJsonVisitor) VisitDArray(d *datas.DArray) interface{} {
	arr := make([]interface{}, 0, len(d.Elements))
	for _, e := range d.Elements {
		if e != nil {
			arr = append(arr, e.Accept(v))
		} else {
			arr = append(arr, nil)
		}
	}
	return arr
}

func (v *ToJsonVisitor) VisitDList(d *datas.DList) interface{} {
	arr := make([]interface{}, 0, len(d.Elements))
	for _, e := range d.Elements {
		if e != nil {
			arr = append(arr, e.Accept(v))
		} else {
			arr = append(arr, nil)
		}
	}
	return arr
}

func (v *ToJsonVisitor) VisitDSet(d *datas.DSet) interface{} {
	arr := make([]interface{}, 0, len(d.Elements))
	for _, e := range d.Elements {
		if e != nil {
			arr = append(arr, e.Accept(v))
		} else {
			arr = append(arr, nil)
		}
	}
	return arr
}

func (v *ToJsonVisitor) VisitDMap(d *datas.DMap) interface{} {
	m := make(map[string]interface{})
	for k, val := range d.Datas {
		// Map的Key在JSON中通常必须是字符串
		keyStr := fmt.Sprintf("%v", k.Accept(v))
		if val != nil {
			m[keyStr] = val.Accept(v)
		} else {
			m[keyStr] = nil
		}
	}
	return m
}

func (v *ToJsonVisitor) VisitDDateTime(d *datas.DDateTime) interface{} {
	return d.String()
}
