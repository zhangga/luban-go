package validator

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/zhangga/luban-go/core/datas"
)

// PathValidator 处理资源路径校验
type PathValidator struct {
	BaseDir string
}

func (v *PathValidator) Name() string {
	return "path"
}

func (v *PathValidator) Validate(ctx *ValidatorContext, data datas.DType, rule string) error {
	// 目前仅支持字符串类型的路径校验
	dStr, ok := data.(*datas.DString)
	if !ok {
		return nil
	}

	path := strings.TrimSpace(dStr.Value)
	if path == "" {
		return nil // 允许为空
	}

	// 拼接基础路径
	fullPath := path
	if v.BaseDir != "" && !filepath.IsAbs(path) {
		fullPath = filepath.Join(v.BaseDir, path)
	}

	// 如果 rule 指定了后缀名要求 (如 path="*.png")，暂时简单处理
	if strings.HasPrefix(rule, "*.") {
		ext := strings.TrimPrefix(rule, "*")
		if !strings.HasSuffix(strings.ToLower(path), strings.ToLower(ext)) {
			return fmt.Errorf("path %s does not match extension %s", path, ext)
		}
	}

	// 检查文件是否存在
	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		return fmt.Errorf("path validation failed: file %s not found (resolved to: %s)", path, fullPath)
	}

	return nil
}
