package defs

import (
	"github.com/zhangga/luban-go/core/rawdefs"
	"testing"
)

func TestDefAssemblyImplConversion(t *testing.T) {
	rawAss := rawdefs.NewRawAssembly()

	rawAss.Targets = append(rawAss.Targets, &rawdefs.RawTarget{
		Name:   "test-target",
		Groups: []string{"default"},
	})

	rawAss.Enums = append(rawAss.Enums, &rawdefs.RawEnum{
		Namespace: "demo",
		Name:      "RoleType",
	})

	rawAss.Beans = append(rawAss.Beans, &rawdefs.RawBean{
		Namespace: "demo",
		Name:      "Role",
	})

	rawAss.Tables = append(rawAss.Tables, &rawdefs.RawTable{
		Namespace: "demo",
		Name:      "TbRole",
	})

	assembly, err := NewDefAssemblyImpl(rawAss, "test-target", []string{"demo.TbRole"})
	if err != nil {
		t.Fatalf("Failed to create assembly: %v", err)
	}

	if assembly.Target.Name != "test-target" {
		t.Errorf("Target name mismatch. Expected test-target, got %s", assembly.Target.Name)
	}

	if len(assembly.TypeList) != 3 {
		t.Errorf("Expected 3 types (Enum, Bean, Table), got %d", len(assembly.TypeList))
	}

	if len(assembly.TablesByName) != 1 {
		t.Errorf("Expected 1 table in map, got %d", len(assembly.TablesByName))
	}

	if len(assembly.ExportTables) != 1 {
		t.Errorf("Expected 1 export table, got %d", len(assembly.ExportTables))
	}

	roleType := assembly.GetType("demo.Role")
	if roleType == nil {
		t.Errorf("Failed to find demo.Role in assembly")
	}

	if roleType.Name() != "Role" || roleType.Namespace() != "demo" {
		t.Errorf("Role metadata mismatch")
	}
}
