package defs

import (
	"fmt"
	"github.com/zhangga/luban/core/refs"
	"github.com/zhangga/luban/internal/rawdefs"
	"github.com/zhangga/luban/internal/utils"
)

var _ refs.IDefField = (*DefField)(nil)

type DefField struct {
	Assembly             *DefAssembly
	HostType             *DefBean
	name                 string
	Type                 string
	CType                refs.TType
	Comment              string
	AutoId               int
	Tags                 map[string]string
	IgnoreNameValidation bool
	Groups               []string
	RawField             *rawdefs.RawField
}

func NewDefField(host *DefBean, f *rawdefs.RawField, idOffset int) *DefField {
	return &DefField{
		HostType:             host,
		name:                 f.Name,
		Type:                 f.Type,
		Comment:              f.Comment,
		Tags:                 f.Tags,
		IgnoreNameValidation: f.NotNameValidation,
		Groups:               f.Groups,
		RawField:             f,
	}
}

func (f *DefField) Name() string {
	return f.name
}

func (f *DefField) SetAutoId(id int) {
	f.AutoId = id
}

func (f *DefField) IsNullable() bool {
	return f.CType.IsNullable()
}

func (f *DefField) HasTag(attrName string) bool {
	if len(f.Tags) == 0 {
		return false
	}
	_, ok := f.Tags[attrName]
	return ok
}

func (f *DefField) GetTag(attrName string) (string, bool) {
	if len(f.Tags) == 0 {
		return "", false
	}
	v, ok := f.Tags[attrName]
	return v, ok
}

func (f *DefField) NeedExport() bool {
	return f.Assembly.NeedExport(f.Groups)
}

func (f *DefField) String() string {
	return fmt.Sprintf("%s.%s", f.HostType.FullName(), f.Name)
}

func (f *DefField) Compile() {
	panic("implement me")
}

func (f *DefField) PostCompile() {
	panic("implement me")
}

func CompileFields[T refs.IDefField](hostType *DefTypeBase, fields []T) {
	names := make(map[string]struct{})
	nextAutoId := 1
	for _, field := range fields {
		name := field.Name()
		if len(name) == 0 {
			panic(fmt.Errorf("type: %s field name can't be empty", hostType.FullName()))
		}
		if _, ok := names[name]; ok {
			panic(fmt.Errorf("type: %s, field name: %s duplicate", hostType.FullName(), name))
		}
		if utils.ToCsStyleName(name) == hostType.Name {
			panic(fmt.Errorf("type: %s, field name: %s 生成的c#字段名与类型名相同，会引起编译错误", hostType.FullName(), name))
		}
		field.SetAutoId(nextAutoId)
		nextAutoId++
	}

	for _, field := range fields {
		field.Compile()
	}
}
