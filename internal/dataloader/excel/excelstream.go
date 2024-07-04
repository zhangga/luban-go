package excel

import "strings"

const endOfList = "}"

type ExcelStream struct {
	overrideDefault string
	row             []*Cell
	curIndex        int
	toIndex         int
	LastReadIndex   int
}

func NewExcelStream(row []*Cell, fromIndex, toIndex int, sep, overrideDefault string) *ExcelStream {
	stream := &ExcelStream{
		overrideDefault: overrideDefault,
	}

	if len(sep) == 0 || sep == " " {
		if len(overrideDefault) == 0 {
			stream.row = row
			stream.curIndex = fromIndex
			stream.toIndex = toIndex
		} else {
			for i := fromIndex; i <= toIndex; i++ {
				cell := row[i]
				value := cell.ValueOrEmpty()
				if len(value) == 0 {
					stream.row = append(stream.row, NewCell(cell.Row, cell.Column, overrideDefault, true))
				} else {
					stream.row = append(stream.row, cell)
				}
			}
			stream.curIndex = 0
			stream.toIndex = len(stream.row) - 1
		}
	} else {
		for i := fromIndex; i <= toIndex; i++ {
			cell := row[i]
			value := cell.ValueOrEmpty()
			if len(value) > 0 {
				for _, v := range strings.Split(value, sep) {
					stream.row = append(stream.row, NewCell(cell.Row, cell.Column, v, true))
				}
			} else if len(overrideDefault) > 0 {
				stream.row = append(stream.row, NewCell(cell.Row, cell.Column, overrideDefault, true))
			}
		}
		stream.curIndex = 0
		stream.toIndex = len(stream.row) - 1
	}
	return stream
}

func NewExcelStreamByRows(rows [][]*Cell, fromIndex, toIndex int, sep, overrideDefault string) *ExcelStream {
	stream := &ExcelStream{
		overrideDefault: overrideDefault,
	}

	if len(sep) == 0 || sep == " " {
		if len(overrideDefault) == 0 {
			for _, row := range rows {
				for i := fromIndex; i <= toIndex; i++ {
					stream.row = append(stream.row, row[i])
				}
			}
		} else {
			panic("concated multi rows don't support 'default'")
		}
	} else {
		for _, row := range rows {
			for i := fromIndex; i <= toIndex; i++ {
				cell := row[i]
				value := cell.ValueOrEmpty()
				if len(value) > 0 {
					for _, v := range strings.Split(value, sep) {
						stream.row = append(stream.row, NewCell(cell.Row, cell.Column, v, true))
					}
				} else if len(overrideDefault) > 0 {
					stream.row = append(stream.row, NewCell(cell.Row, cell.Column, overrideDefault, true))
				}
			}
		}
	}
	stream.curIndex = 0
	stream.toIndex = len(stream.row) - 1
	return stream
}

func (s *ExcelStream) TryReadEOF() bool {
	oldIndex := s.curIndex
	for s.curIndex <= s.toIndex {
		value := s.row[s.curIndex].ValueOrEmpty()
		s.curIndex++
		if len(value) > 0 {
			if value == endOfList {
				s.LastReadIndex = s.curIndex - 1
				return true
			} else {
				s.curIndex = oldIndex
				return false
			}
		}
	}
	s.curIndex = oldIndex
	return true
}
