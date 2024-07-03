package dataloader

import "github.com/zhangga/luban/core/refs"

type IDataLoader interface {
	Names() []string
	RawUrl() string
	ReadOne(ttype refs.TType) *refs.Record
	ReadMulti(ttype refs.TType) []*refs.Record
	Load(rawUrl, subAsset string) error
}
