package rawref

type RawTarget struct {
	Name      string   `json:"name"`
	Manager   string   `json:"manager"`
	TopModule string   `json:"top_module"`
	Groups    []string `json:"groups"`
}
