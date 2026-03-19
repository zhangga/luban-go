package rawdefs

type RawBean struct {
	Namespace   string
	Name        string
	Parent      string
	IsValueType bool
	Comment     string
	Tags        map[string]string
	Alias       string
	Sep         string
	Groups      []string
	Fields      []*RawField
	TypeMappers []*TypeMapper
}

func (b *RawBean) FullName() string {
	if len(b.Namespace) > 0 {
		return b.Namespace + "." + b.Name
	}
	return b.Name
}
