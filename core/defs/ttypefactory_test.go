package defs

import (
	"github.com/zhangga/luban-go/core/rawdefs"
	"github.com/zhangga/luban-go/core/types"
	"testing"
)

func TestTypeFactory(t *testing.T) {
	rawAss := rawdefs.NewRawAssembly()

	// Add an Enum
	rawAss.Enums = append(rawAss.Enums, &rawdefs.RawEnum{
		Namespace: "demo",
		Name:      "Gender",
	})

	// Add a Bean
	rawAss.Beans = append(rawAss.Beans, &rawdefs.RawBean{
		Namespace: "demo",
		Name:      "Item",
	})

	assembly, err := NewDefAssemblyImpl(rawAss, "all", nil)
	if err != nil {
		t.Fatalf("Failed to create assembly: %v", err)
	}

	factory := NewTTypeFactory(assembly)

	// Test primitive types
	tType, err := factory.CreateType("int")
	if err != nil || tType.TypeName() != "int" {
		t.Errorf("Failed to parse int")
	}

	// Test collection types
	tType, err = factory.CreateType("list, string")
	if err != nil || !tType.IsCollection() || tType.TypeName() != "list" {
		t.Errorf("Failed to parse list, string")
	}
	if tType.ElementType().TypeName() != "string" {
		t.Errorf("List element should be string")
	}

	tType, err = factory.CreateType("map, int, demo.Item")
	if err != nil || !tType.IsCollection() || tType.TypeName() != "map" {
		t.Fatalf("Failed to parse map, int, demo.Item: %v", err)
	}

	mapType := tType.(*types.TMap)
	if mapType.KeyType().TypeName() != "int" {
		t.Errorf("Map key should be int")
	}
	if mapType.ValueType().TypeName() != "bean" {
		t.Errorf("Map value should be bean, got %s", mapType.ValueType().TypeName())
	}
}
