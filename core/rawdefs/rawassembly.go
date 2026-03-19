package rawdefs

type RawGroup struct {
	IsDefault bool
	Names     []string
}

type RawRefGroup struct {
	Name string
	Refs []string
}

type RawTarget struct {
	Name      string
	Manager   string
	TopModule string
	Groups    []string
}

type RawAssembly struct {
	Beans        []*RawBean
	Enums        []*RawEnum
	Tables       []*RawTable
	Groups       []*RawGroup
	Targets      []*RawTarget
	RefGroups    []*RawRefGroup
	ConstAliases map[string]string
}

func NewRawAssembly() *RawAssembly {
	return &RawAssembly{
		Beans:        make([]*RawBean, 0),
		Enums:        make([]*RawEnum, 0),
		Tables:       make([]*RawTable, 0),
		Groups:       make([]*RawGroup, 0),
		Targets:      make([]*RawTarget, 0),
		RefGroups:    make([]*RawRefGroup, 0),
		ConstAliases: make(map[string]string),
	}
}
