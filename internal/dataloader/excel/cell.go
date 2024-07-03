package excel

import "strings"

type Cell struct {
	Row    int // 从0开始
	Column int // 从0开始
	NullableString
}

func NewCell(row, column int, value string, valid bool) *Cell {
	return &Cell{
		Row:            row,
		Column:         column,
		NullableString: NullableString{valid: valid, value: value},
	}
}

type NullableString struct {
	valid bool
	value string
}

func (c *Cell) Value() (string, bool) {
	return strings.TrimSpace(c.value), c.valid
}

func (c *Cell) ValueOrDefault(defaultValue string) string {
	if c.valid {
		return strings.TrimSpace(c.value)
	}
	return defaultValue
}

func (c *Cell) ValueOrEmpty() string {
	return c.ValueOrDefault("")
}
