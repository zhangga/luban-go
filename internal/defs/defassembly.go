package defs

import (
	"github.com/zhangga/luban/internal/rawdefs"
	"github.com/zhangga/luban/pkg/logger"
)

type DefAssembly struct {
	logger                 logger.Logger
	TypeMap                map[string]DefTypeBase
	TypeList               []DefTypeBase
	notCaseSenseTypes      map[string]DefTypeBase
	namespaces             map[string]struct{}
	notCaseSenseNamespaces map[string]DefTypeBase
	targets                []rawdefs.RawTarget
	Target                 rawdefs.RawTarget
}
