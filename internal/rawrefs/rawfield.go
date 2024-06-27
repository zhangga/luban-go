package rawrefs

type RawField struct {
	Name              string
	Type              string
	Comment           string
	Tags              map[string]string
	NotNameValidation bool
	Groups            []string
}
