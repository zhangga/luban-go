package options

import (
	"fmt"
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
	CustomOpts          map[string]string // 自定义参数
}

func CreateArguments(opts CommandOptions) Arguments {
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
		CustomOpts:          options,
	}
}

func (args Arguments) GetOptionOrDefault(namespace, name, defaultValue string, useGlobalIfNotExits bool) string {
	if val, ok := args.TryGetOption(namespace, name, useGlobalIfNotExits); ok {
		return val
	}
	return defaultValue
}

func (args Arguments) TryGetOption(namespace, name string, useGlobalIfNotExits bool) (string, bool) {
	var fullName string
	for {
		if len(namespace) == 0 {
			fullName = name
		} else {
			fullName = namespace + "." + name
		}

		if val, ok := args.CustomOpts[fullName]; ok {
			return val, true
		}

		if len(namespace) == 0 || !useGlobalIfNotExits {
			return "", false
		}

		if idx := strings.LastIndex(namespace, "."); idx == -1 {
			namespace = ""
		} else {
			namespace = namespace[:idx]
		}
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
