package datatarget

import (
	"fmt"
	"strings"

	"github.com/zhangga/luban-go/core/codetarget"
	"github.com/zhangga/luban-go/core/defs"
)

type BinDataTarget struct {
	targetGroups []string
}

func NewBinDataTarget(targetGroups []string) *BinDataTarget {
	return &BinDataTarget{
		targetGroups: targetGroups,
	}
}

func (t *BinDataTarget) AggregationType() AggregationType {
	return AggregationTable
}

func (t *BinDataTarget) ExportAllRecords() bool {
	return true
}

func (t *BinDataTarget) ExportTable(table *defs.DefTable, records []*defs.Record) (*codetarget.OutputFile, error) {
	visitor := NewBinDataVisitor(t.targetGroups)
	
	// 在原版中，通常二进制文件头包含一个数量信息或者校验码等，这里我们先简单写入数量
	visitor.buf.WriteInt(int32(len(records)))

	for _, r := range records {
		if r.Data != nil {
			r.Data.Accept(visitor)
		}
	}

	fileName := fmt.Sprintf("%s.bytes", strings.ToLower(table.Name()))
	return codetarget.NewOutputFile(fileName, visitor.buf.Bytes()), nil
}

func (t *BinDataTarget) ExportTables(tables []*defs.DefTable) (*codetarget.OutputFile, error) {
	return nil, fmt.Errorf("BinDataTarget does not support ExportTables")
}

func (t *BinDataTarget) ExportRecord(table *defs.DefTable, record *defs.Record) (*codetarget.OutputFile, error) {
	return nil, fmt.Errorf("BinDataTarget does not support ExportRecord")
}