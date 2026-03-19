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
