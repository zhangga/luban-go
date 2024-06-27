package core

import (
	"github.com/zhangga/luban/core/codeformat"
	"github.com/zhangga/luban/core/codetarget"
	"github.com/zhangga/luban/core/dataloader"
	"github.com/zhangga/luban/core/datatarget"
	"github.com/zhangga/luban/core/l10n"
	"github.com/zhangga/luban/core/manager"
	"github.com/zhangga/luban/core/outputsaver"
	"github.com/zhangga/luban/core/pipeline"
	"github.com/zhangga/luban/core/postprocess"
	"github.com/zhangga/luban/core/schema"
	"github.com/zhangga/luban/core/template"
	"github.com/zhangga/luban/core/validator"
	pipeline2 "github.com/zhangga/luban/internal/pipeline"
	schema2 "github.com/zhangga/luban/internal/schema"
)

// init 注册管理器
func init() {
	manager.Register[*schema.Manager]()
	manager.Register[*template.Manager]()
	manager.Register[*codeformat.Manager]()
	manager.Register[*codetarget.Manager]()
	manager.Register[*postprocess.Manager]()
	manager.Register[*outputsaver.Manager]()
	manager.Register[*dataloader.Manager]()
	manager.Register[*validator.Manager]()
	manager.Register[*datatarget.Manager]()
	manager.Register[*pipeline.Manager]()
	manager.Register[*l10n.Manager]()
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
