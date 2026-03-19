package rawdefs

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
	IsFlags        bool
	IsUniqueItemId bool
	Comment        string
	Tags           map[string]string
	Items          []*EnumItem
	Groups         []string
	TypeMappers    []*TypeMapper
}

func (e *RawEnum) FullName() string {
	if len(e.Namespace) > 0 {
		return e.Namespace + "." + e.Name
	}
	return e.Name
}
