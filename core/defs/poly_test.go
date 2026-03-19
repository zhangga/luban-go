package defs

import (
	"github.com/zhangga/luban-go/core/rawdefs"
	"testing"
)

func TestPolymorphicBean(t *testing.T) {
	rawAss := rawdefs.NewRawAssembly()

	shapeBean := &rawdefs.RawBean{
		Namespace: "demo",
		Name:      "Shape",
		Fields: []*rawdefs.RawField{
			{Name: "Id", Type: "int"},
		},
	}

	circleBean := &rawdefs.RawBean{
		Namespace: "demo",
		Name:      "Circle",
		Parent:    "demo.Shape",
		Fields: []*rawdefs.RawField{
			{Name: "Radius", Type: "float"},
		},
	}

	rawAss.Beans = append(rawAss.Beans, shapeBean, circleBean)

	assembly, err := NewDefAssemblyImpl(rawAss, "all", nil)
	if err != nil {
		t.Fatalf("Failed to create assembly: %v", err)
	}

	shapeDef := assembly.GetType("demo.Shape").(*DefBeanImpl)
	circleDef := assembly.GetType("demo.Circle").(*DefBeanImpl)

	// 测试
	if shapeDef.TryGetChild("Circle") == nil {
		t.Errorf("Expected to find child Circle")
	}
	if len(shapeDef.HierarchyNotAbstractChildren()) != 1 {
		t.Errorf("Expected 1 non abstract child")
	}

	if len(circleDef.HierarchyFields) != 2 {
		t.Errorf("Expected 2 fields in Circle, got %d", len(circleDef.HierarchyFields))
	}
}
