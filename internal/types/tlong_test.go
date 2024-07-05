package types_test

import (
	"github.com/zhangga/luban/internal/types"
	"math"
	"strconv"
	"testing"
)

func TestTLong_TryParseFrom(t *testing.T) {
	tb := &types.TLong{}

	tests := []struct {
		str      string
		expected bool
	}{
		{"0", true},
		{strconv.Itoa(math.MaxInt64), true},
		{strconv.Itoa(math.MaxInt64) + "1", false},
		{strconv.Itoa(math.MinInt64), true},
		{strconv.Itoa(math.MinInt64) + "1", false},
	}

	for _, tt := range tests {
		t.Run(tt.str, func(t *testing.T) {
			ok := tb.TryParseFrom(tt.str)
			if ok != tt.expected {
				t.Errorf("TLong.TryParseFrom(%q) = %v; want %v", tt.str, ok, tt.expected)
			}
		})
	}
}
