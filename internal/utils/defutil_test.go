package utils_test

import (
	"github.com/zhangga/luban/internal/utils"
	"reflect"
	"testing"
)

// TestParseAttrs 测试 ParseAttrs 函数
func TestParseAttrs(t *testing.T) {
	tests := []struct {
		tags     string
		expected map[string]string
	}{
		{"key", map[string]string{"key": "key"}},
		{"key:value", map[string]string{"key": "value"}},
		{"(key=value)", map[string]string{"key": "value"}},
		{"(key=value)#(another=pair)", map[string]string{"key": "value", "another": "pair"}},
		{"#(key=value)", map[string]string{"": "", "key": "value"}},
		{"key=value#(another:pair)", map[string]string{"key": "value", "another": "pair"}},
		{"", map[string]string{}}, // Empty string
		// Add more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.tags, func(t *testing.T) {
			actual := utils.ParseAttrs(tt.tags)
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("ParseAttrs(%q) = %v, expected %v", tt.tags, actual, tt.expected)
			}
		})
	}
}

func TestTrimBracePairs(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"Empty String", "", ""},
		{"No Braces", "example", "example"},
		{"Single Pair", "(example)", "example"},
		{"Multiple Pairs", "((example))", "example"},
		{"Unbalanced Braces", "(example", "(example"},
		{"Braces Inside", "(ex(amp)le)", "ex(amp)le"},
		{"Mismatched Braces", "(example))", "(example))"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := utils.TrimBracePairs(tt.input)
			if actual != tt.expected {
				t.Errorf("TrimBracePairs(%q) = %q, expected %q", tt.input, actual, tt.expected)
			}
		})
	}
}
