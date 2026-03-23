package defs

import (
	"testing"

	"github.com/zhangga/luban-go/core/rawdefs"
)

func TestDefAssemblyImpl_ValidateBean(t *testing.T) {
	rawAss := rawdefs.NewRawAssembly()
	rawAss.Beans = append(rawAss.Beans, &rawdefs.RawBean{
		Name: "TestBean",
		Fields: []*rawdefs.RawField{
			{Name: "Id", Type: "int"},
			{Name: "Id", Type: "string"}, // 重复字段
		},
	})

	_, err := NewDefAssemblyImpl(rawAss, "all", nil)
	if err == nil {
		t.Fatal("Expected error for duplicated bean fields, but got nil")
	}
	if err.Error() != "bean 'TestBean' has duplicated field name: 'Id'" {
		t.Fatalf("Unexpected error message: %v", err)
	}
}

func TestDefAssemblyImpl_ValidateEnum(t *testing.T) {
	rawAss := rawdefs.NewRawAssembly()
	rawAss.Enums = append(rawAss.Enums, &rawdefs.RawEnum{
		Name: "TestEnum",
		Items: []*rawdefs.EnumItem{
			{Name: "A", Value: "1"},
			{Name: "A", Value: "2"}, // 重复名称
		},
	})

	_, err := NewDefAssemblyImpl(rawAss, "all", nil)
	if err == nil {
		t.Fatal("Expected error for duplicated enum items, but got nil")
	}
	if err.Error() != "enum 'TestEnum' has duplicated item name: 'A'" {
		t.Fatalf("Unexpected error message: %v", err)
	}

	rawAss2 := rawdefs.NewRawAssembly()
	rawAss2.Enums = append(rawAss2.Enums, &rawdefs.RawEnum{
		Name: "TestEnum2",
		Items: []*rawdefs.EnumItem{
			{Name: "A", Value: "1"},
			{Name: "B", Value: "1"}, // 重复值
		},
	})

	_, err2 := NewDefAssemblyImpl(rawAss2, "all", nil)
	if err2 == nil {
		t.Fatal("Expected error for duplicated enum values, but got nil")
	}
	if err2.Error() != "enum 'TestEnum2' has duplicated item value: 1 (item: B)" {
		t.Fatalf("Unexpected error message: %v", err2)
	}
}

func TestDefAssemblyImpl_ValidateTable(t *testing.T) {
	rawAss := rawdefs.NewRawAssembly()
	rawAss.Beans = append(rawAss.Beans, &rawdefs.RawBean{
		Name: "TestBean",
		Fields: []*rawdefs.RawField{
			{Name: "Id", Type: "int"},
		},
	})
	rawAss.Tables = append(rawAss.Tables, &rawdefs.RawTable{
		Name: "TestTable",
		ValueType: "TestBean",
		Index: "NotExistId", // 不存在的索引
	})

	_, err := NewDefAssemblyImpl(rawAss, "all", nil)
	if err == nil {
		t.Fatal("Expected error for invalid table index, but got nil")
	}
	if err.Error() != "table 'TestTable' index field 'NotExistId' not found in Bean 'TestBean'" {
		t.Fatalf("Unexpected error message: %v", err)
	}
}