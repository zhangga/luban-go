package lubanconf

import (
	json "github.com/json-iterator/go"
	"github.com/zhangga/luban/pkg/logger"
	"os"
	"path/filepath"
)

type IConfigLoader interface {
	Load(filePath string) (*Conf, error)
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

func (l *GlobalConfigLoader) Load(filePath string) (*Conf, error) {
	l.logger.Debugf("load luban config file: %s", filePath)
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	var config Conf
	if err = decoder.Decode(&config); err != nil {
		return nil, err
	}

	curDir, err := filepath.Abs(filepath.Dir(filePath))
	if err != nil {
		return nil, err
	}
	_ = curDir

	return &config, nil
}
