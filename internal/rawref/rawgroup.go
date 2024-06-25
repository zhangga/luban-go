package rawref

type RawGroup struct {
	IsDefault bool     `json:"is_default"`
	Names     []string `json:"names"`
}
