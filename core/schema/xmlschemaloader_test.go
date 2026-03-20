package schema

import (
	"testing"
)

func TestXmlSchemaLoader(t *testing.T) {
	collector := newMockSchemaCollector()
	loader := NewXmlSchemaLoader(collector)

	err := loader.Load("test_schema.xml")
	if err != nil {
		t.Fatalf("Failed to load xml schema: %v", err)
	}

	assembly := collector.CreateRawAssembly()

	if len(assembly.Enums) != 1 {
		t.Errorf("Expected 1 enum, got %d", len(assembly.Enums))
	} else {
		enum := assembly.Enums[0]
		if enum.FullName() != "demo.Gender" {
			t.Errorf("Expected demo.Gender, got %s", enum.FullName())
		}
		if len(enum.Items) != 2 {
			t.Errorf("Expected 2 enum items, got %d", len(enum.Items))
		}
	}

	if len(assembly.Beans) != 4 {
		t.Errorf("Expected 4 beans, got %d", len(assembly.Beans))
	} else {
		bean := assembly.Beans[0]
		if bean.FullName() != "demo.Item" {
			t.Errorf("Expected demo.Item, got %s", bean.FullName())
		}
		if len(bean.Fields) != 3 {
			t.Errorf("Expected 3 fields, got %d", len(bean.Fields))
		}
		if bean.Fields[0].Name != "Id" || bean.Fields[0].Type != "int" {
			t.Errorf("Expected Id:int, got %s:%s", bean.Fields[0].Name, bean.Fields[0].Type)
		}
	}

	if len(assembly.Tables) != 1 {
		t.Errorf("Expected 1 table, got %d", len(assembly.Tables))
	} else {
		table := assembly.Tables[0]
		if table.FullName() != "demo.sub.TbItem" {
			t.Errorf("Expected demo.sub.TbItem, got %s", table.FullName())
		}
		if table.ValueType != "Item" {
			t.Errorf("Expected value Item, got %s", table.ValueType)
		}
	}
}
