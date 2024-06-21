package pipeline

import "github.com/zhangga/luban/core/options"

// Arguments is the arguments of the pipeline
type Arguments struct {
	Target              string
	ForceLoadTableDatas bool
	IncludeTags         []string
	ExcludeTags         []string
	CodeTargets         []string
	DataTargets         []string
	SchemaCollector     string
	ConfFile            string
	OutputTables        []string
	TimeZone            string
}

func CreateArguments(opts options.CommandOptions) Arguments {
	return Arguments{
		Target:              opts.Target,
		ForceLoadTableDatas: opts.ForceLoadTableDatas,
		IncludeTags:         opts.IncludeTags,
		ExcludeTags:         opts.ExcludeTags,
		CodeTargets:         opts.CodeTargets,
		DataTargets:         opts.DataTargets,
		SchemaCollector:     opts.SchemaCollector,
		ConfFile:            opts.Conf,
		OutputTables:        opts.OutputTables,
		TimeZone:            opts.TimeZone,
	}
}
