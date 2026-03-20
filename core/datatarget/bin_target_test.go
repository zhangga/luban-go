package datatarget

import (
	"encoding/binary"
	"math"
	"testing"

	"github.com/zhangga/luban-go/core/datas"
	"github.com/zhangga/luban-go/core/defs"
	"github.com/zhangga/luban-go/core/rawdefs"
	"github.com/zhangga/luban-go/core/types"
)

func TestBinDataVisitor_Primitives(t *testing.T) {
	visitor := NewBinDataVisitor(nil)
	
	// Test Int
	visitor.VisitDInt(datas.NewDInt(100))
	buf := visitor.Buf().Bytes()
	if len(buf) != 4 {
		t.Fatalf("Expected 4 bytes, got %d", len(buf))
	}
	if binary.LittleEndian.Uint32(buf) != 100 {
		t.Errorf("Expected 100, got %d", binary.LittleEndian.Uint32(buf))
	}
	
	// Test Float
	visitor = NewBinDataVisitor(nil)
	visitor.VisitDFloat(datas.NewDFloat(3.14))
	buf = visitor.Buf().Bytes()
	fBits := binary.LittleEndian.Uint32(buf)
	fVal := math.Float32frombits(fBits)
	if math.Abs(float64(fVal-3.14)) > 0.001 {
		t.Errorf("Expected 3.14, got %f", fVal)
	}

	// Test String
	visitor = NewBinDataVisitor(nil)
	visitor.VisitDString(datas.NewDString("hello"))
	buf = visitor.Buf().Bytes()
	strLen := binary.LittleEndian.Uint32(buf[:4])
	if strLen != 5 {
		t.Errorf("Expected string length 5, got %d", strLen)
	}
	if string(buf[4:]) != "hello" {
		t.Errorf("Expected 'hello', got %s", string(buf[4:]))
	}
}

func TestBinDataTarget_ExportTable(t *testing.T) {
	rawAss := rawdefs.NewRawAssembly()
	rawAss.Beans = append(rawAss.Beans, &rawdefs.RawBean{
		Name: "TestItem",
		Fields: []*rawdefs.RawField{
			{Name: "Id", Type: "int"},
			{Name: "Name", Type: "string"},
		},
	})
	
	rawAss.Tables = append(rawAss.Tables, &rawdefs.RawTable{
		Name: "TbItem",
		ValueType: "TestItem",
	})
	
	assembly, _ := defs.NewDefAssemblyImpl(rawAss, "all", nil)
	
	beanImpl := assembly.GetType("TestItem").(*defs.DefBeanImpl)
	// mock fields
	for _, f := range rawAss.Beans[0].Fields {
		df := defs.NewDefField(f)
		beanImpl.Fields = append(beanImpl.Fields, df)
		beanImpl.HierarchyFields = append(beanImpl.HierarchyFields, df)
	}
	
	beanType := types.NewTBean(false, beanImpl, nil)
	
	dBean1 := datas.NewDBean(beanType, beanType, []datas.DType{
		datas.NewDInt(1),
		datas.NewDString("Sword"),
	})
	
	records := []*defs.Record{
		defs.NewRecord(dBean1, "test.xlsx", nil),
	}
	
	table := assembly.ExportTables[0]
	target := NewBinDataTarget(nil)
	
	out, err := target.ExportTable(table, records)
	if err != nil {
		t.Fatalf("Failed to export table: %v", err)
	}
	
	if out.File != "tbitem.bytes" {
		t.Errorf("Expected tbitem.bytes, got %s", out.File)
	}
	
	buf := out.Content
	// count: 4 bytes
	// item1.Id: 4 bytes
	// item1.Name len: 4 bytes
	// item1.Name string: 5 bytes
	expectedLen := 4 + 4 + 4 + 5
	if len(buf) != expectedLen {
		t.Errorf("Expected %d bytes, got %d", expectedLen, len(buf))
	}
}