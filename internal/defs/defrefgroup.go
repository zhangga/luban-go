package defs

import "github.com/zhangga/luban/internal/rawdefs"

type DefRefGroup struct {
	Name string   `json:"name"`
	Refs []string `json:"refs"`
}

func NewDefRefGroup(group *rawdefs.RawRefGroup) *DefRefGroup {
	return &DefRefGroup{
		Name: group.Name,
		Refs: group.Refs,
	}
}
