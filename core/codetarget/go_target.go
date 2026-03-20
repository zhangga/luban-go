package codetarget

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"

	"github.com/zhangga/luban-go/core/defs"
)

// GoCodeTarget 实现了针对 Go 语言的基础代码生成
type GoCodeTarget struct {
	name      string
	beanTmpl  *template.Template
	tableTmpl *template.Template
}

func NewGoCodeTarget() *GoCodeTarget {
	beanTemplateStr := `package {{.Namespace}}

// {{.Name}} {{.Comment}}
type {{.Name}} struct {
{{- range .Fields}}
	{{.Name}} {{.Type}} // {{.Comment}}
{{- end}}
}
`
	tableTemplateStr := `package {{.Namespace}}

// {{.Name}} 管理表数据
type {{.Name}} struct {
	DataList []*{{.ValueType}}
	DataMap  map[{{.IndexType}}]*{{.ValueType}}
}
`
	beanTmpl := template.Must(template.New("bean").Parse(beanTemplateStr))
	tableTmpl := template.Must(template.New("table").Parse(tableTemplateStr))

	return &GoCodeTarget{
		name:      "go-bin",
		beanTmpl:  beanTmpl,
		tableTmpl: tableTmpl,
	}
}

func (t *GoCodeTarget) Name() string {
	return t.name
}

func (t *GoCodeTarget) ValidateDefinition(assembly *defs.DefAssemblyImpl) error {
	// 简单的验证，真实环境可做复杂依赖检查
	return nil
}

func (t *GoCodeTarget) Handle(assembly *defs.DefAssemblyImpl, manifest *OutputFileManifest) error {
	targetGroups := assembly.Target.Groups

	// 生成 Bean 代码
	for _, rawType := range assembly.Types {
		if bean, ok := rawType.(*defs.DefBeanImpl); ok {
			out, err := t.GenerateBean(bean, targetGroups)
			if err != nil {
				return err
			}
			filePath := fmt.Sprintf("%s.go", strings.ToLower(bean.Name()))
			manifest.AddCodeFile(NewOutputFile(filePath, out))
		}
	}

	// 生成 Table 代码
	for _, table := range assembly.ExportTables {
		out, err := t.GenerateTable(assembly, table)
		if err != nil {
			return err
		}
		filePath := fmt.Sprintf("%s.go", strings.ToLower(table.Name()))
		manifest.AddCodeFile(NewOutputFile(filePath, out))
	}

	return nil
}

// typeMapping 简单的 Go 类型映射
func typeMapping(lubanType string) string {
	switch lubanType {
	case "int":
		return "int32"
	case "long":
		return "int64"
	case "float":
		return "float32"
	case "double":
		return "float64"
	case "bool":
		return "bool"
	case "string":
		return "string"
	default:
		// 复杂类型（Bean 或 List 等）暂时原样返回或需要更复杂的映射解析
		return lubanType
	}
}

func (t *GoCodeTarget) GenerateBean(bean *defs.DefBeanImpl, targetGroups []string) ([]byte, error) {
	type FieldInfo struct {
		Name    string
		Type    string
		Comment string
	}

	fields := make([]FieldInfo, 0, len(bean.Fields))
	for _, f := range bean.Fields {
		if f.NeedExport(targetGroups) {
			fields = append(fields, FieldInfo{
				Name:    strings.Title(f.Raw.Name), // Go 导出字段大写
				Type:    typeMapping(f.Raw.Type),
				Comment: f.Raw.Comment,
			})
		}
	}

	ns := bean.Namespace()
	if ns == "" {
		ns = "config"
	}

	data := map[string]interface{}{
		"Namespace": strings.ToLower(ns),
		"Name":      bean.Name(),
		"Comment":   bean.Raw.Comment,
		"Fields":    fields,
	}

	var buf bytes.Buffer
	if err := t.beanTmpl.Execute(&buf, data); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (t *GoCodeTarget) GenerateTable(assembly *defs.DefAssemblyImpl, table *defs.DefTable) ([]byte, error) {
	ns := table.Namespace()
	if ns == "" {
		ns = "config"
	}

	data := map[string]interface{}{
		"Namespace": strings.ToLower(ns),
		"Name":      table.Name(),
		"ValueType": table.Raw.ValueType,
		"IndexType": "int32", // 暂时写死，真实应根据 table.Index 找寻具体字段类型
	}

	var buf bytes.Buffer
	if err := t.tableTmpl.Execute(&buf, data); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
