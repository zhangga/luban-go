package codetarget

import (
	"strings"
	"testing"

	"github.com/zhangga/luban-go/core/defs"
	"github.com/zhangga/luban-go/core/rawdefs"
)

func TestTSTarget(t *testing.T) {
	target, err := NewTSTarget("../../templates/ts")
	if err != nil {
		t.Fatalf("Failed to create TSTarget: %v", err)
	}

	rawAss := rawdefs.NewRawAssembly()
	rawAss.Beans = append(rawAss.Beans, &rawdefs.RawBean{
		Name: "TestItemTS",
		Fields: []*rawdefs.RawField{
			{Name: "Id", Type: "int"},
			{Name: "Name", Type: "string"},
			{Name: "PriceList", Type: "list,float"},
		},
	})
	assembly, _ := defs.NewDefAssemblyImpl(rawAss, "all", nil)
	beanImpl := assembly.GetType("TestItemTS").(*defs.DefBeanImpl)
	
	for _, f := range rawAss.Beans[0].Fields {
		df := defs.NewDefField(f)
		beanImpl.Fields = append(beanImpl.Fields, df)
		beanImpl.HierarchyFields = append(beanImpl.HierarchyFields, df)
	}

	out, err := target.GenerateBean(beanImpl, []string{})
	if err != nil {
		t.Fatalf("Failed to generate bean: %v", err)
	}
	
	tsStr := string(out)
	
	if !strings.Contains(tsStr, "export class TestItemTS") {
		t.Errorf("Missing class TestItemTS: %s", tsStr)
	}
	if !strings.Contains(tsStr, "public readonly Id: number;") {
		t.Errorf("Missing field Id: %s", tsStr)
	}
	if !strings.Contains(tsStr, "public readonly Name: string;") {
		t.Errorf("Missing field Name: %s", tsStr)
	}
	if !strings.Contains(tsStr, "public readonly PriceList: number[];") {
		t.Errorf("Missing field PriceList: %s", tsStr)
	}
}

func TestTSTypeMapper(t *testing.T) {
	tests := map[string]string{
		"int": "number",
		"string": "string",
		"bool": "boolean",
		"list,int": "number[]",
		"map,string,float": "Map<string, number>",
		"set,int": "Set<number>",
		"array,string": "string[]",
		"demo.Item": "demo.Item",
	}

	for in, expected := range tests {
		actual := tsTypeMapper(in)
		if actual != expected {
			t.Errorf("tsTypeMapper(%s) = %s, expected %s", in, actual, expected)
		}
	}
}