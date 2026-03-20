package codetarget

import (
	"strings"
	"testing"

	"github.com/zhangga/luban-go/core/defs"
	"github.com/zhangga/luban-go/core/rawdefs"
)

func TestPbTarget(t *testing.T) {
	// 创建一个临时的 TemplateDir 以供测试，我们需要使用到 ../../templates/pb
	target, err := NewPbTarget("../../templates/pb")
	if err != nil {
		t.Fatalf("Failed to create PbTarget: %v", err)
	}

	// Mock assembly
	rawAss := rawdefs.NewRawAssembly()
	rawAss.Beans = append(rawAss.Beans, &rawdefs.RawBean{
		Name: "TestItemPb",
		Fields: []*rawdefs.RawField{
			{Name: "Id", Type: "int"},
			{Name: "Name", Type: "string"},
			{Name: "PriceList", Type: "list,float"},
			{Name: "PropMap", Type: "map,string,int"},
		},
	})
	assembly, _ := defs.NewDefAssemblyImpl(rawAss, "all", nil)
	beanImpl := assembly.GetType("TestItemPb").(*defs.DefBeanImpl)
	
	for _, f := range rawAss.Beans[0].Fields {
		df := defs.NewDefField(f)
		beanImpl.Fields = append(beanImpl.Fields, df)
		beanImpl.HierarchyFields = append(beanImpl.HierarchyFields, df)
	}

	// 1. 测试 Bean 导出
	out, err := target.GenerateBean(beanImpl, []string{})
	if err != nil {
		t.Fatalf("Failed to generate bean: %v", err)
	}
	
	pbStr := string(out)
	
	if !strings.Contains(pbStr, "syntax = \"proto3\";") {
		t.Errorf("Missing proto3 syntax")
	}
	if !strings.Contains(pbStr, "message TestItemPb") {
		t.Errorf("Missing message TestItemPb")
	}
	if !strings.Contains(pbStr, "int32 id = 1;") {
		t.Errorf("Missing field id: %s", pbStr)
	}
	if !strings.Contains(pbStr, "string name = 2;") {
		t.Errorf("Missing field name: %s", pbStr)
	}
	if !strings.Contains(pbStr, "repeated float pricelist = 3;") {
		t.Errorf("Missing field pricelist: %s", pbStr)
	}
	if !strings.Contains(pbStr, "map<string, int32> propmap = 4;") {
		t.Errorf("Missing field propmap: %s", pbStr)
	}

	// 2. 测试 TypeMapper
	if pbTypeMapper("int") != "int32" {
		t.Errorf("pbTypeMapper int failed")
	}
	if pbTypeMapper("list,string") != "repeated string" {
		t.Errorf("pbTypeMapper list,string failed")
	}
	if pbTypeMapper("map,string,float") != "map<string, float>" {
		t.Errorf("pbTypeMapper map failed: %s", pbTypeMapper("map,string,float"))
	}
	if pbTypeMapper("demo.Item") != "demo_Item" {
		t.Errorf("pbTypeMapper custom type failed: %s", pbTypeMapper("demo.Item"))
	}
}
