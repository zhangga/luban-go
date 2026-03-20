package datatarget

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/zhangga/luban-go/core/codetarget"
	"github.com/zhangga/luban-go/core/defs"
)

type LuaDataTarget struct {
	visitor *ToLuaVisitor
}

func NewLuaDataTarget(targetGroups []string) *LuaDataTarget {
	return &LuaDataTarget{
		visitor: NewToLuaVisitor(targetGroups),
	}
}

func (t *LuaDataTarget) AggregationType() AggregationType {
	return AggregationTable
}

func (t *LuaDataTarget) ExportAllRecords() bool {
	return true
}

func (t *LuaDataTarget) ExportTable(table *defs.DefTable, records []*defs.Record) (*codetarget.OutputFile, error) {
	var buf bytes.Buffer
	
	buf.WriteString("local ")
	buf.WriteString(table.Name())
	buf.WriteString(" = {\n")

	// 判断表的模式 (List 还是 Map)，决定使用数组包裹还是字典包裹
	isMap := table.Raw.Mode == "map" || table.Raw.Mode == ""
	indexField := table.Raw.Index

	for _, r := range records {
		if r.Data != nil {
			luaStr := fmt.Sprintf("%v", r.Data.Accept(t.visitor))

			buf.WriteString("    ")
			
			if isMap {
				// 获取主键的值
				dBean := r.Data
				var keyStr string
				if beanImpl, ok := dBean.ImplType.DefBean.(*defs.DefBeanImpl); ok {
					for i, field := range beanImpl.HierarchyFields {
						if field.Raw.Name == indexField && i < len(dBean.Fields) {
							keyStr = fmt.Sprintf("%v", dBean.Fields[i].Accept(t.visitor))
							break
						}
					}
				}
				
				if keyStr == "" {
					return nil, fmt.Errorf("failed to find index field '%s' in table '%s'", indexField, table.Name())
				}
				
				buf.WriteString(fmt.Sprintf("[%s] = %s,\n", keyStr, luaStr))
			} else {
				// 列表模式，直接追加
				buf.WriteString(luaStr)
				buf.WriteString(",\n")
			}
		}
	}

	buf.WriteString("}\n\nreturn ")
	buf.WriteString(table.Name())
	buf.WriteString("\n")

	fileName := fmt.Sprintf("%s.lua", strings.ToLower(table.Name()))
	return codetarget.NewOutputFile(fileName, buf.Bytes()), nil
}

func (t *LuaDataTarget) ExportTables(tables []*defs.DefTable) (*codetarget.OutputFile, error) {
	return nil, fmt.Errorf("LuaDataTarget does not support ExportTables")
}

func (t *LuaDataTarget) ExportRecord(table *defs.DefTable, record *defs.Record) (*codetarget.OutputFile, error) {
	return nil, fmt.Errorf("LuaDataTarget does not support ExportRecord")
}
