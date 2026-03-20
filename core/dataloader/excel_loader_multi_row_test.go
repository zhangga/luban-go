package dataloader

import (
	"bytes"
	"testing"

	"github.com/xuri/excelize/v2"
	"github.com/zhangga/luban-go/core/defs"
	"github.com/zhangga/luban-go/core/rawdefs"
	"github.com/zhangga/luban-go/core/types"
)

func createMultiRowTestExcel() []byte {
	f := excelize.NewFile()
	sheet := "Sheet1"
	
	// ## 控制行
	f.SetCellValue(sheet, "A1", "##var")
	f.SetCellValue(sheet, "B1", "Id")
	f.SetCellValue(sheet, "C1", "Name")
	f.SetCellValue(sheet, "D1", "Skills")

	f.SetCellValue(sheet, "A2", "##type")
	f.SetCellValue(sheet, "B2", "int")
	f.SetCellValue(sheet, "C2", "string")
	f.SetCellValue(sheet, "D2", "list,string")

	// 记录1: Id=101
	f.SetCellValue(sheet, "A3", "")
	f.SetCellValue(sheet, "B3", "101")
	f.SetCellValue(sheet, "C3", "Warrior")
	f.SetCellValue(sheet, "D3", "Slash")

	// 记录1的第2行 (Id和Name为空)
	f.SetCellValue(sheet, "A4", "")
	f.SetCellValue(sheet, "B4", "")
	f.SetCellValue(sheet, "C4", "")
	f.SetCellValue(sheet, "D4", "Charge")

	// 记录1的第3行
	f.SetCellValue(sheet, "A5", "")
	f.SetCellValue(sheet, "B5", "")
	f.SetCellValue(sheet, "C5", "")
	f.SetCellValue(sheet, "D5", "Block")

	// 记录2: Id=102
	f.SetCellValue(sheet, "A6", "")
	f.SetCellValue(sheet, "B6", "102")
	f.SetCellValue(sheet, "C6", "Mage")
	f.SetCellValue(sheet, "D6", "Fireball")

	// 记录2的第2行
	f.SetCellValue(sheet, "A7", "")
	f.SetCellValue(sheet, "B7", "")
	f.SetCellValue(sheet, "C7", "")
	f.SetCellValue(sheet, "D7", "IceBlock")

	var buf bytes.Buffer
	f.Write(&buf)
	return buf.Bytes()
}

func TestExcelDataLoader_MultiRow(t *testing.T) {
	excelData := createMultiRowTestExcel()

	// 构造一个模拟的装配体和工厂
	rawAss := rawdefs.NewRawAssembly()
	rawAss.Beans = append(rawAss.Beans, &rawdefs.RawBean{
		Name: "TestHero",
		Fields: []*rawdefs.RawField{
			{Name: "Id", Type: "int"},
			{Name: "Name", Type: "string"},
			{Name: "Skills", Type: "list,string"},
		},
	})
	assembly, _ := defs.NewDefAssemblyImpl(rawAss, "all", nil)
	typeFactory := defs.NewTTypeFactory(assembly)

	loader := NewExcelDataLoader(typeFactory)
	err := loader.Load("test.xlsx", "", bytes.NewReader(excelData))
	if err != nil {
		t.Fatalf("Failed to load excel: %v", err)
	}

	beanImpl := assembly.GetType("TestHero").(*defs.DefBeanImpl)
	// 将 RawField 转换为 DefField
	for _, f := range rawAss.Beans[0].Fields {
		beanImpl.Fields = append(beanImpl.Fields, defs.NewDefField(f))
		beanImpl.HierarchyFields = append(beanImpl.HierarchyFields, defs.NewDefField(f))
	}
	dummyBeanType := types.NewTBean(false, beanImpl, nil)

	records := loader.ReadMulti(dummyBeanType)

	// 应该只有 2 条记录（ID=101, 102）
	if len(records) != 2 {
		t.Fatalf("Expected 2 records, got %d", len(records))
	}

	// 检查第一条记录合并
	r1 := records[0]
	if r1.Data.Fields[0].String() != "101" {
		t.Errorf("Expected 101, got %s", r1.Data.Fields[0].String())
	}
	if r1.Data.Fields[2].String() != "[Slash, Charge, Block]" {
		t.Errorf("Expected [Slash, Charge, Block], got %s", r1.Data.Fields[2].String())
	}

	// 检查第二条记录合并
	r2 := records[1]
	if r2.Data.Fields[0].String() != "102" {
		t.Errorf("Expected 102, got %s", r2.Data.Fields[0].String())
	}
	if r2.Data.Fields[2].String() != "[Fireball, IceBlock]" {
		t.Errorf("Expected [Fireball, IceBlock], got %s", r2.Data.Fields[2].String())
	}
}