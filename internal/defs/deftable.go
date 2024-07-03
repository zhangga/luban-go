package defs

import (
	"github.com/zhangga/luban/core/pctx"
	"github.com/zhangga/luban/core/refs"
	"github.com/zhangga/luban/internal/rawdefs"
)

var _ refs.IDefType = (*DefTable)(nil)

// DefTable 表定义
type DefTable struct {
	DefTypeBase
	Index              string
	ValueType          string
	Mode               rawdefs.TableMode
	InputFiles         []string
	ReadSchemaFromFile bool
	outputFile         string
	IsExported         bool
}

func NewDefTable(rawTable *rawdefs.RawTable) *DefTable {
	return &DefTable{
		DefTypeBase: DefTypeBase{
			Name:      rawTable.Name,
			namespace: rawTable.Namespace,
			Groups:    rawTable.Groups,
			Comment:   rawTable.Comment,
			Tags:      rawTable.Tags,
		},
		Index:              rawTable.Index,
		ValueType:          rawTable.ValueType,
		Mode:               rawTable.Mode,
		InputFiles:         rawTable.InputFiles,
		ReadSchemaFromFile: rawTable.ReadSchemaFromFile,
		outputFile:         rawTable.OutputFile,
	}
}

func (t *DefTable) PreCompile(ctx pctx.Context) {
	//TODO implement me
	panic("implement me")
}

func (t *DefTable) Compile(ctx pctx.Context) {
	//TODO implement me
	panic("implement me")
}

func (t *DefTable) PostCompile(ctx pctx.Context) {
	//TODO implement me
	panic("implement me")
}
