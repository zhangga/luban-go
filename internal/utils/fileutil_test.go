package utils_test

import (
	"github.com/zhangga/luban/internal/utils"
	"testing"
)

func TestSplitFileAndSheetName(t *testing.T) {
	// 定义一组测试用例
	tests := []struct {
		url           string // 输入URL
		expectedFile  string // 期望的文件名结果
		expectedSheet string // 期望的表名结果
	}{
		{"file.xlsx@Sheet1", "file.xlsx", "Sheet1"},
		{"http://example.com/path/to/file.xlsx@Sheet1", "file.xlsx", "Sheet1"},
		{"https://example.com/path/to/file.xlsx@Sheet1", "file.xlsx", "Sheet1"},
		{"file.xlsx", "file.xlsx", ""},
		{"file@Sheet1", "file", "Sheet1"},
		{"file@Sheet1@Sheet2", "file", "Sheet1@Sheet2"},
		{"@Sheet1", "", "Sheet1"},
		{"path/to/file.xlsx", "path/to/file.xlsx", ""},
		{"path/to/@Sheet1", "", "Sheet1"},
	}

	// 遍历测试用例并运行测试
	for _, tt := range tests {
		t.Run(tt.url, func(t *testing.T) {
			file, sheet := utils.SplitFileAndSheetName(tt.url)
			if file != tt.expectedFile || sheet != tt.expectedSheet {
				t.Errorf("SplitFileAndSheetName(%q) = %q, %q; want %q, %q",
					tt.url, file, sheet, tt.expectedFile, tt.expectedSheet)
			}
		})
	}
}
