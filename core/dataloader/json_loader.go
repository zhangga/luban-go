package dataloader

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"strings"

	"github.com/zhangga/luban-go/core/datas"
	"github.com/zhangga/luban-go/core/defs"
	"github.com/zhangga/luban-go/core/types"
)

type JsonDataLoader struct {
	rawUrl      string
	recordsData []map[string]interface{}
	typeFactory *defs.TTypeFactory
}

func NewJsonDataLoader(typeFactory *defs.TTypeFactory) *JsonDataLoader {
	return &JsonDataLoader{
		typeFactory: typeFactory,
	}
}

func (l *JsonDataLoader) RawUrl() string {
	return l.rawUrl
}

func (l *JsonDataLoader) Load(rawUrl string, subAsset string, stream io.Reader) error {
	l.rawUrl = rawUrl

	bytes, err := ioutil.ReadAll(stream)
	if err != nil {
		return fmt.Errorf("failed to read json: %w", err)
	}

	// 尝试解析为数组（多条记录）
	var arrayData []map[string]interface{}
	if err := json.Unmarshal(bytes, &arrayData); err == nil {
		l.recordsData = arrayData
		return nil
	}

	// 尝试解析为单条记录
	var singleData map[string]interface{}
	if err := json.Unmarshal(bytes, &singleData); err == nil {
		l.recordsData = []map[string]interface{}{singleData}
		return nil
	}

	return fmt.Errorf("failed to parse json, neither array nor object: %v", err)
}

func (l *JsonDataLoader) ReadOne(t *types.TBean) *defs.Record {
	records := l.ReadMulti(t)
	if len(records) > 0 {
		return records[0]
	}
	return nil
}

func (l *JsonDataLoader) ReadMulti(t *types.TBean) []*defs.Record {
	var result []*defs.Record

	for _, rawData := range l.recordsData {
		dBean := l.readBean(t, rawData)
		result = append(result, defs.NewRecord(dBean, l.rawUrl, nil))
	}

	return result
}

func (l *JsonDataLoader) readBean(beanType *types.TBean, data map[string]interface{}) *datas.DBean {
	fieldsData := make([]datas.DType, 0)
	
	var implBeanType = beanType

	// 多态支持
	if beanType.IsDynamic() {
		if typeVal, ok := data["$type"]; ok {
			if typeStr, ok := typeVal.(string); ok {
				if beanImpl, ok := beanType.DefBean.(*defs.DefBeanImpl); ok {
					if childDef := beanImpl.TryGetChild(typeStr); childDef != nil {
						implBeanType = types.NewTBean(false, childDef, nil)
					}
				}
			}
		}
	}

	beanImpl, ok := implBeanType.DefBean.(*defs.DefBeanImpl)
	if !ok {
		return datas.NewDBean(beanType, implBeanType, fieldsData)
	}

	creator := NewDataCreator()

	for _, f := range beanImpl.HierarchyFields {
		fieldName := f.Raw.Name
		fieldType, _ := l.typeFactory.CreateType(f.Raw.Type)
		
		if fieldType == nil {
			fieldsData = append(fieldsData, nil)
			continue
		}

		if val, exists := data[fieldName]; exists {
			// 如果是嵌套的 Bean 且传入的是 map
			if nestedBean, isBean := fieldType.(*types.TBean); isBean {
				if nestedMap, ok := val.(map[string]interface{}); ok {
					nestedData := l.readBean(nestedBean, nestedMap)
					fieldsData = append(fieldsData, nestedData)
					continue
				}
			}
			
			// 如果是基本类型或列表，我们将值转为字符串交给 DataCreator 创建
			// TODO: 对 Json 原生的 List/Map 做更优雅的原生解析，这里先降级为字符串转换统一处理
			strVal := fmt.Sprintf("%v", val)
			// 特殊处理 JSON 数组的字符串化格式，便于 Creator 解析
			if arr, isArr := val.([]interface{}); isArr {
				var strVals []string
				for _, v := range arr {
					strVals = append(strVals, fmt.Sprintf("%v", v))
				}
				strVal = strings.Join(strVals, ",")
			}

			dVal, err := creator.CreateField(fieldType, strVal)
			if err != nil {
				fmt.Printf("Warning: failed to create data for field %s from json: %v\n", fieldName, err)
				fieldsData = append(fieldsData, datas.NewDString(strVal))
			} else {
				fieldsData = append(fieldsData, dVal)
			}
		} else {
			fieldsData = append(fieldsData, nil)
		}
	}

	return datas.NewDBean(beanType, implBeanType, fieldsData)
}