package codetarget

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/zhangga/luban-go/core/defs"
)

// TemplateCodeTarget 提供了一个通用的基于 text/template 的代码生成器实现
type TemplateCodeTarget struct {
	name          string
	fileExt       string
	templateDir   string
	typeMapper    func(string) string
	beanTmplName  string
	tableTmplName string
	tmpl          *template.Template
}

// NewTemplateCodeTarget 创建一个模板生成器
func NewTemplateCodeTarget(name string, fileExt string, templateDir string, typeMapper func(string) string) (*TemplateCodeTarget, error) {
	t := &TemplateCodeTarget{
		name:          name,
		fileExt:       fileExt,
		templateDir:   templateDir,
		typeMapper:    typeMapper,
		beanTmplName:  "bean.tpl",
		tableTmplName: "table.tpl",
		tmpl:          template.New("root"),
	}

	// 注册一些通用的模板函数
	t.tmpl.Funcs(template.FuncMap{
		"title": strings.Title,
		"lower": strings.ToLower,
		"upper": strings.ToUpper,
		"type":  t.typeMapper,
		"add": func(a, b int) int { return a + b },
	})

	if err := t.loadTemplates(); err != nil {
		return nil, err
	}

	return t, nil
}

func (t *TemplateCodeTarget) loadTemplates() error {
	if t.templateDir == "" {
		return fmt.Errorf("template directory is empty")
	}

	// 读取 bean.tpl
	beanTmplPath := filepath.Join(t.templateDir, t.beanTmplName)
	b, err := ioutil.ReadFile(beanTmplPath)
	if err != nil {
		return fmt.Errorf("failed to read %s: %w", beanTmplPath, err)
	}
	if _, err := t.tmpl.New(t.beanTmplName).Parse(string(b)); err != nil {
		return err
	}

	// 读取 table.tpl
	tableTmplPath := filepath.Join(t.templateDir, t.tableTmplName)
	b, err = ioutil.ReadFile(tableTmplPath)
	if err != nil {
		return fmt.Errorf("failed to read %s: %w", tableTmplPath, err)
	}
	if _, err := t.tmpl.New(t.tableTmplName).Parse(string(b)); err != nil {
		return err
	}

	return nil
}

func (t *TemplateCodeTarget) Name() string {
	return t.name
}

func (t *TemplateCodeTarget) ValidateDefinition(assembly *defs.DefAssemblyImpl) error {
	return nil
}

func (t *TemplateCodeTarget) Handle(assembly *defs.DefAssemblyImpl, manifest *OutputFileManifest) error {
	targetGroups := assembly.Target.Groups

	// 遍历并生成所有的 Bean
	for _, rawType := range assembly.Types {
		if bean, ok := rawType.(*defs.DefBeanImpl); ok {
			out, err := t.GenerateBean(bean, targetGroups)
			if err != nil {
				return err
			}
			filePath := fmt.Sprintf("%s.%s", strings.ToLower(bean.Name()), t.fileExt)
			manifest.AddCodeFile(NewOutputFile(filePath, out))
		}
	}

	// 遍历并生成所有的 Table
	for _, table := range assembly.ExportTables {
		out, err := t.GenerateTable(assembly, table)
		if err != nil {
			return err
		}
		filePath := fmt.Sprintf("%s.%s", strings.ToLower(table.Name()), t.fileExt)
		manifest.AddCodeFile(NewOutputFile(filePath, out))
	}

	return nil
}

func (t *TemplateCodeTarget) GenerateBean(bean *defs.DefBeanImpl, targetGroups []string) ([]byte, error) {
	type FieldInfo struct {
		Name    string
		Type    string
		Comment string
	}

	exportFields := make([]FieldInfo, 0)
	for _, f := range bean.Fields {
		if f.NeedExport(targetGroups) {
			exportFields = append(exportFields, FieldInfo{
				Name:    f.Raw.Name,
				Type:    f.Raw.Type,
				Comment: f.Raw.Comment,
			})
		}
	}

	parentFields := make([]FieldInfo, 0)
	if bean.ParentDefType != nil {
		for _, f := range bean.ParentDefType.HierarchyFields {
			if f.NeedExport(targetGroups) {
				parentFields = append(parentFields, FieldInfo{
					Name:    f.Raw.Name,
					Type:    f.Raw.Type,
					Comment: f.Raw.Comment,
				})
			}
		}
	}

	hierarchyFields := make([]FieldInfo, 0)
	for _, f := range bean.HierarchyFields {
		if f.NeedExport(targetGroups) {
			hierarchyFields = append(hierarchyFields, FieldInfo{
				Name:    f.Raw.Name,
				Type:    f.Raw.Type,
				Comment: f.Raw.Comment,
			})
		}
	}

	data := map[string]interface{}{
		"Namespace":       bean.Namespace(),
		"Name":            bean.Name(),
		"Comment":         bean.Raw.Comment,
		"ExportFields":    exportFields,
		"ParentFields":    parentFields,
		"HierarchyFields": hierarchyFields,
		"Parent":          bean.Parent,
		"IsDynamic":       bean.IsAbstractType,
	}

	var buf bytes.Buffer
	if err := t.tmpl.ExecuteTemplate(&buf, t.beanTmplName, data); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (t *TemplateCodeTarget) GenerateTable(assembly *defs.DefAssemblyImpl, table *defs.DefTable) ([]byte, error) {
	indexType := "int" // default fallback

	// 查找 ValueType 对应的 Bean，并获取 Index 字段的类型
	if bean, ok := assembly.Types[table.Raw.ValueType].(*defs.DefBeanImpl); ok {
		for _, f := range bean.HierarchyFields {
			if f.Raw.Name == table.Raw.Index {
				indexType = f.Raw.Type
				break
			}
		}
	}

	data := map[string]interface{}{
		"Namespace": table.Namespace(),
		"Name":      table.Name(),
		"ValueType": table.Raw.ValueType,
		"Index":     table.Raw.Index,
		"IndexType": indexType,
		"Mode":      string(table.Raw.Mode),
	}

	var buf bytes.Buffer
	if err := t.tmpl.ExecuteTemplate(&buf, t.tableTmplName, data); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
