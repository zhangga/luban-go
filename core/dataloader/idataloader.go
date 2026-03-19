package dataloader

import (
	"github.com/zhangga/luban-go/core/defs"
	"github.com/zhangga/luban-go/core/types"
	"io"
)

type IDataLoader interface {
	RawUrl() string
	ReadOne(t *types.TBean) *defs.Record
	ReadMulti(t *types.TBean) []*defs.Record
	Load(rawUrl string, subAsset string, stream io.Reader) error
}
