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

		// 找到字段的列索引映射，并处理控制列 ## 和忽略列 #
		// 为了支持同名多列的情况，我们将映射从 map[string]int 改为 map[string][]int
		fieldIndices := make(map[string][]int)
		
		// 寻找包含 ##var（或没有##的普通变量行）的那一行作为表头
		headerRowIndex := 0
		for rIdx, row := range sheet.Cells {
			if len(row) > 0 {
				firstCellStr := fmt.Sprintf("%v", row[0].Value)
				// 原版 luban 中，第一列经常被用来做控制标记
				// 以 ## 起始的行是控制行，##var 表示这是变量名行
				if strings.HasPrefix(firstCellStr, "##var") || (!strings.HasPrefix(firstCellStr, "##") && !strings.HasPrefix(firstCellStr, "#")) {
					headerRowIndex = rIdx
					break
				}
			}
		}

		headers := sheet.Cells[headerRowIndex]
		for _, cell := range headers {
			if strVal, ok := cell.Value.(string); ok {
				strVal = strings.TrimSpace(strVal)
				// 忽略以 # 开头的列
				if strings.HasPrefix(strVal, "#") {
					continue
				}
				if strVal != "" {
					fieldIndices[strVal] = append(fieldIndices[strVal], cell.Column)
				}
			}
		}
		fmt.Printf("DEBUG headers mapping: %v\n", fieldIndices)

		// 找出主键列（通常是第一列，或者通过一些方式指定，这里简单取第一列不是 ## 且非空的第一个有效字段的列号，或者干脆默认 Id 的列，或者默认非忽略的第一列）
		// 在这里，如果一个数据行除了控制列外，前几列（或者是主键列）为空，则认为是多行延续。
		// 更好的策略：如果发现当前行的“主键”为空，则合并到上一行。
		// 我们假设 $type 和 第一个字段 为关键列。如果有，且均为空，则为多行。
		
		var primaryColIdx = -1
		if len(t.DefBean.(*defs.DefBeanImpl).HierarchyFields) > 0 {
			firstFieldName := t.DefBean.(*defs.DefBeanImpl).HierarchyFields[0].Raw.Name
			if idxs, ok := fieldIndices[firstFieldName]; ok && len(idxs) > 0 {
				primaryColIdx = idxs[0]
			}
		}

		// 迭代每一行数据 (从表头行的下一行开始)
		for i := headerRowIndex + 1; i < len(sheet.Cells); {
			row := sheet.Cells[i]

			// 过滤空行
			if len(row) == 0 {
				i++
				continue
			}

			// 处理行控制标记
			firstCellStr := strings.TrimSpace(fmt.Sprintf("%v", row[0].Value))
			// 忽略注释行（以 ## 起始）或数据禁用行（通常是 # 起始，或者视业务而定）
			if strings.HasPrefix(firstCellStr, "##") || strings.HasPrefix(firstCellStr, "#") {
				i++
				continue
			}

			// 解析 TBean 字段
			fieldsData := make([]datas.DType, 0)

			var implBeanType *types.TBean = t

			fmt.Printf("DEBUG resolving row %d for bean %s. IsDynamic: %v\n", i, t.DefBean.Name(), t.IsDynamic())

			// 如果是多态的话，通常会有一个字段或列标识它的子类型 (比如 $type)
			if t.IsDynamic() {
				// 尝试在列中查找 $type
				if typeColIdxs, ok := fieldIndices["$type"]; ok && len(typeColIdxs) > 0 {
					typeColIdx := typeColIdxs[0]
					if typeColIdx < len(row) {
						typeStr := fmt.Sprintf("%v", row[typeColIdx].Value)
						if beanImpl, ok := t.DefBean.(*defs.DefBeanImpl); ok {
							if childDef := beanImpl.TryGetChild(typeStr); childDef != nil {
								implBeanType = types.NewTBean(false, childDef, nil)
								fmt.Printf("DEBUG dynamically resolved poly type: %s for row %d\n", childDef.FullName(), i)
							} else {
								fmt.Printf("Warning: failed to find poly child type %s for bean %s\n", typeStr, t.DefBean.Name())
							}
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
					colIdxs, found := fieldIndices[fieldName]
					fmt.Printf("DEBUG field '%s' found=%v colIdxs=%v\n", fieldName, found, colIdxs)

					// 解析多列合并（比如数组/列表）或者单列
					fieldType, _ := l.typeFactory.CreateType(f.Raw.Type)
					if fieldType == nil {
						fieldsData = append(fieldsData, nil)
						continue
					}

					var strVals []string
					if found {
						// 收集当前行以及后续属于同一条记录的行
						for r := i; r < len(sheet.Cells); r++ {
							currRow := sheet.Cells[r]
							
							// 跳过控制行
							if len(currRow) > 0 {
								firstCellStr := strings.TrimSpace(fmt.Sprintf("%v", currRow[0].Value))
								if strings.HasPrefix(firstCellStr, "##") || strings.HasPrefix(firstCellStr, "#") {
									continue
								}
							}

							// 如果不是第一行，检查主键是否为空，如果不为空，说明是新记录，不能再继续收集
							if r > i && primaryColIdx >= 0 && primaryColIdx < len(currRow) {
								pkVal := strings.TrimSpace(fmt.Sprintf("%v", currRow[primaryColIdx].Value))
								if pkVal != "" {
									break // 遇到新记录，停止收集
								}
							}

							for _, colIdx := range colIdxs {
								if colIdx < len(currRow) {
									cell := currRow[colIdx]
									strVal := strings.TrimSpace(fmt.Sprintf("%v", cell.Value))
									if strVal != "" {
										strVals = append(strVals, strVal)
									}
								}
							}
						}
					}

					if len(strVals) > 0 {
						// 拼装多列的值
						// 简单处理：如果是集合类型且有多列，将它们用逗号连接起来交给 Creator 解析
						mergedStr := strings.Join(strVals, ",")
						
						creator := NewDataCreator()
						dVal, err := creator.CreateField(fieldType, mergedStr)
						if err != nil {
							fmt.Printf("Warning: failed to create data for field %s, err: %v\n", fieldName, err)
							fieldsData = append(fieldsData, datas.NewDString(mergedStr)) // 兜底
						} else {
							fieldsData = append(fieldsData, dVal)
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

			// 计算跳过多少行（因为内部已经合并了后面的空主键行）
			nextI := i + 1
			for r := i + 1; r < len(sheet.Cells); r++ {
				currRow := sheet.Cells[r]
				
				// 跳过控制行和忽略行，同时也要算在 nextI 里面，因为要跳过它们
				if len(currRow) > 0 {
					firstCellStr := strings.TrimSpace(fmt.Sprintf("%v", currRow[0].Value))
					if strings.HasPrefix(firstCellStr, "##") || strings.HasPrefix(firstCellStr, "#") {
						nextI++
						continue
					}
				}

				// 如果能够检查到主键列
				if primaryColIdx >= 0 {
					if primaryColIdx < len(currRow) {
						pkVal := strings.TrimSpace(fmt.Sprintf("%v", currRow[primaryColIdx].Value))
						if pkVal != "" {
							// 遇到新记录了，跳出计算
							break
						} else {
							// 是同一条记录的延续
							nextI++
						}
					} else {
						// 主键列不存在，等同于主键为空（延续行）
						nextI++
					}
				} else {
					// 无法判断主键（可能根本没有），那就默认当成一行记录
					break
				}
			}
			
			// 防御性编程：如果 nextI 没有增加，强制增加避免死循环
			if nextI <= i {
				nextI = i + 1
			}
			i = nextI
		}
	}

	return records
}
