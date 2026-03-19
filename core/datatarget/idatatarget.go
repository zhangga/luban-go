package datatarget

import (
	"github.com/zhangga/luban-go/core/codetarget"
	"github.com/zhangga/luban-go/core/defs"
)

type AggregationType string

const (
	AggregationTable  AggregationType = "table"
	AggregationTables AggregationType = "tables"
	AggregationRecord AggregationType = "record"
	AggregationOther  AggregationType = "other"
)

// IDataTarget 定义数据生成目标（例如：json，bin，bson 等）
type IDataTarget interface {
	AggregationType() AggregationType
	ExportAllRecords() bool

	ExportTable(table *defs.DefTable, records []*defs.Record) (*codetarget.OutputFile, error)
	ExportTables(tables []*defs.DefTable) (*codetarget.OutputFile, error)
	ExportRecord(table *defs.DefTable, record *defs.Record) (*codetarget.OutputFile, error)
}

// IDataExporter 定义数据导出器
type IDataExporter interface {
	Handle(assembly *defs.DefAssemblyImpl, dataTarget IDataTarget, manifest *codetarget.OutputFileManifest) error
}
