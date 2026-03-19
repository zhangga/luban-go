package schema

import "github.com/zhangga/luban-go/core/rawdefs"

type SchemaFileInfo struct {
	FileName string
	Type     string
}

type LubanConfig struct {
	ConfigFileName string
	Groups         []*rawdefs.RawGroup
	Targets        []*rawdefs.RawTarget
	Imports        []*SchemaFileInfo
	Xargs          []string
	InputDataDir   string
}
