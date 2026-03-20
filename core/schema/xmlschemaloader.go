package schema

import (
	"encoding/xml"
	"fmt"
	"github.com/zhangga/luban-go/core/rawdefs"
	"io"
	"os"
	"strings"
)

type SchemaLoaderBase struct {
	LoaderType string
	Col        ISchemaCollector
}

func (b *SchemaLoaderBase) Type() string {
	return b.LoaderType
}

func (b *SchemaLoaderBase) Collector() ISchemaCollector {
	return b.Col
}

type XmlSchemaLoader struct {
	*SchemaLoaderBase
	fileName       string
	namespaceStack []string
}

func NewXmlSchemaLoader(collector ISchemaCollector) *XmlSchemaLoader {
	return &XmlSchemaLoader{
		SchemaLoaderBase: &SchemaLoaderBase{
			LoaderType: "xml",
			Col:        collector,
		},
		namespaceStack: make([]string, 0),
	}
}

func (l *XmlSchemaLoader) curNamespace() string {
	if len(l.namespaceStack) > 0 {
		return l.namespaceStack[len(l.namespaceStack)-1]
	}
	return ""
}

func (l *XmlSchemaLoader) Load(fileName string) error {
	l.fileName = fileName

	f, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer f.Close()

	// 采用简单的 Decoder 逐行解析 XML (或反序列化为 struct)
	// 由于 Luban 的 xml 定义存在大量动态嵌套，使用 decoder() 更灵活
	decoder := xml.NewDecoder(f)

	for {
		t, err := decoder.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		switch se := t.(type) {
		case xml.StartElement:
			if err := l.handleElement(decoder, &se); err != nil {
				return err
			}
		}
	}
	return nil
}

func (l *XmlSchemaLoader) handleElement(decoder *xml.Decoder, se *xml.StartElement) error {
	tagName := se.Name.Local
	switch tagName {
	case "module":
		return l.addModule(decoder, se)
	case "enum":
		return l.addEnum(decoder, se)
	case "bean":
		return l.addBean(decoder, se)
	case "table":
		return l.addTable(decoder, se)
	case "refgroup":
		// TODO
	case "constalias":
		// TODO
	}
	return nil
}

// 获取XML属性的辅助函数
func getAttr(se *xml.StartElement, name string) string {
	for _, a := range se.Attr {
		if a.Name.Local == name {
			return a.Value
		}
	}
	return ""
}

func (l *XmlSchemaLoader) addModule(decoder *xml.Decoder, se *xml.StartElement) error {
	name := getAttr(se, "name")
	name = strings.TrimSpace(name)

	newNs := name
	if len(l.namespaceStack) > 0 {
		parentNs := l.namespaceStack[len(l.namespaceStack)-1]
		if parentNs != "" && name != "" {
			newNs = parentNs + "." + name
		} else if parentNs != "" {
			newNs = parentNs
		}
	}
	l.namespaceStack = append(l.namespaceStack, newNs)

	for {
		t, err := decoder.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		switch child := t.(type) {
		case xml.StartElement:
			if err := l.handleElement(decoder, &child); err != nil {
				return err
			}
		case xml.EndElement:
			if child.Name.Local == se.Name.Local {
				l.namespaceStack = l.namespaceStack[:len(l.namespaceStack)-1]
				return nil
			}
		}
	}

	return nil
}

func (l *XmlSchemaLoader) addEnum(decoder *xml.Decoder, se *xml.StartElement) error {
	name := getAttr(se, "name")
	if name == "" {
		return fmt.Errorf("enum requires name attribute in %s", l.fileName)
	}

	enum := &rawdefs.RawEnum{
		Namespace: l.curNamespace(),
		Name:      name,
		Comment:   getAttr(se, "comment"),
		Items:     make([]*rawdefs.EnumItem, 0),
	}

	// 读取 enum_item 子节点或直接闭合
	for {
		t, err := decoder.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		switch child := t.(type) {
		case xml.StartElement:
			if child.Name.Local == "var" || child.Name.Local == "item" { // 注意有些地方用 var 作为枚举项
				itemName := getAttr(&child, "name")
				itemValue := getAttr(&child, "value")
				itemAlias := getAttr(&child, "alias")
				enum.Items = append(enum.Items, &rawdefs.EnumItem{
					Name:  itemName,
					Value: itemValue,
					Alias: itemAlias,
				})
			}
		case xml.EndElement:
			if child.Name.Local == se.Name.Local {
				l.Col.AddEnum(enum)
				return nil
			}
		}
	}

	return nil
}

func (l *XmlSchemaLoader) addBean(decoder *xml.Decoder, se *xml.StartElement) error {
	name := getAttr(se, "name")
	if name == "" {
		return fmt.Errorf("bean requires name attribute in %s", l.fileName)
	}

	bean := &rawdefs.RawBean{
		Namespace: l.curNamespace(),
		Name:      name,
		Parent:    getAttr(se, "parent"),
		Comment:   getAttr(se, "comment"),
		Fields:    make([]*rawdefs.RawField, 0),
	}

	// 读取 var (field)
	for {
		t, err := decoder.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		switch child := t.(type) {
		case xml.StartElement:
			if child.Name.Local == "var" {
				fieldName := getAttr(&child, "name")
				fieldType := getAttr(&child, "type")
				fieldAlias := getAttr(&child, "alias")
				bean.Fields = append(bean.Fields, &rawdefs.RawField{
					Name:  fieldName,
					Type:  fieldType,
					Alias: fieldAlias,
				})
			}
		case xml.EndElement:
			if child.Name.Local == se.Name.Local {
				l.Col.AddBean(bean)
				return nil
			}
		}
	}
	return nil
}

func (l *XmlSchemaLoader) addTable(decoder *xml.Decoder, se *xml.StartElement) error {
	name := getAttr(se, "name")
	value := getAttr(se, "value")
	index := getAttr(se, "index")

	if name == "" || value == "" {
		return fmt.Errorf("table requires name and value attributes in %s", l.fileName)
	}

	modeAttr := getAttr(se, "mode")
	var mode rawdefs.TableMode
	switch modeAttr {
	case "one":
		mode = rawdefs.TableModeOne
	case "map":
		mode = rawdefs.TableModeMap
	case "list":
		mode = rawdefs.TableModeList
	default:
		mode = rawdefs.TableModeMap // 默认为 map
	}

	table := &rawdefs.RawTable{
		Namespace: l.curNamespace(),
		Name:      name,
		ValueType: value,
		Index:     index,
		Mode:      mode,
	}

	fileAttr := getAttr(se, "file")
	if fileAttr == "" {
		fileAttr = getAttr(se, "input")
	}
	if fileAttr != "" {
		table.InputFiles = strings.Split(fileAttr, ",")
	}

	for {
		t, err := decoder.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		switch child := t.(type) {
		case xml.StartElement:
			// table 内部结构目前可忽略或添加扩展
		case xml.EndElement:
			if child.Name.Local == se.Name.Local {
				l.Col.AddTable(table)
				return nil
			}
		}
	}
	return nil
}
