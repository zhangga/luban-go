package dataloader

import "github.com/zhangga/luban/core/refs"

type IDataLoader interface {
	RawUrl() string
	ReadOne() refs.Record
}
