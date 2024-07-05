package types_test

import (
	"github.com/zhangga/luban/internal/types"
	"testing"
)

func TestTDouble_TryParseFrom(t *testing.T) {
	tb := &types.TDouble{}

	tests := []struct {
		str      string
		expected bool
	}{
		{"0", true},
		{"0.0", true},
		{"111.12345678", true},
		{"-111.12345678", true},
		{"haha", false},
	}

	for _, tt := range tests {
		t.Run(tt.str, func(t *testing.T) {
			ok := tb.TryParseFrom(tt.str)
			if ok != tt.expected {
				t.Errorf("TDouble.TryParseFrom(%q) = %v; want %v", tt.str, ok, tt.expected)
			}
		})
	}
}
