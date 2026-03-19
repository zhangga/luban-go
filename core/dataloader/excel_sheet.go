package dataloader

import (
	"fmt"
)

type Cell struct {
	Row    int
	Column int
	Value  interface{}
}

func NewCell(row, col int, value interface{}) *Cell {
	return &Cell{
		Row:    row,
		Column: col,
		Value:  value,
	}
}

func (c *Cell) String() string {
	return fmt.Sprintf("[%s:%d] %v", toAlphaString(c.Column), c.Row+1, c.Value)
}

func toAlphaString(column int) string {
	h := column / 26
	n := column % 26
	if h > 0 {
		return string(rune('A'+h-1)) + string(rune('A'+n))
	}
	return string(rune('A' + n))
}

type RawSheet struct {
	SheetName string
	Cells     [][]*Cell
}

func NewRawSheet(name string) *RawSheet {
	return &RawSheet{
		SheetName: name,
		Cells:     make([][]*Cell, 0),
	}
}
