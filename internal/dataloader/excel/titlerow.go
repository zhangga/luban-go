package excel

type TitleRow struct {
	Tags      []string
	SelfTitle *Title
	Row       []*Cell
	Rows      [][]*Cell
	Elements  []*TitleRow
	Fields    map[string]*TitleRow
}
