package validator

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/zhangga/luban-go/core/datas"
)

func TestPathValidator(t *testing.T) {
	// 创建一个临时文件用于测试
	tmpDir := os.TempDir()
	testFile := filepath.Join(tmpDir, "test_resource.png")
	os.WriteFile(testFile, []byte("dummy"), 0644)
	defer os.Remove(testFile)

	v := &PathValidator{BaseDir: tmpDir}

	// Test valid file
	err := v.Validate(nil, datas.NewDString("test_resource.png"), "")
	if err != nil {
		t.Errorf("Expected nil for existing file, got %v", err)
	}

	// Test invalid file
	err = v.Validate(nil, datas.NewDString("non_existent.png"), "")
	if err == nil {
		t.Errorf("Expected error for non-existent file")
	}

	// Test file extension validation (success)
	err = v.Validate(nil, datas.NewDString("test_resource.png"), "*.png")
	if err != nil {
		t.Errorf("Expected nil for matching extension, got %v", err)
	}

	// Test file extension validation (failure)
	err = v.Validate(nil, datas.NewDString("test_resource.png"), "*.jpg")
	if err == nil {
		t.Errorf("Expected error for mismatched extension")
	}
}
