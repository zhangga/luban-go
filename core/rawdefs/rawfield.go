package rawdefs

type RawField struct {
	Name              string
	Alias             string
	Type              string
	Comment           string
	Tags              map[string]string
	Variants          []string
	NotNameValidation bool
	Groups            []string
}
