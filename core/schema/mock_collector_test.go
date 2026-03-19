package schema

import "github.com/zhangga/luban-go/core/rawdefs"

// mockSchemaCollector 用于测试
type mockSchemaCollector struct {
	Assembly *rawdefs.RawAssembly
}

func newMockSchemaCollector() *mockSchemaCollector {
	return &mockSchemaCollector{
		Assembly: rawdefs.NewRawAssembly(),
	}
}

func (m *mockSchemaCollector) Load(config *LubanConfig) error          { return nil }
func (m *mockSchemaCollector) CreateRawAssembly() *rawdefs.RawAssembly { return m.Assembly }

func (m *mockSchemaCollector) AddTable(table *rawdefs.RawTable) {
	m.Assembly.Tables = append(m.Assembly.Tables, table)
}

func (m *mockSchemaCollector) AddBean(bean *rawdefs.RawBean) {
	m.Assembly.Beans = append(m.Assembly.Beans, bean)
}

func (m *mockSchemaCollector) AddEnum(enum *rawdefs.RawEnum) {
	m.Assembly.Enums = append(m.Assembly.Enums, enum)
}

func (m *mockSchemaCollector) AddRefGroup(refGroup *rawdefs.RawRefGroup) {
	m.Assembly.RefGroups = append(m.Assembly.RefGroups, refGroup)
}

func (m *mockSchemaCollector) AddConstAlias(name string, alias string) {
	m.Assembly.ConstAliases[name] = alias
}
