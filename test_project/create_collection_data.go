package main

import (
	"github.com/xuri/excelize/v2"
)

func main() {
	f := excelize.NewFile()
	sheet := "Sheet1"

	// Headers
	f.SetCellValue(sheet, "A1", "Id")
	f.SetCellValue(sheet, "B1", "Name")
	f.SetCellValue(sheet, "C1", "PriceList") // 测试 List
	f.SetCellValue(sheet, "D1", "PropMap")   // 测试 Map

	// Data
	f.SetCellValue(sheet, "A2", "1001")
	f.SetCellValue(sheet, "B2", "Bag")
	f.SetCellValue(sheet, "C2", "10,20,30")     // 逗号分隔的int list
	f.SetCellValue(sheet, "D2", "hp,100,mp,50") // 逗号分隔的 string:int map

	if err := f.SaveAs("test_project/Datas/TbBag.xlsx"); err != nil {
		panic(err)
	}
}
