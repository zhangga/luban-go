package codetarget

import (
	"github.com/zhangga/luban-go/core/defs"
)

// OutputFile 定义代码或数据导出的文件
type OutputFile struct {
	File     string
	Content  []byte
	Encoding string // 记录编码格式 (如 "utf8")
}

func NewOutputFile(file string, content []byte) *OutputFile {
	return &OutputFile{
		File:    file,
		Content: content,
	}
}

// OutputFileManifest 聚合所有的导出文件
type OutputFileManifest struct {
	DataFiles []*OutputFile
	CodeFiles []*OutputFile
}

func NewOutputFileManifest() *OutputFileManifest {
	return &OutputFileManifest{
		DataFiles: make([]*OutputFile, 0),
		CodeFiles: make([]*OutputFile, 0),
	}
}

func (m *OutputFileManifest) AddDataFile(f *OutputFile) {
	m.DataFiles = append(m.DataFiles, f)
}

func (m *OutputFileManifest) AddCodeFile(f *OutputFile) {
	m.CodeFiles = append(m.CodeFiles, f)
}

// ICodeTarget 定义代码生成目标（例如：csharp-json，go-bin 等）
type ICodeTarget interface {
	Name() string
	ValidateDefinition(assembly *defs.DefAssemblyImpl) error
	Handle(assembly *defs.DefAssemblyImpl, manifest *OutputFileManifest) error
}
