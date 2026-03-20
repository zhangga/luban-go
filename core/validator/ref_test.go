package validator

import (
	"testing"

	"github.com/zhangga/luban-go/core/datas"
	"github.com/zhangga/luban-go/core/defs"
	"github.com/zhangga/luban-go/core/rawdefs"
	"github.com/zhangga/luban-go/core/types"
)

func TestRefValidator(t *testing.T) {
	// 准备假的上下文
	rawAss := rawdefs.NewRawAssembly()
	rawAss.Tables = append(rawAss.Tables, &rawdefs.RawTable{
		Namespace: "demo",
		Name:      "TbItem",
		Index:     "Id",
		ValueType: "demo.Item",
	})
	
	rawAss.Beans = append(rawAss.Beans, &rawdefs.RawBean{
		Namespace: "demo",
		Name:      "Item",
		Fields: []*rawdefs.RawField{
			{Name: "Id", Type: "int"},
		},
	})

	assembly, _ := defs.NewDefAssemblyImpl(rawAss, "all", nil)

	// 模拟数据记录
	tBeanImpl := assembly.Types["demo.Item"].(*defs.DefBeanImpl)
	tBeanType := types.NewTBean(false, tBeanImpl, nil)
	dBean1 := datas.NewDBean(tBeanType, tBeanType, []datas.DType{datas.NewDInt(1001)})
	dBean2 := datas.NewDBean(tBeanType, tBeanType, []datas.DType{datas.NewDInt(1002)})

	records := []*defs.Record{
		defs.NewRecord(dBean1, "", nil),
		defs.NewRecord(dBean2, "", nil),
	}

	tableMap := map[string][]*defs.Record{
		"demo.TbItem": records,
	}

	ctx := NewValidatorContext(assembly, tableMap)
	v := &RefValidator{}

	// Test valid ref
	err := v.Validate(ctx, datas.NewDInt(1001), "TbItem")
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}

	// Test invalid ref
	err = v.Validate(ctx, datas.NewDInt(9999), "TbItem")
	if err == nil {
		t.Errorf("Expected error for non-existent ref")
	}

	// Test default/empty value (should pass even if not in target table)
	err = v.Validate(ctx, datas.NewDInt(0), "TbItem")
	if err != nil {
		t.Errorf("Expected nil for default value, got %v", err)
	}
}
