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

var _ schema.ISchemaLoader = (*XmlSchemaLoader)(nil)

type ITagHandler func(ele *etree.Element)

// XmlSchemaLoader xml文件加载器
type XmlSchemaLoader struct {
	schemaLoaderBase
	fileName       string
	tagHandlers    map[string]ITagHandler
	namespaceStack []string
}

func NewXmlSchemaLoader(logger logger.Logger, dataType string, collector schema.ISchemaCollector) schema.ISchemaLoader {
	l := &XmlSchemaLoader{
		schemaLoaderBase: schemaLoaderBase{
			logger:    logger,
			dataType:  dataType,
			collector: collector,
		},
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
		l.namespaceStack = append(l.namespaceStack, utils.MakeFullName(l.CurNameSpace(), name))
	}

	// 加载所有module定义，允许嵌套
	for _, child := range ele.ChildElements() {
		tagName := child.Tag
		handler, ok := l.tagHandlers[tagName]
		if !ok {
			panic(fmt.Errorf("xml定义文件: %s, module: %s, 不支持的tag: %s", l.fileName, l.CurNameSpace(), tagName))
		}
		handler(child)
	}
	l.namespaceStack = l.namespaceStack[:len(l.namespaceStack)-1]
}

const (
	enumKeyName    = "name"
	enumKeyFlags   = "flags"
	enumKeyComment = "comment"
	enumKeyTags    = "tags"
	enumKeyUnique  = "unique"
	enumKeyGroup   = "group"

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

func (l *XmlSchemaLoader) AddEnum(ele *etree.Element) {
	validAttrKeys(l.fileName, ele, enumOptionalAttrs, enumRequiredAttrs)

	rawEnum := &rawrefs.RawEnum{
		Name:           getRequiredAttr(l.fileName, ele, enumKeyName),
		Namespace:      l.CurNameSpace(),
		Comment:        getOptionalAttr(ele, enumKeyComment),
		IsFlags:        getOptionalBoolAttr(ele, enumKeyFlags),
		Tags:           utils.ParseAttrs(getOptionalAttr(ele, enumKeyTags)),
		IsUniqueItemId: getOptionalBoolAttr(ele, enumKeyUnique, true),
		Groups:         createGroups(getOptionalAttr(ele, enumKeyGroup)),
	}
	rawEnum.FullName = utils.MakeFullName(rawEnum.Namespace, rawEnum.Name)

	for _, itemEle := range ele.ChildElements() {
		switch itemEle.Tag {
		case "var":
			validAttrKeys(l.fileName, itemEle, enumItemOptionalAttrs, enumItemRequiredAttrs)
			rawEnum.Items = append(rawEnum.Items, rawrefs.EnumItem{
				Name:    getRequiredAttr(l.fileName, itemEle, enumItemKeyName),
				Alias:   getOptionalAttr(itemEle, enumItemKeyAlias),
				Value:   getOptionalAttr(itemEle, enumItemKeyValue),
				Comment: getOptionalAttr(itemEle, enumItemKeyComment),
				Tags:    utils.ParseAttrs(getOptionalAttr(itemEle, enumItemKeyTags)),
			})
		case "mapper":
			rawEnum.TypeMappers = append(rawEnum.TypeMappers, createTypeMapper(itemEle, l.fileName, rawEnum.FullName))
		default:
			panic(fmt.Errorf("xml定义文件: %s, enum: %s, 不支持的tag: %s", l.fileName, rawEnum.Name, itemEle.Tag))
		}
	}
	l.collector.AddEnum(rawEnum)
	l.logger.Debugf("load enum: %s", rawEnum.FullName)
}

const (
	tableKeyName               = "name"
	tableKeyValue              = "value"
	tableKeyInput              = "input"
	tableKeyIndex              = "index"
	tableKeyMode               = "mode"
	tableKeyGroup              = "group"
	tableKeyComment            = "comment"
	tableKeyReadSchemaFromFile = "readSchemaFromFile"
	tableKeyOutput             = "output"
	tableKeyTags               = "tags"
)

var (
	tableRequiredAttrs = []string{tableKeyName, tableKeyValue, tableKeyInput}
	tableOptionalAttrs = []string{tableKeyIndex, tableKeyMode, tableKeyGroup, tableKeyComment, tableKeyReadSchemaFromFile, tableKeyOutput, tableKeyTags}
)

func (l *XmlSchemaLoader) AddTable(ele *etree.Element) {
	validAttrKeys(l.fileName, ele, tableOptionalAttrs, tableRequiredAttrs)

	name := getRequiredAttr(l.fileName, ele, tableKeyName)
	module := l.CurNameSpace()
	valueType := getRequiredAttr(l.fileName, ele, tableKeyValue)
	input := getRequiredAttr(l.fileName, ele, tableKeyInput)
	defineFromFile := getOptionalBoolAttr(ele, tableKeyReadSchemaFromFile)
	index := getOptionalAttr(ele, tableKeyIndex)
	group := getOptionalAttr(ele, tableKeyGroup)
	comment := getOptionalAttr(ele, tableKeyComment)
	mode := getOptionalAttr(ele, tableKeyMode)
	tags := getOptionalAttr(ele, tableKeyTags)
	output := getOptionalAttr(ele, tableKeyOutput)
	table := &rawrefs.RawTable{
		Name:               name,
		Namespace:          module,
		ValueType:          valueType,
		ReadSchemaFromFile: defineFromFile,
		Index:              index,
		Groups:             createGroups(group),
		Comment:            comment,
		Mode:               rawrefs.ConvertTableMode(l.fileName, name, mode, index),
		Tags:               utils.ParseAttrs(tags),
		OutputFile:         output,
	}
	if len(table.Name) == 0 {
		panic(fmt.Errorf("xml定义文件: %s, table: %s, name属性不能为空", l.fileName, ele.Tag))
	}
	if len(table.ValueType) == 0 {
		panic(fmt.Errorf("xml定义文件: %s, table: %s, value属性不能为空", l.fileName, ele.Tag))
	}
	inputs := strings.FieldsFunc(input, func(c rune) bool {
		return c == ',' || c == ';' || c == '|'
	})
	for _, in := range inputs {
		if s := strings.TrimSpace(in); len(s) > 0 {
			table.InputFiles = append(table.InputFiles, s)
		}
	}
	l.collector.AddTable(table)
}

const (
	beanKeyName      = "name"
	beanKeyParent    = "parent"
	beanKeyValueType = "valueType"
	beanKeyAlias     = "alias"
	beanKeySep       = "sep"
	beanKeyComment   = "comment"
	beanKeyTags      = "tags"
	beanKeyGroup     = "group"
)

var (
	beanRequiredAttrs = []string{beanKeyName}
	beanOptionalAttrs = []string{beanKeyParent, beanKeyValueType, beanKeyAlias, beanKeySep, beanKeyComment, beanKeyTags, beanKeyGroup}
)

func (l *XmlSchemaLoader) AddBean(ele *etree.Element) {
	l.AddBeanByParent(ele, "")
}

func (l *XmlSchemaLoader) AddBeanByParent(ele *etree.Element, parent string) {
	validAttrKeys(l.fileName, ele, beanOptionalAttrs, beanRequiredAttrs)
	l.tryUpdateParent(ele, &parent)
	rawBean := &rawrefs.RawBean{
		Name:        getRequiredAttr(l.fileName, ele, beanKeyName),
		Namespace:   l.CurNameSpace(),
		Parent:      parent,
		IsValueType: getOptionalBoolAttr(ele, beanKeyValueType),
		Alias:       getOptionalAttr(ele, beanKeyAlias),
		Sep:         getOptionalAttr(ele, beanKeySep),
		Comment:     getOptionalAttr(ele, beanKeyComment),
		Tags:        utils.ParseAttrs(getOptionalAttr(ele, beanKeyTags)),
		Groups:      createGroups(getOptionalAttr(ele, beanKeyGroup)),
	}
	rawBean.FullName = utils.MakeFullName(rawBean.Namespace, rawBean.Name)

	var childBeans []*etree.Element
	defineAnyChildBean := false
	for _, itemEle := range ele.ChildElements() {
		switch itemEle.Tag {
		case "var":
			if defineAnyChildBean {
				panic(fmt.Errorf("xml定义文件: %s, bean: %s 的多态子bean必须在所有成员字段 <var> 之后定义", l.fileName, rawBean.FullName))
			}
			field := createFieldByElement(itemEle, l.fileName)
			rawBean.Fields = append(rawBean.Fields, field)
		case "mapper":
			rawBean.TypeMappers = append(rawBean.TypeMappers, createTypeMapper(itemEle, l.fileName, rawBean.FullName))
		case "bean":
			defineAnyChildBean = true
			childBeans = append(childBeans, itemEle)
		default:
			panic(fmt.Errorf("xml定义文件: %s, bean: %s, 不支持的tag: %s", l.fileName, rawBean.FullName, itemEle.Tag))
		}
	}
	l.collector.AddBean(rawBean)
	l.logger.Debugf("load bean: %s", rawBean.FullName)

	for _, child := range childBeans {
		l.AddBeanByParent(child, rawBean.FullName)
	}
}

const (
	groupKeyName = "name"
	groupKeyRef  = "ref"
)

var (
	groupRequiredAttrs = []string{groupKeyName, groupKeyRef}
)

func (l *XmlSchemaLoader) AddRefGroup(ele *etree.Element) {
	validAttrKeys(l.fileName, ele, nil, groupRequiredAttrs)

	refSplits := strings.FieldsFunc(getRequiredAttr(l.fileName, ele, groupKeyRef), func(r rune) bool {
		return r == ',' || r == ';' || r == '|'
	})
	var refs []string
	for _, ref := range refSplits {
		if r := strings.TrimSpace(ref); len(r) > 0 {
			refs = append(refs, r)
		}
	}
	rawRefGroup := &rawrefs.RawRefGroup{
		Name: getRequiredAttr(l.fileName, ele, groupKeyName),
		Refs: refs,
	}
	l.collector.AddRefGroup(rawRefGroup)
}

func (l *XmlSchemaLoader) CurNameSpace() string {
	if len(l.namespaceStack) == 0 {
		return ""
	}
	return l.namespaceStack[len(l.namespaceStack)-1]
}

func (l *XmlSchemaLoader) tryUpdateParent(ele *etree.Element, parent *string) {
	selfDefParent := getOptionalAttr(ele, beanKeyParent)
	if len(selfDefParent) > 0 {
		if len(*parent) > 0 {
			panic(fmt.Errorf("xml定义文件: %s, 嵌套在: %s, 中定义的子bean: %s, 不能再定义parent属性: %s", l.fileName, parent, ele.Tag, selfDefParent))
		}
		*parent = selfDefParent
	}
}

func validAttrKeys(defineFile string, ele *etree.Element, optionKeys []string, requireKeys []string) {
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

func createGroups(str string) []string {
	split := func(c rune) bool {
		return c == ',' || c == ';'
	}
	parts := strings.FieldsFunc(str, split)

	var groups []string
	for _, part := range parts {
		trimmed := strings.TrimSpace(part)
		if len(trimmed) > 0 {
			groups = append(groups, trimmed)
		}
	}
	return groups
}

const (
	fieldKeyName    = "name"
	fieldKeyType    = "type"
	fieldKeyGroup   = "group"
	fieldKeyComment = "comment"
	fieldKeyTags    = "tags"
)

var (
	fieldRequiredAttrs = []string{fieldKeyName, fieldKeyType}
	fieldOptionalAttrs = []string{fieldKeyGroup, fieldKeyComment, fieldKeyTags}
)

// createFieldByElement 创建字段
func createFieldByElement(ele *etree.Element, fileName string) rawrefs.RawField {
	validAttrKeys(fileName, ele, fieldOptionalAttrs, fieldRequiredAttrs)
	name := getRequiredAttr(fileName, ele, fieldKeyName)
	typeStr := getRequiredAttr(fileName, ele, fieldKeyType)
	group := getOptionalAttr(ele, fieldKeyGroup)
	comment := getOptionalAttr(ele, fieldKeyComment)
	tags := getOptionalAttr(ele, fieldKeyTags)
	return createField(fileName, name, typeStr, group, comment, tags, false)
}

func createField(fileName, name, typeStr, group, comment, tags string, ignoreNameValidation bool) rawrefs.RawField {
	field := rawrefs.RawField{
		Name:              name,
		Groups:            createGroups(group),
		Comment:           comment,
		Tags:              utils.ParseAttrs(tags),
		NotNameValidation: ignoreNameValidation,
		Type:              typeStr,
	}
	return field
}

const (
	enumItemMapperKeyTarget     = "target"
	enumItemMapperKeyCodeTarget = "codeTarget"
	enumItemMapperKeyName       = "name"
	enumItemMapperKeyValue      = "value"
)

// createTypeMapper 创建类型映射
func createTypeMapper(ele *etree.Element, fileName, fullName string) rawrefs.TypeMapper {
	target := getRequiredAttr(fileName, ele, enumItemMapperKeyTarget)
	targets := strings.FieldsFunc(target, func(c rune) bool {
		return c == ',' || c == ';' || c == '|'
	})
	codeTarget := getRequiredAttr(fileName, ele, enumItemMapperKeyCodeTarget)
	codeTargets := strings.FieldsFunc(codeTarget, func(c rune) bool {
		return c == ',' || c == ';' || c == '|'
	})
	opts := make(map[string]string)
	for _, item := range ele.ChildElements() {
		key := getRequiredAttr(fileName, item, enumItemMapperKeyName)
		value := getRequiredAttr(fileName, item, enumItemMapperKeyValue)
		if _, ok := opts[key]; ok {
			panic(fmt.Errorf("xml定义文件: %s, enum: %s, mapper: %s, 重复的name属性: %s", fileName, fullName, ele.Tag, key))
		}
		opts[key] = value
	}

	return rawrefs.TypeMapper{
		Targets:     targets,
		CodeTargets: codeTargets,
		Options:     opts,
	}
}
