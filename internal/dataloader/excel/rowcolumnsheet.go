package excel

import (
	"fmt"
	"github.com/zhangga/luban/internal/utils"
)

type RowColumnSheet struct {
	Name             string
	UrlWithoutParams string
	SheetName        string
	Rows             []*Row
}

type Row struct {
	Tag string
	Row *TitleRow
}

func NewRowColumnSheet(rawUrl, name, sheetName string) *RowColumnSheet {
	return &RowColumnSheet{
		UrlWithoutParams: rawUrl,
		Name:             name,
		SheetName:        sheetName,
	}
}

func (s *RowColumnSheet) Load(rawSheet *RawSheet) {
	cells := rawSheet.Cells
	title := rawSheet.Title
	if !title.HierarchyMultiRows {
		for _, row := range cells {
			if isBlankRow(row, title) {
				continue
			}
			s.Rows = append(s.Rows, &Row{
				Tag: getRowTag(row),
				Row: parseOneLineTitleRow(title, row),
			})
		}
	} else {
		for _, table := range splitRows(title, cells) {
			s.Rows = append(s.Rows, &Row{
				Tag: getRowTag(table[0]),
				Row: parseMultiLineTitleRow(title, table),
			})
		}
	}
}

func (s *RowColumnSheet) UrlWithParams() string {
	if len(s.SheetName) == 0 {
		return s.UrlWithoutParams
	}
	return fmt.Sprintf("%s@%s", s.SheetName, s.UrlWithoutParams)
}

func getRowTag(row []*Cell) string {
	if len(row) == 0 {
		return ""
	}
	return row[0].ValueOrEmpty()
}

type oneTable [][]*Cell

func splitRows(title *Title, rows [][]*Cell) []oneTable {
	var allTable []oneTable

	var table oneTable
	for _, row := range rows {
		if isBlankRowByIndex(row, title.FromIndex, title.ToIndex) {
			continue
		}

		if table == nil {
			table = append(table, row)
			continue
		}

		if isMultiRowsExtendRow(row, title) {
			table = append(table, row)
		} else {
			allTable = append(allTable, table)
			table = nil
			table = append(table, row)
		}
	}
	if len(table) > 0 {
		allTable = append(allTable, table)
	}

	return allTable
}

func isMultiRowsExtendRow(row []*Cell, title *Title) bool {
	if title.HasSubTitle() {
		for _, sub := range title.SubTitleList {
			if !sub.SelfMultiRows && !isMultiRowsExtendRow(row, sub) {
				return false
			}
		}
		return true
	}

	return isBlankRowByIndex(row, title.FromIndex, title.ToIndex)
}

func parseMultiLineTitleRow(title *Title, table oneTable) *TitleRow {
	if !title.HasSubTitle() {
		if title.SelfMultiRows {
			return &TitleRow{SelfTitle: title, Rows: table}
		} else {
			return &TitleRow{SelfTitle: title, Row: table[0]}
		}
	}

	fields := make(map[string]*TitleRow)
	for _, sub := range title.SubTitleList {
		if sub.SelfMultiRows {
			var eles []*TitleRow
			if sub.SubHierarchyMultiRows {
				for _, eleRows := range splitRows(sub, table) {
					eles = append(eles, parseMultiLineTitleRow(sub, eleRows))
				}
			} else {
				for _, eleRow := range table {
					if isBlankRowByIndex(eleRow, sub.FromIndex, sub.ToIndex) {
						continue
					}
					eles = append(eles, parseOneLineTitleRow(sub, eleRow))
				}
			}
			fields[sub.Name] = &TitleRow{SelfTitle: sub, Elements: eles}
		} else {
			if sub.SubHierarchyMultiRows {
				fields[sub.Name] = parseMultiLineTitleRow(sub, table)
			} else {
				fields[sub.Name] = parseOneLineTitleRow(sub, table[0])
			}
		}
	}
	return &TitleRow{
		SelfTitle: title,
		Fields:    fields,
	}
}

func parseOneLineTitleRow(title *Title, row []*Cell) *TitleRow {
	if !title.HasSubTitle() {
		return &TitleRow{SelfTitle: title, Row: row}
	}

	fields := make(map[string]*TitleRow)
	for _, sub := range title.SubTitleList {
		fields[sub.Name] = parseOneLineTitleRow(sub, row)
	}
	return &TitleRow{
		SelfTitle: title,
		Fields:    fields,
	}
}

func isBlankRow(row []*Cell, title *Title) bool {
	if len(title.SubTitleList) == 0 {
		return isBlankRowByIndex(row, title.FromIndex, title.ToIndex)
	}
	return utils.All(title.SubTitleList, func(t *Title) bool {
		return isBlankRow(row, t)
	})
}

func isBlankRowByIndex(row []*Cell, fromIndex, toIndex int) bool {
	i := 1
	if fromIndex > i {
		i = fromIndex
	}
	n := toIndex
	if len(row)-1 < n {
		n = len(row) - 1
	}

	for ; i <= n; i++ {
		if len(row[i].ValueOrEmpty()) > 0 {
			return false
		}
	}
	return true
}
