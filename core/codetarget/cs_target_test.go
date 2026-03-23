package codetarget

import (
	"strings"
	"testing"

	"github.com/zhangga/luban-go/core/defs"
	"github.com/zhangga/luban-go/core/rawdefs"
)

func TestCSTarget(t *testing.T) {
	target, err := NewCSTarget("../../templates/cs")
	if err != nil {
		t.Fatalf("Failed to create CSTarget: %v", err)
	}

	rawAss := rawdefs.NewRawAssembly()
	rawAss.Beans = append(rawAss.Beans, &rawdefs.RawBean{
		Name: "TestItemCS",
		Fields: []*rawdefs.RawField{
			{Name: "Id", Type: "int"},
			{Name: "Name", Type: "string"},
			{Name: "PriceList", Type: "list,float"},
		},
	})
	assembly, _ := defs.NewDefAssemblyImpl(rawAss, "all", nil)
	beanImpl := assembly.GetType("TestItemCS").(*defs.DefBeanImpl)
	
	for _, f := range rawAss.Beans[0].Fields {
		df := defs.NewDefField(f)
		beanImpl.Fields = append(beanImpl.Fields, df)
		beanImpl.HierarchyFields = append(beanImpl.HierarchyFields, df)
	}

	out, err := target.GenerateBean(beanImpl, []string{})
	if err != nil {
		t.Fatalf("Failed to generate bean: %v", err)
	}
	
	csStr := string(out)
	
	if !strings.Contains(csStr, "public class TestItemCS") {
		t.Errorf("Missing class TestItemCS: %s", csStr)
	}
	if !strings.Contains(csStr, "public readonly int Id;") {
		t.Errorf("Missing field Id: %s", csStr)
	}
	if !strings.Contains(csStr, "public readonly string Name;") {
		t.Errorf("Missing field Name: %s", csStr)
	}
	if !strings.Contains(csStr, "public readonly System.Collections.Generic.List<float> PriceList;") {
		t.Errorf("Missing field PriceList: %s", csStr)
	}
}

func TestCSTypeMapper(t *testing.T) {
	tests := map[string]string{
		"int": "int",
		"string": "string",
		"bool": "bool",
		"list,int": "System.Collections.Generic.List<int>",
		"map,string,float": "System.Collections.Generic.Dictionary<string, float>",
		"set,int": "System.Collections.Generic.HashSet<int>",
		"array,string": "string[]",
		"demo.Item": "demo.Item",
	}

	for in, expected := range tests {
		actual := csTypeMapper(in)
		if actual != expected {
			t.Errorf("csTypeMapper(%s) = %s, expected %s", in, actual, expected)
		}
	}
}