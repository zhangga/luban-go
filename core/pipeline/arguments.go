package pipeline

import (
	"fmt"
	"github.com/zhangga/luban/core/options"
	"strings"
)

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
	CustomArgs          map[string]string
}

func CreateArguments(opts options.CommandOptions) Arguments {
	options := parseOptions(opts.Xargs...)
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
		CustomArgs:          options,
	}
}

func parseOptions(xargs ...string) map[string]string {
	options := make(map[string]string)
	for _, arg := range xargs {
		kv := strings.Split(arg, "=")
		if len(kv) != 2 {
			err := fmt.Errorf("invalid xargs: %s", arg)
			panic(err)
		}
		if _, ok := options[kv[0]]; ok {
			err := fmt.Errorf("duplicate xargs: %s", arg)
			panic(err)
		}
		options[kv[0]] = kv[1]
	}
	return options
}
