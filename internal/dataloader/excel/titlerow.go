package excel

type TitleRow struct {
	Tags      []string
	SelfTitle *Title
	Row       []*Cell
	Rows      [][]*Cell
	Elements  []*TitleRow
	Fields    map[string]*TitleRow
}

func (t *TitleRow) AsStream(sep string) *ExcelStream {
	return NewExcelStream(t.Row, t.SelfTitle.FromIndex, t.SelfTitle.ToIndex, sep, t.SelfTitle.Default)
}

func (t *TitleRow) AsMultiRowConcatStream(sep string) *ExcelStream {
	return NewExcelStreamByRows(t.Rows, t.SelfTitle.FromIndex, t.SelfTitle.ToIndex, sep, t.SelfTitle.Default)
}

func (t *TitleRow) AsMultiRowConcatElements(sep string) *ExcelStream {
	rows := make([][]*Cell, 0, len(t.Elements))
	for _, e := range t.Elements {
		rows = append(rows, e.Row)
	}
	return NewExcelStreamByRows(rows, t.SelfTitle.FromIndex, t.SelfTitle.ToIndex, sep, t.SelfTitle.Default)
}
