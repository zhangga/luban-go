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
	f.SetCellValue(sheet, "C1", "Desc")

	// Data
	f.SetCellValue(sheet, "A2", "1001")
	f.SetCellValue(sheet, "B2", "Excalibur")
	f.SetCellValue(sheet, "C2", "Legendary Sword")

	f.SetCellValue(sheet, "A3", "1002")
	f.SetCellValue(sheet, "B3", "Iron Shield")
	f.SetCellValue(sheet, "C3", "Basic Defense")

	if err := f.SaveAs("test_project/Datas/TbItem.xlsx"); err != nil {
		panic(err)
	}
}
