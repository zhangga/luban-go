package datatarget

import (
	"strings"
	"testing"

	"github.com/zhangga/luban-go/core/datas"
	"github.com/zhangga/luban-go/core/defs"
	"github.com/zhangga/luban-go/core/rawdefs"
	"github.com/zhangga/luban-go/core/types"
)

func TestLuaDataTarget_ExportTable(t *testing.T) {
	// 准备元数据
	rawBean := &rawdefs.RawBean{
		Name: "Item",
		Fields: []*rawdefs.RawField{
			{Name: "Id", Type: "int"},
			{Name: "Name", Type: "string"},
		},
	}
	defBean := defs.NewDefBeanImpl(rawBean)
	defBean.Fields = append(defBean.Fields, defs.NewDefField(rawBean.Fields[0]))
	defBean.Fields = append(defBean.Fields, defs.NewDefField(rawBean.Fields[1]))
	defBean.HierarchyFields = append(defBean.HierarchyFields, defs.NewDefField(rawBean.Fields[0]))
	defBean.HierarchyFields = append(defBean.HierarchyFields, defs.NewDefField(rawBean.Fields[1]))

	tBean := types.NewTBean(false, defBean, nil)

	// 准备数据
	dBean1 := datas.NewDBean(tBean, tBean, []datas.DType{
		datas.NewDInt(101),
		datas.NewDString("Sword"),
	})
	record1 := defs.NewRecord(dBean1, "test.xlsx", nil)

	dBean2 := datas.NewDBean(tBean, tBean, []datas.DType{
		datas.NewDInt(102),
		datas.NewDString("Shield"),
	})
	record2 := defs.NewRecord(dBean2, "test.xlsx", nil)

	rawTable := &rawdefs.RawTable{Name: "TbItem", Index: "Id", Mode: "map"}
	defTable := defs.NewDefTable(rawTable)

	// 导出
	target := NewLuaDataTarget([]string{})
	out, err := target.ExportTable(defTable, []*defs.Record{record1, record2})
	if err != nil {
		t.Fatalf("Failed to export table: %v", err)
	}

	if out.File != "tbitem.lua" {
		t.Errorf("Expected filename tbitem.lua, got %s", out.File)
	}

	outStr := string(out.Content)
	if !strings.Contains(outStr, `local TbItem = {`) {
		t.Errorf("Expected to contain 'local TbItem = {', got %s", outStr)
	}
	if !strings.Contains(outStr, `[101] = {Id=101, Name="Sword"}`) {
		t.Errorf("Expected to contain '[101] = {Id=101, Name=\"Sword\"}', got %s", outStr)
	}
	if !strings.Contains(outStr, `[102] = {Id=102, Name="Shield"}`) {
		t.Errorf("Expected to contain '[102] = {Id=102, Name=\"Shield\"}', got %s", outStr)
	}
	if !strings.Contains(outStr, `return TbItem`) {
		t.Errorf("Expected to contain 'return TbItem', got %s", outStr)
	}
}

func TestLuaVisitor_List(t *testing.T) {
	visitor := NewToLuaVisitor(nil)
	
	dList := datas.NewDList(nil, []datas.DType{
		datas.NewDInt(1),
		datas.NewDInt(2),
		datas.NewDInt(3),
	})

	luaStr := dList.Accept(visitor).(string)
	if luaStr != "{1, 2, 3}" {
		t.Errorf("Expected {1, 2, 3}, got %s", luaStr)
	}
}

func TestLuaVisitor_Map(t *testing.T) {
	visitor := NewToLuaVisitor(nil)
	
	dMap := datas.NewDMap(nil, map[datas.DType]datas.DType{
		datas.NewDString("hp"): datas.NewDInt(100),
	})

	luaStr := dMap.Accept(visitor).(string)
	if luaStr != "{[\"hp\"]=100}" {
		t.Errorf("Expected {[\"hp\"]=100}, got %s", luaStr)
	}
}
