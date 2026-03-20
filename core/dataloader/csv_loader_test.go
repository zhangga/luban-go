package dataloader

import (
	"strings"
	"testing"

	"github.com/zhangga/luban-go/core/defs"
	"github.com/zhangga/luban-go/core/rawdefs"
	"github.com/zhangga/luban-go/core/types"
)

func createTestCsv() string {
	// 跟 Excel 测试保持一致，第一列是控制列，后面是字段
	return `##var,Id,Name,List,List
##type,int,string,list,int
,101,Sword,10,20
,102,Shield,30,40
`
}

func TestCsvDataLoader(t *testing.T) {
	csvData := createTestCsv()

	rawAss := rawdefs.NewRawAssembly()
	rawAss.Beans = append(rawAss.Beans, &rawdefs.RawBean{
		Name: "TestCsvItem",
		Fields: []*rawdefs.RawField{
			{Name: "Id", Type: "int"},
			{Name: "Name", Type: "string"},
			{Name: "List", Type: "list,int"},
		},
	})
	assembly, _ := defs.NewDefAssemblyImpl(rawAss, "all", nil)
	typeFactory := defs.NewTTypeFactory(assembly)

	loader := NewCsvDataLoader(typeFactory)
	err := loader.Load("test.csv", "", strings.NewReader(csvData))
	if err != nil {
		t.Fatalf("Failed to load csv: %v", err)
	}

	beanImpl := assembly.GetType("TestCsvItem").(*defs.DefBeanImpl)
	for _, f := range rawAss.Beans[0].Fields {
		beanImpl.Fields = append(beanImpl.Fields, defs.NewDefField(f))
		beanImpl.HierarchyFields = append(beanImpl.HierarchyFields, defs.NewDefField(f))
	}
	dummyBeanType := types.NewTBean(false, beanImpl, nil)

	records := loader.ReadMulti(dummyBeanType)

	if len(records) != 2 {
		t.Fatalf("Expected 2 records, got %d", len(records))
	}

	r1 := records[0]
	if r1.Data.Fields[0].String() != "101" {
		t.Errorf("Expected 101, got %s", r1.Data.Fields[0].String())
	}
	if r1.Data.Fields[1].String() != "Sword" {
		t.Errorf("Expected Sword, got %s", r1.Data.Fields[1].String())
	}
	if r1.Data.Fields[2].String() != "[10, 20]" {
		t.Errorf("Expected [10, 20], got %s", r1.Data.Fields[2].String())
	}
}