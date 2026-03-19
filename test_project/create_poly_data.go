package main

import (
	"github.com/xuri/excelize/v2"
)

func main() {
	f := excelize.NewFile()
	sheet := "Sheet1"

	// Headers (包括多态字段 $type)
	f.SetCellValue(sheet, "A1", "Id")
	f.SetCellValue(sheet, "B1", "$type")
	f.SetCellValue(sheet, "C1", "Radius")
	f.SetCellValue(sheet, "D1", "Width")
	f.SetCellValue(sheet, "E1", "Height")

	// Data: Circle
	f.SetCellValue(sheet, "A2", "101")
	f.SetCellValue(sheet, "B2", "Circle")
	f.SetCellValue(sheet, "C2", "5.5")
	f.SetCellValue(sheet, "D2", "")
	f.SetCellValue(sheet, "E2", "")

	// Data: Rectangle
	f.SetCellValue(sheet, "A3", "102")
	f.SetCellValue(sheet, "B3", "Rectangle")
	f.SetCellValue(sheet, "C3", "")
	f.SetCellValue(sheet, "D3", "10.0")
	f.SetCellValue(sheet, "E3", "20.0")

	if err := f.SaveAs("test_project/Datas/TbShape.xlsx"); err != nil {
		panic(err)
	}
}
