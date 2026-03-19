package codetarget

import (
	"strings"
	"testing"

	"github.com/zhangga/luban-go/core/defs"
	"github.com/zhangga/luban-go/core/rawdefs"
)

func TestGoCodeTarget_GenerateBean(t *testing.T) {
	rawBean := &rawdefs.RawBean{
		Namespace: "Demo",
		Name:      "Item",
		Comment:   "这是一个物品配置",
		Fields: []*rawdefs.RawField{
			{Name: "id", Type: "int", Comment: "物品ID"},
			{Name: "name", Type: "string", Comment: "物品名称"},
		},
	}

	defBean := defs.NewDefBeanImpl(rawBean)
	// 将 RawField 转换为 DefField
	for _, f := range rawBean.Fields {
		defBean.Fields = append(defBean.Fields, defs.NewDefField(f))
	}

	target := NewGoCodeTarget()
	outBytes, err := target.GenerateBean(defBean)
	if err != nil {
		t.Fatalf("Failed to generate bean: %v", err)
	}

	outStr := string(outBytes)

	if !strings.Contains(outStr, "package demo") {
		t.Errorf("Expected package demo, got: \n%s", outStr)
	}
	if !strings.Contains(outStr, "type Item struct") {
		t.Errorf("Expected type Item struct, got: \n%s", outStr)
	}
	if !strings.Contains(outStr, "Id int32 // 物品ID") {
		t.Errorf("Expected Id int32, got: \n%s", outStr)
	}
	if !strings.Contains(outStr, "Name string // 物品名称") {
		t.Errorf("Expected Name string, got: \n%s", outStr)
	}
}

func TestGoCodeTarget_GenerateTable(t *testing.T) {
	rawTable := &rawdefs.RawTable{
		Namespace: "Demo",
		Name:      "TbItem",
		ValueType: "Item",
		Index:     "id",
	}

	defTable := defs.NewDefTable(rawTable)
	target := NewGoCodeTarget()

	outBytes, err := target.GenerateTable(defTable)
	if err != nil {
		t.Fatalf("Failed to generate table: %v", err)
	}

	outStr := string(outBytes)

	if !strings.Contains(outStr, "package demo") {
		t.Errorf("Expected package demo, got: \n%s", outStr)
	}
	if !strings.Contains(outStr, "type TbItem struct") {
		t.Errorf("Expected type TbItem struct, got: \n%s", outStr)
	}
	if !strings.Contains(outStr, "DataList []*Item") {
		t.Errorf("Expected DataList []*Item, got: \n%s", outStr)
	}
	if !strings.Contains(outStr, "DataMap  map[int32]*Item") {
		t.Errorf("Expected DataMap  map[int32]*Item, got: \n%s", outStr)
	}
}
