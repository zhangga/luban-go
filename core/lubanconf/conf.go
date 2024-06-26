package lubanconf

// Conf luban.conf中的配置数据
type Conf struct {
	DataDir     string       `json:"dataDir"`
	Groups      []Group      `json:"groups"`
	SchemaFiles []SchemaFile `json:"schemaFiles"`
	Targets     []Target     `json:"targets"`
}

// LubanConfig Load解析后的数据
type LubanConfig struct {
	ConfigFileName string
	InputDataDir   string
	Groups         []Group
	Targets        []Target
	Imports        []SchemaFile
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
