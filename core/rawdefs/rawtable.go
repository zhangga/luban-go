package rawdefs

type TableMode string

const (
	TableModeOne  TableMode = "one"
	TableModeMap  TableMode = "map"
	TableModeList TableMode = "list"
)

type RawTable struct {
	Namespace          string
	Name               string
	Index              string
	ValueType          string
	ReadSchemaFromFile bool
	Mode               TableMode
	Comment            string
	Tags               map[string]string
	Groups             []string
	InputFiles         []string
	OutputFile         string
}

func (t *RawTable) FullName() string {
	if len(t.Namespace) > 0 {
		return t.Namespace + "." + t.Name
	}
	return t.Name
}
