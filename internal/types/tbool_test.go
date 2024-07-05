package types_test

import (
	"github.com/zhangga/luban/internal/types"
	"testing"
)

func TestTBool_TryParseFrom(t *testing.T) {
	tb := &types.TBool{}

	tests := []struct {
		str      string
		expected bool
	}{
		{"1", true},
		{"0", true},
		{"true", true},
		{"false", true},
		{"True", true},
		{"False", true},
		{"No", false},
		{"Yes", false},
		{"NaB", false},
	}

	for _, tt := range tests {
		t.Run(tt.str, func(t *testing.T) {
			ok := tb.TryParseFrom(tt.str)
			if ok != tt.expected {
				t.Errorf("TBool.TryParseFrom(%q) = %v; want %v", tt.str, ok, tt.expected)
			}
		})
	}
}
