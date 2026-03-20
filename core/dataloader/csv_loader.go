package dataloader

import (
	"encoding/csv"
	"fmt"
	"io"
	"strings"

	"github.com/zhangga/luban-go/core/datas"
	"github.com/zhangga/luban-go/core/defs"
	"github.com/zhangga/luban-go/core/types"
)

type CsvDataLoader struct {
	rawUrl      string
	records     [][]string
	typeFactory *defs.TTypeFactory
}

func NewCsvDataLoader(typeFactory *defs.TTypeFactory) *CsvDataLoader {
	return &CsvDataLoader{
		typeFactory: typeFactory,
	}
}

func (l *CsvDataLoader) RawUrl() string {
	return l.rawUrl
}

func (l *CsvDataLoader) Load(rawUrl string, subAsset string, stream io.Reader) error {
	l.rawUrl = rawUrl

	reader := csv.NewReader(stream)
	// 允许可变列数
	reader.FieldsPerRecord = -1
	
	records, err := reader.ReadAll()
	if err != nil {
		return fmt.Errorf("failed to read csv: %w", err)
	}

	l.records = records
	return nil
}

func (l *CsvDataLoader) ReadOne(t *types.TBean) *defs.Record {
	records := l.ReadMulti(t)
	if len(records) > 0 {
		return records[0]
	}
	return nil
}

func (l *CsvDataLoader) ReadMulti(t *types.TBean) []*defs.Record {
	var result []*defs.Record

	if len(l.records) < 2 {
		return result
	}

	fieldIndices := make(map[string][]int)
	headerRowIndex := 0

	// 1. 寻找表头
	for rIdx, row := range l.records {
		if len(row) > 0 {
			firstCellStr := strings.TrimSpace(row[0])
			if strings.HasPrefix(firstCellStr, "##var") || (!strings.HasPrefix(firstCellStr, "##") && !strings.HasPrefix(firstCellStr, "#")) {
				headerRowIndex = rIdx
				break
			}
		}
	}

	headers := l.records[headerRowIndex]
	for cIdx, header := range headers {
		strVal := strings.TrimSpace(header)
		if strings.HasPrefix(strVal, "#") {
			continue
		}
		if strVal != "" {
			fieldIndices[strVal] = append(fieldIndices[strVal], cIdx)
		}
	}

	// 2. 确定主键列
	var primaryColIdx = -1
	if len(t.DefBean.(*defs.DefBeanImpl).HierarchyFields) > 0 {
		firstFieldName := t.DefBean.(*defs.DefBeanImpl).HierarchyFields[0].Raw.Name
		if idxs, ok := fieldIndices[firstFieldName]; ok && len(idxs) > 0 {
			primaryColIdx = idxs[0]
		}
	}

	// 3. 解析记录，支持多行
	for i := headerRowIndex + 1; i < len(l.records); {
		row := l.records[i]

		if len(row) == 0 {
			i++
			continue
		}

		firstCellStr := strings.TrimSpace(row[0])
		if strings.HasPrefix(firstCellStr, "##") || strings.HasPrefix(firstCellStr, "#") {
			i++
			continue
		}

		fieldsData := make([]datas.DType, 0)
		var implBeanType *types.TBean = t

		if t.IsDynamic() {
			if typeColIdxs, ok := fieldIndices["$type"]; ok && len(typeColIdxs) > 0 {
				typeColIdx := typeColIdxs[0]
				if typeColIdx < len(row) {
					typeStr := strings.TrimSpace(row[typeColIdx])
					if beanImpl, ok := t.DefBean.(*defs.DefBeanImpl); ok {
						if childDef := beanImpl.TryGetChild(typeStr); childDef != nil {
							implBeanType = types.NewTBean(false, childDef, nil)
						}
					}
				}
			}
		}

		if beanImpl, ok := implBeanType.DefBean.(*defs.DefBeanImpl); ok {
			for _, f := range beanImpl.HierarchyFields {
				fieldName := f.Raw.Name
				colIdxs, found := fieldIndices[fieldName]

				fieldType, _ := l.typeFactory.CreateType(f.Raw.Type)
				if fieldType == nil {
					fieldsData = append(fieldsData, nil)
					continue
				}

				var strVals []string
				if found {
					for r := i; r < len(l.records); r++ {
						currRow := l.records[r]
						
						if len(currRow) > 0 {
							fStr := strings.TrimSpace(currRow[0])
							if strings.HasPrefix(fStr, "##") || strings.HasPrefix(fStr, "#") {
								continue
							}
						}

						if r > i && primaryColIdx >= 0 && primaryColIdx < len(currRow) {
							pkVal := strings.TrimSpace(currRow[primaryColIdx])
							if pkVal != "" {
								break
							}
						}

						for _, colIdx := range colIdxs {
							if colIdx < len(currRow) {
								strVal := strings.TrimSpace(currRow[colIdx])
								if strVal != "" {
									strVals = append(strVals, strVal)
								}
							}
						}
					}
				}

				if len(strVals) > 0 {
					mergedStr := strings.Join(strVals, ",")
					creator := NewDataCreator()
					dVal, err := creator.CreateField(fieldType, mergedStr)
					if err != nil {
						fieldsData = append(fieldsData, datas.NewDString(mergedStr))
					} else {
						fieldsData = append(fieldsData, dVal)
					}
				} else {
					fieldsData = append(fieldsData, nil)
				}
			}
		} else {
			for _, cell := range row {
				fieldsData = append(fieldsData, datas.NewDString(strings.TrimSpace(cell)))
			}
		}

		dBean := datas.NewDBean(t, implBeanType, fieldsData)
		result = append(result, defs.NewRecord(dBean, l.rawUrl, nil))

		// 计算 nextI
		nextI := i + 1
		for r := i + 1; r < len(l.records); r++ {
			currRow := l.records[r]
			
			if len(currRow) > 0 {
				fStr := strings.TrimSpace(currRow[0])
				if strings.HasPrefix(fStr, "##") || strings.HasPrefix(fStr, "#") {
					nextI++
					continue
				}
			}

			if primaryColIdx >= 0 {
				if primaryColIdx < len(currRow) {
					pkVal := strings.TrimSpace(currRow[primaryColIdx])
					if pkVal != "" {
						break
					} else {
						nextI++
					}
				} else {
					nextI++
				}
			} else {
				break
			}
		}
		
		if nextI <= i {
			nextI = i + 1
		}
		i = nextI
	}

	return result
}