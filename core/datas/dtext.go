package datas

import "fmt"

// DText 表示本地化的文本数据
// 通常它可能包含一个 Key 和一个源语言的文本内容
type DText struct {
	Key        string
	RawText    string
}

func NewDText(key string, rawText string) *DText {
	return &DText{
		Key:     key,
		RawText: rawText,
	}
}

func (d *DText) TypeName() string {
	return "text"
}

func (d *DText) Accept(visitor IDataVisitor) interface{} {
	return visitor.VisitDText(d)
}

func (d *DText) String() string {
	return fmt.Sprintf("{key:%s, text:%s}", d.Key, d.RawText)
}
