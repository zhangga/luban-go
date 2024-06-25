package lubanconf

type Conf struct {
	Groups      []Group      `json:"groups"`
	SchemaFiles []SchemaFile `json:"schemaFiles"`
	DataDir     string       `json:"dataDir"`
	Targets     []Target     `json:"targets"`
}

type Group struct {
	Names   []string `json:"names"`
	Default bool     `json:"default"`
}

type SchemaFile struct {
	FileName string `json:"fileName"`
	Type     string `json:"type"`
}

type Target struct {
	Name      string   `json:"name"`
	Manager   string   `json:"manager"`
	Groups    []string `json:"groups"`
	TopModule string   `json:"topModule"`
}
