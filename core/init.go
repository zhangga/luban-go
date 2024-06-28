package core

import (
	"github.com/zhangga/luban/core/manager"
	"github.com/zhangga/luban/core/pipeline"
	"github.com/zhangga/luban/core/schema"
	codeformat2 "github.com/zhangga/luban/internal/codeformat"
	codetarget2 "github.com/zhangga/luban/internal/codetarget"
	dataloader2 "github.com/zhangga/luban/internal/dataloader"
	datatarget2 "github.com/zhangga/luban/internal/datatarget"
	l10n2 "github.com/zhangga/luban/internal/l10n"
	outputsaver2 "github.com/zhangga/luban/internal/outputsaver"
	pipeline2 "github.com/zhangga/luban/internal/pipeline"
	postprocess2 "github.com/zhangga/luban/internal/postprocess"
	schema2 "github.com/zhangga/luban/internal/schema"
	template2 "github.com/zhangga/luban/internal/template"
	validator2 "github.com/zhangga/luban/internal/validator"
)

// init 注册管理器
func init() {
	manager.Register[*schema2.Manager]()
	manager.Register[*template2.Manager]()
	manager.Register[*codeformat2.Manager]()
	manager.Register[*codetarget2.Manager]()
	manager.Register[*postprocess2.Manager]()
	manager.Register[*outputsaver2.Manager]()
	manager.Register[*dataloader2.Manager]()
	manager.Register[*validator2.Manager]()
	manager.Register[*datatarget2.Manager]()
	manager.Register[*pipeline2.Manager]()
	manager.Register[*l10n2.Manager]()
}

// init 注册pipeline
func init() {
	pipeline.Register(pipeline2.NewDefaultPipeline)
}

// init 注册schema
func init() {
	// 注册collector
	schema.RegisterCollector(schema2.NewDefaultSchemaCollector)
	// 注册schema loader
	schema.RegisterSchemaLoader(schema2.NewXmlSchemaLoader, 0, "", ".xml")
	schema.RegisterSchemaLoader(schema2.NewExcelSchemaLoader, 0, "table", ".xlsx", ".xls", ".xlsm", ".csv")
	schema.RegisterSchemaLoader(schema2.NewExcelSchemaLoader, 0, "bean", ".xlsx", ".xls", ".xlsm", ".csv")
	schema.RegisterSchemaLoader(schema2.NewExcelSchemaLoader, 0, "enum", ".xlsx", ".xls", ".xlsm", ".csv")
	// 注册bean loader
	schema.RegisterBeanLoaderCreator(schema2.NewBeanSchemaFromExcelHeaderLoader)
}
