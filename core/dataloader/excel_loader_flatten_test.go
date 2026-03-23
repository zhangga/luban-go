package dataloader

import (
	"bytes"
	"testing"

	"github.com/xuri/excelize/v2"
	"github.com/zhangga/luban-go/core/defs"
	"github.com/zhangga/luban-go/core/rawdefs"
	"github.com/zhangga/luban-go/core/types"
)

func createFlattenTestExcel() []byte {
	f := excelize.NewFile()
	sheet := "Sheet1"
	
	// ## 控制行
	f.SetCellValue(sheet, "A1", "##var")
	f.SetCellValue(sheet, "B1", "Id")
	f.SetCellValue(sheet, "C1", "Info.Name")
	f.SetCellValue(sheet, "D1", "Info.Level")
	f.SetCellValue(sheet, "E1", "Price")

	f.SetCellValue(sheet, "A2", "##type")
	f.SetCellValue(sheet, "B2", "int")
	f.SetCellValue(sheet, "C2", "string")
	f.SetCellValue(sheet, "D2", "int")
	f.SetCellValue(sheet, "E2", "int")

	// 记录1
	f.SetCellValue(sheet, "A3", "")
	f.SetCellValue(sheet, "B3", "101")
	f.SetCellValue(sheet, "C3", "Sword")
	f.SetCellValue(sheet, "D3", "5")
	f.SetCellValue(sheet, "E3", "100")

	// 记录2
	f.SetCellValue(sheet, "A4", "")
	f.SetCellValue(sheet, "B4", "102")
	f.SetCellValue(sheet, "C4", "Shield")
	f.SetCellValue(sheet, "D4", "10")
	f.SetCellValue(sheet, "E4", "200")

	var buf bytes.Buffer
	f.Write(&buf)
	return buf.Bytes()
}

func TestExcelDataLoader_Flatten(t *testing.T) {
	excelData := createFlattenTestExcel()

	// 构造装配体和工厂
	rawAss := rawdefs.NewRawAssembly()
	
	// 内部结构体
	itemInfoBean := &rawdefs.RawBean{
		Name: "ItemInfo",
		Fields: []*rawdefs.RawField{
			{Name: "Name", Type: "string"},
			{Name: "Level", Type: "int"},
		},
	}
	rawAss.Beans = append(rawAss.Beans, itemInfoBean)

	// 主结构体
	mainBean := &rawdefs.RawBean{
		Name: "TestFlattenItem",
		Fields: []*rawdefs.RawField{
			{Name: "Id", Type: "int"},
			{Name: "Info", Type: "ItemInfo"}, // 嵌套的 Bean
			{Name: "Price", Type: "int"},
		},
	}
	rawAss.Beans = append(rawAss.Beans, mainBean)

	assembly, _ := defs.NewDefAssemblyImpl(rawAss, "all", nil)
	typeFactory := defs.NewTTypeFactory(assembly)

	loader := NewExcelDataLoader(typeFactory)
	err := loader.Load("test.xlsx", "", bytes.NewReader(excelData))
	if err != nil {
		t.Fatalf("Failed to load excel: %v", err)
	}

	mainBeanImpl := assembly.GetType("TestFlattenItem").(*defs.DefBeanImpl)
	dummyBeanType := types.NewTBean(false, mainBeanImpl, nil)

	records := loader.ReadMulti(dummyBeanType)

	// 应该有 2 条记录
	if len(records) != 2 {
		t.Fatalf("Expected 2 records, got %d", len(records))
	}

	// 检查第一条记录的数据
	// Id
	if records[0].Data.Fields[0].String() != "101" {
		t.Errorf("Expected ID 101, got %s", records[0].Data.Fields[0].String())
	}
	
	// Info (这是一个 DBean)
	infoData := records[0].Data.Fields[1]
	if infoData.TypeName() != "bean" {
		t.Fatalf("Expected bean type for Info, got %s", infoData.TypeName())
	}
	
	// 这里不一定有直接的方法拿里面的字段值，我们可以转成 json 格式来比较
	// 或者直接读取 field
	// Info.Name
	if infoData.String() != "{Sword, 5}" { // dbean 的 String 默认输出格式
		t.Errorf("Expected nested bean string {Sword, 5}, got %s", infoData.String())
	}

	// Price
	if records[0].Data.Fields[2].String() != "100" {
		t.Errorf("Expected Price 100, got %s", records[0].Data.Fields[2].String())
	}
}