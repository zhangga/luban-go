package excel

type dataReader struct {
	fileName  string
	sheetName string
	rows      [][]string
	curRow    int
	curCol    int
}

func newDataReader(rows [][]string, fileName, sheetName string) *dataReader {
	return &dataReader{
		fileName:  fileName,
		sheetName: sheetName,
		rows:      rows,
		curRow:    0,
		curCol:    0,
	}
}

func (r *dataReader) GetString(row, col int) string {
	if row >= len(r.rows) {
		return ""
	}
	line := r.rows[row]
	if col >= len(line) {
		return ""
	}
	return r.rows[row][col]
}

func (r *dataReader) ReadRow() []string {
	if r.curRow >= len(r.rows) {
		return nil
	}
	row := r.rows[r.curRow]
	r.curRow++
	return row
}

func (r *dataReader) PeekRow() []string {
	if r.curRow >= len(r.rows) {
		return nil
	}
	return r.rows[r.curRow]
}

func (r *dataReader) FieldCount() int {
	if r.curRow >= len(r.rows) {
		return 0
	}
	return len(r.rows[r.curRow])
}

func (r *dataReader) NextRow() {
	r.curRow++
}
