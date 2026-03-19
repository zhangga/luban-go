package dataloader

import (
	"strings"
)

// DataStream 用于解析以分隔符(如 `,` 或 `;`)分割的单元格数据
type DataStream struct {
	tokens []string
	cursor int
}

func NewDataStream(content string, sep string) *DataStream {
	if content == "" {
		return &DataStream{tokens: []string{}, cursor: 0}
	}

	if sep == "" {
		sep = "," // 默认使用逗号分割
	}

	rawTokens := strings.Split(content, sep)
	tokens := make([]string, 0, len(rawTokens))
	for _, t := range rawTokens {
		tokens = append(tokens, strings.TrimSpace(t))
	}

	return &DataStream{
		tokens: tokens,
		cursor: 0,
	}
}

func (s *DataStream) ReadString() (string, bool) {
	if s.cursor >= len(s.tokens) {
		return "", false
	}
	val := s.tokens[s.cursor]
	s.cursor++
	return val, true
}

func (s *DataStream) HasNext() bool {
	return s.cursor < len(s.tokens)
}

func (s *DataStream) Remain() []string {
	if s.cursor >= len(s.tokens) {
		return []string{}
	}
	return s.tokens[s.cursor:]
}
