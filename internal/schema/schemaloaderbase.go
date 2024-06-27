package schema

import (
	"github.com/zhangga/luban/core/schema"
	"github.com/zhangga/luban/pkg/logger"
)

type schemaLoaderBase struct {
	logger    logger.Logger
	dataType  string
	collector schema.ISchemaCollector
}
