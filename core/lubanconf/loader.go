package lubanconf

import (
	"errors"
	"github.com/bytedance/sonic"
	json "github.com/json-iterator/go"
	"github.com/zhangga/luban/pkg/logger"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

var ErrNotDir = errors.New("not dir")

type IConfigLoader interface {
	Load(filePath string) (*LubanConfig, error)
}

var _ IConfigLoader = (*GlobalConfigLoader)(nil)

type GlobalConfigLoader struct {
	logger logger.Logger
}

func NewGlobalConfigLoader(logger logger.Logger) *GlobalConfigLoader {
	return &GlobalConfigLoader{
		logger: logger,
	}
}

func (l *GlobalConfigLoader) Load(filePath string) (*LubanConfig, error) {
	conf, err := loadJson(filePath)
	if err != nil {
		return nil, err
	}

	curDir, err := filepath.Abs(filepath.Dir(filePath))
	if err != nil {
		return nil, err
	}

	configFileName := filepath.Base(filePath)
	dataInputDir := filepath.Join(curDir, conf.DataDir)
	importFiles := make([]SchemaFile, 0, len(conf.SchemaFiles))
	for _, schemaFile := range conf.SchemaFiles {
		fullPath := filepath.Join(curDir, schemaFile.FileName)
		info, err := os.Stat(fullPath)
		if os.IsNotExist(err) {
			l.logger.Errorf("%s schema文件错误: 文件不存在 %s", configFileName, fullPath)
			return nil, err
		}
		if err != nil {
			l.logger.Errorf("%s schema文件错误: %s, err: %s", configFileName, fullPath, err)
			return nil, err
		}

		if len(schemaFile.Type) == 0 && !info.IsDir() { // type为空，必须是文件夹
			l.logger.Errorf("%s schema文件错误: %s 不是文件夹", configFileName, fullPath)
			return nil, ErrNotDir
		}

		// 指定的是文件，直接添加，不需要判断文件名
		if !info.IsDir() {
			importFiles = append(importFiles, SchemaFile{FileName: fullPath, Type: schemaFile.Type})
			continue
		}

		// 遍历文件夹下的所有文件
		filepath.Walk(fullPath, func(path string, info fs.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() {
				return nil
			}
			if strings.HasPrefix(info.Name(), ".") || strings.HasPrefix(info.Name(), "~") {
				return nil
			}
			if strings.HasPrefix(info.Name(), "_") && !strings.HasPrefix(info.Name(), "__") {
				return nil
			}
			if strings.HasSuffix(info.Name(), ".meta") {
				return nil
			}
			importFiles = append(importFiles, SchemaFile{FileName: path, Type: schemaFile.Type})
			return nil
		})
	}

	config := LubanConfig{
		ConfigFileName: configFileName,
		InputDataDir:   dataInputDir,
		Groups:         conf.Groups,
		Targets:        conf.Targets,
		Imports:        importFiles,
	}
	return &config, nil
}

func loadJson(filePath string) (*Conf, error) {
	buf, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var config Conf
	if err = sonic.Unmarshal(buf, &config); err != nil {
		return nil, err
	}
	return &config, nil
}

func loadByFile(filePath string) (*Conf, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	var conf Conf
	if err = decoder.Decode(&conf); err != nil {
		return nil, err
	}

	return &conf, nil
}
