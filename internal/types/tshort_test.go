package types_test

import (
	"github.com/zhangga/luban/internal/types"
	"math"
	"strconv"
	"testing"
)

func TestTShort_TryParseFrom(t *testing.T) {
	tb := &types.TShort{}

	tests := []struct {
		str      string
		expected bool
	}{
		{"0", true},
		{strconv.Itoa(math.MaxInt16), true},
		{strconv.Itoa(math.MaxInt16 + 1), false},
		{strconv.Itoa(math.MinInt16), true},
		{strconv.Itoa(math.MinInt16 - 1), false},
	}

	for _, tt := range tests {
		t.Run(tt.str, func(t *testing.T) {
			ok := tb.TryParseFrom(tt.str)
			if ok != tt.expected {
				t.Errorf("TShort.TryParseFrom(%q) = %v; want %v", tt.str, ok, tt.expected)
			}
		})
	}
}
