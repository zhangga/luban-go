package dataloader

import (
	"strings"
	"testing"

	"github.com/zhangga/luban-go/core/defs"
	"github.com/zhangga/luban-go/core/rawdefs"
	"github.com/zhangga/luban-go/core/types"
)

func TestJsonDataLoader(t *testing.T) {
	jsonData := `
[
  {
    "Id": 101,
    "Name": "Sword",
    "List": [10, 20]
  },
  {
    "Id": 102,
    "Name": "Shield",
    "List": [30, 40]
  }
]
`

	rawAss := rawdefs.NewRawAssembly()
	rawAss.Beans = append(rawAss.Beans, &rawdefs.RawBean{
		Name: "TestJsonItem",
		Fields: []*rawdefs.RawField{
			{Name: "Id", Type: "int"},
			{Name: "Name", Type: "string"},
			{Name: "List", Type: "list,int"},
		},
	})
	assembly, _ := defs.NewDefAssemblyImpl(rawAss, "all", nil)
	typeFactory := defs.NewTTypeFactory(assembly)

	loader := NewJsonDataLoader(typeFactory)
	err := loader.Load("test.json", "", strings.NewReader(jsonData))
	if err != nil {
		t.Fatalf("Failed to load json: %v", err)
	}

	beanImpl := assembly.GetType("TestJsonItem").(*defs.DefBeanImpl)
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

func TestJsonDataLoader_SingleObject(t *testing.T) {
	jsonData := `
{
  "Id": 999,
  "Name": "Boss",
  "List": [100, 200]
}
`

	rawAss := rawdefs.NewRawAssembly()
	rawAss.Beans = append(rawAss.Beans, &rawdefs.RawBean{
		Name: "TestJsonItem",
		Fields: []*rawdefs.RawField{
			{Name: "Id", Type: "int"},
			{Name: "Name", Type: "string"},
			{Name: "List", Type: "list,int"},
		},
	})
	assembly, _ := defs.NewDefAssemblyImpl(rawAss, "all", nil)
	typeFactory := defs.NewTTypeFactory(assembly)

	loader := NewJsonDataLoader(typeFactory)
	err := loader.Load("test.json", "", strings.NewReader(jsonData))
	if err != nil {
		t.Fatalf("Failed to load json: %v", err)
	}

	beanImpl := assembly.GetType("TestJsonItem").(*defs.DefBeanImpl)
	dummyBeanType := types.NewTBean(false, beanImpl, nil)

	records := loader.ReadMulti(dummyBeanType)

	if len(records) != 1 {
		t.Fatalf("Expected 1 record, got %d", len(records))
	}

	r1 := records[0]
	if r1.Data.Fields[0].String() != "999" {
		t.Errorf("Expected 999, got %s", r1.Data.Fields[0].String())
	}
}