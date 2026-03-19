package datatarget

import (
	"encoding/json"
	"fmt"
	"github.com/zhangga/luban-go/core/codetarget"
	"github.com/zhangga/luban-go/core/defs"
	"strings"
)

type JsonDataTarget struct {
	visitor *ToJsonVisitor
}

func NewJsonDataTarget() *JsonDataTarget {
	return &JsonDataTarget{
		visitor: NewToJsonVisitor(),
	}
}

func (t *JsonDataTarget) AggregationType() AggregationType {
	return AggregationTable
}

func (t *JsonDataTarget) ExportAllRecords() bool {
	return true
}

func (t *JsonDataTarget) ExportTable(table *defs.DefTable, records []*defs.Record) (*codetarget.OutputFile, error) {
	// 将所有 records 转换成可序列化的对象列表
	var jsonObjList []interface{}
	for _, r := range records {
		if r.Data != nil {
			jsonObj := r.Data.Accept(t.visitor)
			jsonObjList = append(jsonObjList, jsonObj)
		}
	}

	bytes, err := json.MarshalIndent(jsonObjList, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("failed to marshal table %s to json: %v", table.Name(), err)
	}

	fileName := fmt.Sprintf("%s.json", strings.ToLower(table.Name()))
	return codetarget.NewOutputFile(fileName, bytes), nil
}

func (t *JsonDataTarget) ExportTables(tables []*defs.DefTable) (*codetarget.OutputFile, error) {
	return nil, fmt.Errorf("JsonDataTarget does not support ExportTables")
}

func (t *JsonDataTarget) ExportRecord(table *defs.DefTable, record *defs.Record) (*codetarget.OutputFile, error) {
	return nil, fmt.Errorf("JsonDataTarget does not support ExportRecord")
}
