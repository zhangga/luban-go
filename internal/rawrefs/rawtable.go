package rawrefs

import (
	"fmt"
	"github.com/zhangga/luban/core/refs"
	"strings"
)

var _ refs.UnimplementedTable = (*RawTable)(nil)

type RawTable struct {
	Namespace          string
	Name               string
	Index              string
	ValueType          string
	ReadSchemaFromFile bool
	Mode               TableMode
	Comment            string
	Tags               map[string]string
	Groups             []string
	InputFiles         []string
	OutputFile         string
}

func (r RawTable) MustEmbedUnimplementedTable() {}

type TableMode int

const (
	TableModeOne TableMode = iota
	TableModeMap
	TableModeList
)

func ConvertTableMode(schemaFile, tableName, modeStr, indexStr string) TableMode {
	indexes := strings.FieldsFunc(indexStr, func(c rune) bool {
		return c == ',' || c == '+'
	})

	switch modeStr {
	case "one", "single", "singleton":
		if len(indexStr) > 0 {
			panic(fmt.Errorf("xml定义文件: %s, table: %s, mode: %s, index: %s, 不能同时指定index", schemaFile, tableName, modeStr, indexStr))
		}
		return TableModeOne
	case "map":
		if len(indexStr) > 0 && len(indexes) > 1 {
			panic(fmt.Errorf("xml定义文件: %s, table: %s, mode: %s, index: %s, 是单主键表, 不能包含多个key", schemaFile, tableName, modeStr, indexStr))
		}
		return TableModeMap
	case "list":
		return TableModeList
	case "":
		if len(indexStr) == 0 || len(indexes) == 1 {
			return TableModeMap
		} else {
			return TableModeList
		}
	default:
		panic(fmt.Errorf("xml定义文件: %s, table: %s, mode: %s, 不支持的模式", schemaFile, tableName, modeStr))
	}
}
