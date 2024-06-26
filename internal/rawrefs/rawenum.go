package rawrefs

type EnumItem struct {
	Name    string
	Alias   string
	Value   string
	Comment string
	Tags    map[string]string
}

type RawEnum struct {
	Namespace      string
	Name           string
	FullName       string
	IsFlags        bool
	IsUniqueItemId bool
	Comment        string
	Tags           map[string]string
	Items          []EnumItem
	Groups         []string
	TypeMappers    []TypeMapper
}
