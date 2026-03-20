package dataloader

import (
	"bytes"
	"testing"
	"strings"

	"github.com/xuri/excelize/v2"
	"github.com/zhangga/luban-go/core/defs"
	"github.com/zhangga/luban-go/core/rawdefs"
	"github.com/zhangga/luban-go/core/types"
	"github.com/zhangga/luban-go/core/l10n"
	"github.com/zhangga/luban-go/core/datas"
)

func createL10nTestExcel() []byte {
	f := excelize.NewFile()
	sheet := "Sheet1"
	
	f.SetCellValue(sheet, "A1", "Id")
	f.SetCellValue(sheet, "B1", "Name")
	f.SetCellValue(sheet, "C1", "Desc")

	f.SetCellValue(sheet, "A2", "101")
	f.SetCellValue(sheet, "B2", "Sword")
	f.SetCellValue(sheet, "C2", "A sharp sword")

	f.SetCellValue(sheet, "A3", "102")
	f.SetCellValue(sheet, "B3", "Shield")
	f.SetCellValue(sheet, "C3", "A strong shield")

	var buf bytes.Buffer
	f.Write(&buf)
	return buf.Bytes()
}

func TestExcelDataLoader_L10nText(t *testing.T) {
	excelData := createL10nTestExcel()

	// Clear l10n manager before test
	l10n.GetManager().Clear()

	rawAss := rawdefs.NewRawAssembly()
	rawAss.Beans = append(rawAss.Beans, &rawdefs.RawBean{
		Name: "TestItemL10n",
		Fields: []*rawdefs.RawField{
			{Name: "Id", Type: "int"},
			{Name: "Name", Type: "string"},
			{Name: "Desc", Type: "text"},
		},
	})
	assembly, _ := defs.NewDefAssemblyImpl(rawAss, "all", nil)
	typeFactory := defs.NewTTypeFactory(assembly)

	loader := NewExcelDataLoader(typeFactory)
	err := loader.Load("test.xlsx", "", bytes.NewReader(excelData))
	if err != nil {
		t.Fatalf("Failed to load excel: %v", err)
	}

	beanImpl := assembly.GetType("TestItemL10n").(*defs.DefBeanImpl)
	// 将 RawField 转换为 DefField
	for _, f := range rawAss.Beans[0].Fields {
		beanImpl.Fields = append(beanImpl.Fields, defs.NewDefField(f))
		beanImpl.HierarchyFields = append(beanImpl.HierarchyFields, defs.NewDefField(f))
	}
	dummyBeanType := types.NewTBean(false, beanImpl, nil)

	records := loader.ReadMulti(dummyBeanType)

	if len(records) != 2 {
		t.Fatalf("Expected 2 records, got %d", len(records))
	}

	// 检查第一条记录的数据
	desc1 := records[0].Data.Fields[2]
	dtext1, ok := desc1.(*datas.DText)
	if !ok {
		t.Fatalf("Expected DText, got %T", desc1)
	}

	if dtext1.RawText != "A sharp sword" {
		t.Errorf("Expected 'A sharp sword', got '%s'", dtext1.RawText)
	}
	if dtext1.Key == "" {
		t.Errorf("Expected key not to be empty")
	}

	// 导出 L10n 数据检查
	bytes, err := l10n.GetManager().Export()
	if err != nil {
		t.Fatalf("Failed to export l10n: %v", err)
	}
	
	jsonStr := string(bytes)
	if !strings.Contains(jsonStr, "A sharp sword") || !strings.Contains(jsonStr, "A strong shield") {
		t.Errorf("Exported json missing expected texts: %s", jsonStr)
	}
}
