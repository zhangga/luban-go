package excel

import (
	"fmt"
	"github.com/zhangga/luban/internal/utils"
	"golang.org/x/exp/slices"
	"strings"
)

const (
	tagKeySep       = "sep"
	tagKeyNonEmpty  = "non_empty"
	tagKeyMultiRows = "multi_rows"
	tagKeyDefault   = "default"
)

var validTags = map[string]struct{}{
	tagKeySep:       struct{}{},
	tagKeyNonEmpty:  struct{}{},
	tagKeyMultiRows: struct{}{},
	tagKeyDefault:   struct{}{},
}

type Title struct {
	Root                  bool
	FromIndex             int
	ToIndex               int
	Name                  string
	Sep                   string
	Tags                  map[string]string
	SubTitles             map[string]*Title
	SubTitleList          []*Title
	NonEmpty              bool
	Default               string
	SelfMultiRows         bool
	HierarchyMultiRows    bool
	SubHierarchyMultiRows bool
}

func (t *Title) Init() {
	t.SortSubTitles()

	t.Sep = utils.GetValueOrDefault(t.Tags, tagKeySep, "")
	nonEmpty := strings.ToLower(utils.GetValueOrDefault(t.Tags, tagKeyNonEmpty, ""))
	if nonEmpty == "1" || nonEmpty == "true" {
		t.NonEmpty = true
	} else {
		t.NonEmpty = false
	}
	multiRows := strings.ToLower(utils.GetValueOrDefault(t.Tags, tagKeyMultiRows, ""))
	if multiRows == "1" || multiRows == "true" {
		t.SelfMultiRows = true
	} else {
		t.SelfMultiRows = false
	}
	t.Default = utils.GetValueOrDefault(t.Tags, tagKeyDefault, "")

	for k := range t.Tags {
		if !utils.ContainKey(validTags, k) {
			panic(fmt.Errorf("excel标题列: %s, 不支持tag: %s, 请移到##type行", t.Name, k))
		}
	}

	if t.HasSubTitle() {
		if t.Root {
			firstField := utils.FirstOfList(t.SubTitleList, func(t *Title) bool {
				return utils.IsNormalFieldName(t.Name)
			})
			if firstField != nil {
				// 第一个字段一般为key，为了避免失误将空单元格当作key=0的数据，默认非空
				firstField.Tags["non_empty"] = "1"
			}
		}
		for _, sub := range t.SubTitleList {
			sub.Init()
		}
	}
	t.SubHierarchyMultiRows = utils.Any(t.SubTitleList, func(t *Title) bool {
		return t.HierarchyMultiRows
	})
	t.HierarchyMultiRows = t.SelfMultiRows || t.SubHierarchyMultiRows
}

func (t *Title) HasSubTitle() bool {
	return len(t.SubTitleList) > 0
}

func (t *Title) SepOrDefault(defaultValue string) string {
	if t.Sep == "" {
		return defaultValue
	}
	return t.Sep
}

func (t *Title) AddSubTitle(title *Title) {
	if t.SubTitles == nil {
		t.SubTitles = make(map[string]*Title)
	}
	if _, ok := t.SubTitles[title.Name]; ok {
		panic(fmt.Errorf("列: %s 重复", title.Name))
	}
	t.SubTitles[title.Name] = title
	t.SubTitleList = append(t.SubTitleList, title)
}

// SortSubTitles 由于先处理merge再处理只占一列的标题头.
// sub titles 未必是有序的。对于大多数数据并无影响
// 但对于 list类型的多级标题头，有可能导致element 数据次序乱了
func (t *Title) SortSubTitles() {
	if len(t.SubTitleList) > 1 {
		slices.SortFunc(t.SubTitleList, func(t1, t2 *Title) int {
			return t1.FromIndex - t2.FromIndex
		})
	}
	for _, st := range t.SubTitleList {
		st.SortSubTitles()
	}
}

func (t *Title) String() string {
	subTitles := make([]string, 0, len(t.SubTitleList))
	for _, st := range t.SubTitleList {
		subTitles = append(subTitles, st.String())
	}
	return fmt.Sprintf("name: %s [%d, %d], sub titles:[%s]", t.Name, t.FromIndex, t.ToIndex, strings.Join(subTitles, ",\\n"))
}
