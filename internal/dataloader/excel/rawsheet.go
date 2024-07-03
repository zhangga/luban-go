package excel

type RawSheet struct {
	Title     *Title
	TableName string
	SheetName string
	Cells     [][]*Cell
}
