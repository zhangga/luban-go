package datatarget

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/zhangga/luban-go/core/datas"
	"github.com/zhangga/luban-go/core/defs"
	"github.com/zhangga/luban-go/core/rawdefs"
	"github.com/zhangga/luban-go/core/types"
)

func TestJsonDataTarget_ExportTable(t *testing.T) {
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

	rawTable := &rawdefs.RawTable{Name: "TbItem"}
	defTable := defs.NewDefTable(rawTable)

	// 导出
	target := NewJsonDataTarget([]string{})
	out, err := target.ExportTable(defTable, []*defs.Record{record1, record2})
	if err != nil {
		t.Fatalf("Failed to export table: %v", err)
	}

	if out.File != "tbitem.json" {
		t.Errorf("Expected filename tbitem.json, got %s", out.File)
	}

	outStr := string(out.Content)
	if !strings.Contains(outStr, `"Id": 101`) {
		t.Errorf("Expected to contain 'Id: 101', got %s", outStr)
	}
	if !strings.Contains(outStr, `"Name": "Shield"`) {
		t.Errorf("Expected to contain 'Name: Shield', got %s", outStr)
	}

	// 验证 JSON 合法性
	var resultList []map[string]interface{}
	if err := json.Unmarshal(out.Content, &resultList); err != nil {
		t.Fatalf("Output is not valid JSON: %v", err)
	}
	if len(resultList) != 2 {
		t.Fatalf("Expected 2 JSON objects, got %d", len(resultList))
	}
}
