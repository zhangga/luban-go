package options

// CommandOptions 启动参数
type CommandOptions struct {
	SchemaCollector       string   `json:"schema_collector" mapstructure:"schema_collector"`                 // schema collector name
	Conf                  string   `json:"conf" mapstructure:"conf"`                                         // luban conf file path
	Target                string   `json:"target" mapstructure:"target"`                                     // target name
	CodeTargets           []string `json:"code_targets" mapstructure:"code_targets"`                         // code target name list
	DataTargets           []string `json:"data_targets" mapstructure:"data_targets"`                         // data target name list
	Pipeline              string   `json:"pipeline" mapstructure:"pipeline"`                                 // pipeline name
	ForceLoadTableDatas   bool     `json:"force_load_table_datas" mapstructure:"force_load_table_datas"`     // force load table datas when not any dataTarget
	IncludeTags           []string `json:"include_tags" mapstructure:"include_tags"`                         // include tags
	ExcludeTags           []string `json:"exclude_tags" mapstructure:"exclude_tags"`                         // exclude tags
	OutputTables          []string `json:"output_tables" mapstructure:"output_tables"`                       // output tables
	TimeZone              string   `json:"time_zone" mapstructure:"time_zone"`                               // time zone
	CustomTemplateDirs    []string `json:"custom_template_dirs" mapstructure:"custom_template_dirs"`         // custom template dirs
	ValidationFailAsError bool     `json:"validation_fail_as_error" mapstructure:"validation_fail_as_error"` // validation fail as error
	Xargs                 []string `json:"xargs" mapstructure:"xargs"`                                       // args like -x a=1 -x b=2
	Verbose               bool     `json:"verbose" mapstructure:"verbose"`                                   // verbose
	LogConfig             string   `json:"log_config" mapstructure:"log_config"`                             // log config file
}
