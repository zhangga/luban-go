package rawdefs

type RawAssembly struct {
	Beans     []*RawBean
	Enums     []*RawEnum
	Tables    []*RawTable
	Groups    []*RawGroup
	Targets   []*RawTarget
	RefGroups []*RawRefGroup
}
