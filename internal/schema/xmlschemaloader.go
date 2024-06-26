package schema

import (
	"fmt"
	"github.com/beevik/etree"
	"github.com/zhangga/luban/core/schema"
	"github.com/zhangga/luban/internal/rawrefs"
	"github.com/zhangga/luban/internal/utils"
	"github.com/zhangga/luban/pkg/logger"
	"strings"
)

const (
	enumKeyName        = "name"
	enumKeyFlags       = "flags"
	enumKeyComment     = "comment"
	enumKeyTags        = "tags"
	enumKeyUnique      = "unique"
	enumKeyGroup       = "group"
	enumItemKeyName    = "name"
	enumItemKeyValue   = "value"
	enumItemKeyAlias   = "alias"
	enumItemKeyComment = "comment"
	enumItemKeyTags    = "tags"
)

var (
	enumRequiredAttrs     = []string{enumKeyName}
	enumOptionalAttrs     = []string{enumKeyFlags, enumKeyComment, enumKeyTags, enumKeyUnique, enumKeyGroup}
	enumItemRequiredAttrs = []string{enumItemKeyName}
	enumItemOptionalAttrs = []string{enumItemKeyValue, enumItemKeyAlias, enumItemKeyComment, enumItemKeyTags}
)

var _ schema.ISchemaLoader = (*XmlSchemaLoader)(nil)

type ITagHandler func(ele *etree.Element)

// XmlSchemaLoader xml文件加载器
type XmlSchemaLoader struct {
	logger         logger.Logger
	dataType       string
	collector      schema.ISchemaCollector
	fileName       string
	tagHandlers    map[string]ITagHandler
	namespaceStack []string
}

func NewXmlSchemaLoader(logger logger.Logger, dataType string, collector schema.ISchemaCollector) schema.ISchemaLoader {
	l := &XmlSchemaLoader{
		logger:      logger,
		dataType:    dataType,
		collector:   collector,
		tagHandlers: map[string]ITagHandler{},
	}
	l.tagHandlers["module"] = l.AddModule
	l.tagHandlers["enum"] = l.AddEnum
	l.tagHandlers["bean"] = l.AddBean
	l.tagHandlers["table"] = l.AddTable
	l.tagHandlers["refgroup"] = l.AddRefGroup
	return l
}

func (l *XmlSchemaLoader) Load(fileName string) {
	l.fileName = fileName
	doc := etree.NewDocument()
	if err := doc.ReadFromFile(fileName); err != nil {
		l.logger.Errorf("read xml file: %s failed: %v", fileName, err)
		panic(err)
	}
	root := doc.Root()
	handler, ok := l.tagHandlers[root.Tag]
	if !ok {
		panic(fmt.Errorf("xml handler tag: %s, not found", root.Tag))
	}
	handler(root)
}

func (l *XmlSchemaLoader) AddModule(ele *etree.Element) {
	name := strings.TrimSpace(ele.SelectAttr("name").Value)
	if len(l.namespaceStack) == 0 {
		l.namespaceStack = append(l.namespaceStack, name)
	} else {
		l.namespaceStack = append(l.namespaceStack, utils.MakeFullName(l.namespaceStack[len(l.namespaceStack)-1], name))
	}

	// 加载所有module定义，允许嵌套
	for _, child := range ele.ChildElements() {
		tagName := child.Tag
		handler, ok := l.tagHandlers[tagName]
		if !ok {
			panic(fmt.Errorf("xml定义文件: %s, module: %s, 不支持的tag: %s", l.fileName, l.namespaceStack[len(l.namespaceStack)-1], tagName))
		}
		handler(child)
	}
	l.namespaceStack = l.namespaceStack[:len(l.namespaceStack)-1]
}

func (l *XmlSchemaLoader) AddEnum(ele *etree.Element) {
	validAttrKey(l.fileName, ele, enumOptionalAttrs, enumRequiredAttrs)

	rawEnum := &rawrefs.RawEnum{
		Name:           getRequiredAttr(l.fileName, ele, enumKeyName),
		Namespace:      l.namespaceStack[len(l.namespaceStack)-1],
		Comment:        getOptionalAttr(ele, enumKeyComment),
		IsFlags:        getOptionalBoolAttr(ele, enumKeyFlags),
		Tags:           utils.ParseAttrs(getOptionalAttr(ele, enumKeyTags)),
		IsUniqueItemId: getOptionalBoolAttr(ele, enumKeyUnique, true),
		Groups:
	}
	rawEnum.FullName = utils.MakeFullName(rawEnum.Namespace, rawEnum.Name)
}

func (l *XmlSchemaLoader) AddBean(ele *etree.Element) {

}

func (l *XmlSchemaLoader) AddTable(ele *etree.Element) {

}

func (l *XmlSchemaLoader) AddRefGroup(ele *etree.Element) {

}

func validAttrKey(defineFile string, ele *etree.Element, optionKeys []string, requireKeys []string) {
	for _, attr := range ele.Attr {
		name := attr.Key
		if !utils.Contain(requireKeys, name) && !utils.Contain(optionKeys, name) {
			panic(fmt.Errorf("xml定义文件: %s, tag: %s, 不支持的属性: %s", defineFile, ele.Tag, name))
		}
	}
	for _, key := range requireKeys {
		if ele.SelectAttr(key) == nil {
			panic(fmt.Errorf("xml定义文件: %s, tag: %s, 缺少必要属性: %s", defineFile, ele.Tag, key))
		}
	}
}

func getRequiredAttr(fileName string, ele *etree.Element, key string) string {
	attr := ele.SelectAttr(key)
	if attr == nil {
		panic(fmt.Errorf("xml定义文件: %s, tag: %s, 缺少必要属性: %s", fileName, ele.Tag, key))
	}
	value := strings.TrimSpace(attr.Value)
	if len(value) == 0 {
		panic(fmt.Errorf("xml定义文件: %s, tag: %s, 必要属性: %s, 不能为空", fileName, ele.Tag, key))
	}
	return value
}

func getOptionalAttr(ele *etree.Element, key string) string {
	attr := ele.SelectAttr(key)
	if attr == nil {
		return ""
	}
	return strings.TrimSpace(attr.Value)
}

func getOptionalBoolAttr(ele *etree.Element, key string, defaultValues ...bool) bool {
	if len(defaultValues) > 1 {
		panic("getOptionalBoolAttr: too many default values")
	}

	defaultValue := false
	if len(defaultValues) == 1 {
		defaultValue = defaultValues[0]
	}

	attr := ele.SelectAttr(key)
	if attr == nil {
		return defaultValue
	}
	value := strings.TrimSpace(attr.Value)
	if len(value) == 0 {
		return defaultValue
	}
	lv := strings.ToLower(value)
	return lv == "1" || lv == "true"
}
