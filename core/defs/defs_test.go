package defs

import (
	"github.com/zhangga/luban-go/core/rawdefs"
	"testing"
)

func TestDefs(t *testing.T) {
	rawBean := &rawdefs.RawBean{
		Namespace: "TestNs",
		Name:      "TestBean",
	}

	bean := NewDefBeanImpl(rawBean)
	if bean.FullName() != "TestNs.TestBean" {
		t.Errorf("Expected TestNs.TestBean, got %s", bean.FullName())
	}

	rawEnum := &rawdefs.RawEnum{
		Namespace: "TestNs",
		Name:      "TestEnum",
	}

	enum := NewDefEnumImpl(rawEnum)
	if enum.FullName() != "TestNs.TestEnum" {
		t.Errorf("Expected TestNs.TestEnum, got %s", enum.FullName())
	}

	rawTable := &rawdefs.RawTable{
		Namespace: "TestNs",
		Name:      "TestTable",
	}

	table := NewDefTable(rawTable)
	if table.FullName() != "TestNs.TestTable" {
		t.Errorf("Expected TestNs.TestTable, got %s", table.FullName())
	}

	assembly := NewDefAssembly(rawdefs.NewRawAssembly())
	assembly.AddType(bean)
	assembly.AddType(enum)
	assembly.AddTable(table)

	if len(assembly.Types) != 2 {
		t.Errorf("Expected 2 types, got %d", len(assembly.Types))
	}
	if len(assembly.Tables) != 1 {
		t.Errorf("Expected 1 table, got %d", len(assembly.Tables))
	}
}
