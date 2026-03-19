package dataloader

import (
	"fmt"
	"io"
	"strings"

	"github.com/xuri/excelize/v2"
	"github.com/zhangga/luban-go/core/datas"
	"github.com/zhangga/luban-go/core/defs"
	"github.com/zhangga/luban-go/core/types"
)

type ExcelDataLoader struct {
	rawUrl      string
	sheets      []*RawSheet
	typeFactory *defs.TTypeFactory // 加入 Factory 引用以处理 TType 创建
}

func NewExcelDataLoader(typeFactory *defs.TTypeFactory) *ExcelDataLoader {
	return &ExcelDataLoader{
		sheets:      make([]*RawSheet, 0),
		typeFactory: typeFactory,
	}
}

func (l *ExcelDataLoader) RawUrl() string {
	return l.rawUrl
}

func (l *ExcelDataLoader) Load(rawUrl string, subAsset string, stream io.Reader) error {
	l.rawUrl = rawUrl

	f, err := excelize.OpenReader(stream)
	if err != nil {
		return fmt.Errorf("failed to open excel stream: %w", err)
	}
	defer f.Close()

	sheetList := f.GetSheetList()
	for _, sheetName := range sheetList {
		// 可根据 subAsset 过滤 sheet，如果 subAsset 不为空
		if subAsset != "" && sheetName != subAsset {
			continue
		}

		rows, err := f.GetRows(sheetName)
		if err != nil {
			return err
		}

		rawSheet := NewRawSheet(sheetName)
		for rIdx, row := range rows {
			cellRow := make([]*Cell, 0, len(row))
			for cIdx, colCell := range row {
				cellRow = append(cellRow, NewCell(rIdx, cIdx, colCell))
			}
			rawSheet.Cells = append(rawSheet.Cells, cellRow)
		}
		l.sheets = append(l.sheets, rawSheet)
	}

	return nil
}

func (l *ExcelDataLoader) ReadOne(t *types.TBean) *defs.Record {
	// 简单的提取逻辑示例，真正实现需要处理多态和字段偏移
	records := l.ReadMulti(t)
	if len(records) > 0 {
		return records[0]
	}
	return nil
}

func (l *ExcelDataLoader) ReadMulti(t *types.TBean) []*defs.Record {
	var records []*defs.Record

	for _, sheet := range l.sheets {
		// 假定第0行是表头（字段名），从第1行开始是数据
		if len(sheet.Cells) < 2 {
			continue
		}

		// 找到字段的列索引映射
		headers := sheet.Cells[0]
		fieldIndices := make(map[string]int)
		for _, cell := range headers {
			if strVal, ok := cell.Value.(string); ok {
				fieldIndices[strings.TrimSpace(strVal)] = cell.Column
			}
		}
		fmt.Printf("DEBUG headers mapping: %v\n", fieldIndices)

		// 迭代每一行
		for i := 1; i < len(sheet.Cells); i++ {
			row := sheet.Cells[i]

			// 过滤空行
			if len(row) == 0 {
				continue
			}

			// 解析 TBean 字段
			fieldsData := make([]datas.DType, 0)

			var implBeanType *types.TBean = t

			fmt.Printf("DEBUG resolving row %d for bean %s. IsDynamic: %v\n", i, t.DefBean.Name(), t.IsDynamic())

			// 如果是多态的话，通常会有一个字段或列标识它的子类型 (比如 $type)
			if t.IsDynamic() {
				// 尝试在列中查找 $type
				if typeColIdx, ok := fieldIndices["$type"]; ok && typeColIdx < len(row) {
					typeStr := fmt.Sprintf("%v", row[typeColIdx].Value)
					if beanImpl, ok := t.DefBean.(*defs.DefBeanImpl); ok {
						if childDef := beanImpl.TryGetChild(typeStr); childDef != nil {
							implBeanType = types.NewTBean(false, childDef, nil)
							fmt.Printf("DEBUG dynamically resolved poly type: %s for row %d\n", childDef.FullName(), i)
						} else {
							fmt.Printf("Warning: failed to find poly child type %s for bean %s\n", typeStr, t.DefBean.Name())
						}
					}
				} else {
					fmt.Printf("Warning: no $type column found in sheet for polymorphic bean %s\n", t.DefBean.Name())
				}
			}

			// 对于 TBean 的每一个字段进行遍历（按序填充）
			if beanImpl, ok := implBeanType.DefBean.(*defs.DefBeanImpl); ok {
				for _, f := range beanImpl.HierarchyFields {
					fieldName := f.Raw.Name
					colIdx, found := fieldIndices[fieldName]
					fmt.Printf("DEBUG field '%s' found=%v colIdx=%v\n", fieldName, found, colIdx)
					if found && colIdx < len(row) {
						cell := row[colIdx]
						strVal := fmt.Sprintf("%v", cell.Value)

						// 根据该列定义的原始类型字符串，解析出具体的 TType
						fieldType, _ := l.typeFactory.CreateType(f.Raw.Type)
						if fieldType != nil && strVal != "" {
							creator := NewDataCreator()
							dVal, err := creator.CreateField(fieldType, strVal)
							if err != nil {
								fmt.Printf("Warning: failed to create data for field %s, err: %v\n", fieldName, err)
								fieldsData = append(fieldsData, datas.NewDString(strVal)) // 兜底
							} else {
								fieldsData = append(fieldsData, dVal)
							}
						} else {
							// 兜底直接存 string，或留空
							if strVal != "" {
								fieldsData = append(fieldsData, datas.NewDString(strVal))
							} else {
								fieldsData = append(fieldsData, nil)
							}
						}
					} else {
						fieldsData = append(fieldsData, nil)
					}
				}
			} else {
				// Fallback：如果没有完整的元数据，把所有列加进去
				for _, cell := range row {
					strVal := fmt.Sprintf("%v", cell.Value)
					fieldsData = append(fieldsData, datas.NewDString(strVal))
				}
			}

			dBean := datas.NewDBean(t, implBeanType, fieldsData)
			records = append(records, defs.NewRecord(dBean, l.rawUrl, nil))
		}
	}

	return records
}
