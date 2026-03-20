package dataloader

import (
	"bytes"
	"testing"

	"github.com/xuri/excelize/v2"
	"github.com/zhangga/luban-go/core/defs"
	"github.com/zhangga/luban-go/core/rawdefs"
	"github.com/zhangga/luban-go/core/types"
)

func createTestExcel() []byte {
	f := excelize.NewFile()
	sheet := "Sheet1"
	f.SetCellValue(sheet, "A1", "Id")
	f.SetCellValue(sheet, "B1", "Name")

	f.SetCellValue(sheet, "A2", "101")
	f.SetCellValue(sheet, "B2", "Sword")

	f.SetCellValue(sheet, "A3", "102")
	f.SetCellValue(sheet, "B3", "Shield")

	var buf bytes.Buffer
	f.Write(&buf)
	return buf.Bytes()
}

func TestExcelDataLoader(t *testing.T) {
	excelData := createTestExcel()

	// 构造一个模拟的装配体和工厂
	rawAss := rawdefs.NewRawAssembly()
	rawAss.Beans = append(rawAss.Beans, &rawdefs.RawBean{
		Name: "TestItem",
		Fields: []*rawdefs.RawField{
			{Name: "Id", Type: "int"},
			{Name: "Name", Type: "string"},
		},
	})
	assembly, _ := defs.NewDefAssemblyImpl(rawAss, "all", nil)
	typeFactory := defs.NewTTypeFactory(assembly)

	loader := NewExcelDataLoader(typeFactory)
	err := loader.Load("test.xlsx", "", bytes.NewReader(excelData))
	if err != nil {
		t.Fatalf("Failed to load excel: %v", err)
	}

	if len(loader.sheets) != 1 {
		t.Fatalf("Expected 1 sheet, got %d", len(loader.sheets))
	}

	sheet := loader.sheets[0]
	if len(sheet.Cells) != 3 { // 1 header + 2 data rows
		t.Errorf("Expected 3 rows, got %d", len(sheet.Cells))
	}

	// Read records
	beanImpl := assembly.GetType("TestItem").(*defs.DefBeanImpl)
	dummyBeanType := types.NewTBean(false, beanImpl, nil)

	records := loader.ReadMulti(dummyBeanType)

	if len(records) != 2 {
		t.Fatalf("Expected 2 records, got %d", len(records))
	}

	// 检查第一条记录的数据
	if len(records[0].Data.Fields) != 2 {
		t.Errorf("Expected 2 fields in record, got %d", len(records[0].Data.Fields))
	}

	if records[0].Data.Fields[0].String() != "101" {
		t.Errorf("Expected ID 101, got %s", records[0].Data.Fields[0].String())
	}
	if records[0].Data.Fields[1].String() != "Sword" {
		t.Errorf("Expected Name Sword, got %s", records[0].Data.Fields[1].String())
	}
}

func createAdvancedTestExcel() []byte {
	f := excelize.NewFile()
	sheet := "Sheet1"
	
	// ## 控制行
	f.SetCellValue(sheet, "A1", "##var")
	f.SetCellValue(sheet, "B1", "Id")
	f.SetCellValue(sheet, "C1", "Name")
	f.SetCellValue(sheet, "D1", "List") // 同名列1
	f.SetCellValue(sheet, "E1", "List") // 同名列2
	f.SetCellValue(sheet, "F1", "#IgnoreCol")

	f.SetCellValue(sheet, "A2", "##type")
	f.SetCellValue(sheet, "B2", "int")
	f.SetCellValue(sheet, "C2", "string")
	f.SetCellValue(sheet, "D2", "list,int")
	f.SetCellValue(sheet, "E2", "")
	f.SetCellValue(sheet, "F2", "")

	// 正常数据行
	f.SetCellValue(sheet, "A3", "")
	f.SetCellValue(sheet, "B3", "101")
	f.SetCellValue(sheet, "C3", "Sword")
	f.SetCellValue(sheet, "D3", "10")
	f.SetCellValue(sheet, "E3", "20")
	f.SetCellValue(sheet, "F3", "ignored")

	// 忽略行
	f.SetCellValue(sheet, "A4", "##")
	f.SetCellValue(sheet, "B4", "102")
	
	f.SetCellValue(sheet, "A5", "#")
	f.SetCellValue(sheet, "B5", "103")

	var buf bytes.Buffer
	f.Write(&buf)
	return buf.Bytes()
}

func TestExcelDataLoader_Advanced(t *testing.T) {
	excelData := createAdvancedTestExcel()

	// 构造一个模拟的装配体和工厂
	rawAss := rawdefs.NewRawAssembly()
	rawAss.Beans = append(rawAss.Beans, &rawdefs.RawBean{
		Name: "TestItemAdv",
		Fields: []*rawdefs.RawField{
			{Name: "Id", Type: "int"},
			{Name: "Name", Type: "string"},
			{Name: "List", Type: "list,int"},
		},
	})
	assembly, _ := defs.NewDefAssemblyImpl(rawAss, "all", nil)
	typeFactory := defs.NewTTypeFactory(assembly)

	loader := NewExcelDataLoader(typeFactory)
	err := loader.Load("test.xlsx", "", bytes.NewReader(excelData))
	if err != nil {
		t.Fatalf("Failed to load excel: %v", err)
	}

	beanImpl := assembly.GetType("TestItemAdv").(*defs.DefBeanImpl)
	// 将 RawField 转换为 DefField
	for _, f := range rawAss.Beans[0].Fields {
		beanImpl.Fields = append(beanImpl.Fields, defs.NewDefField(f))
		beanImpl.HierarchyFields = append(beanImpl.HierarchyFields, defs.NewDefField(f))
	}
	dummyBeanType := types.NewTBean(false, beanImpl, nil)

	records := loader.ReadMulti(dummyBeanType)

	// 应该只有 1 条记录（ID=101的那条），102被 ## 忽略，103被 # 忽略
	if len(records) != 1 {
		t.Fatalf("Expected 1 record, got %d", len(records))
	}

	// 检查 List 合并
	listVal := records[0].Data.Fields[2]
	if listVal.String() != "[10, 20]" {
		t.Errorf("Expected merged list [10, 20], got %s", listVal.String())
	}
}
