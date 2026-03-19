package defs

import "github.com/zhangga/luban-go/core/datas"

type Record struct {
	AutoIndex int
	Data      *datas.DBean
	Source    string
	Tags      []string
}

func NewRecord(data *datas.DBean, source string, tags []string) *Record {
	return &Record{
		Data:   data,
		Source: source,
		Tags:   tags,
	}
}

func (r *Record) IsNotFiltered(includeTags []string, excludeTags []string) bool {
	if len(r.Tags) == 0 {
		return true
	}

	if len(includeTags) > 0 {
		for _, tag := range r.Tags {
			for _, includeTag := range includeTags {
				if tag == includeTag {
					return true
				}
			}
		}
		return false
	}

	for _, tag := range r.Tags {
		for _, excludeTag := range excludeTags {
			if tag == excludeTag {
				return false
			}
		}
	}
	return true
}
