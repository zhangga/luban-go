package logger

import (
	"errors"
	"github.com/spf13/viper"
	"os"
)

var (
	errFilePathIsEmpty = errors.New("config filepath is empty")
)

// loadConfig 读取配置
func loadConfig[T any](filePath string, config *T) error {
	var (
		err   error
		viper = viper.New()
	)
	// 从配置文件读取
	if err = loadConfigByFile(filePath, viper); err != nil {
		return err
	}

	// 反序列化配置
	if err = viper.Unmarshal(config); err != nil {
		return err
	}

	return viper.Sub("log").UnmarshalExact(config)
}

// loadConfigByFile 从配置文件中读取
func loadConfigByFile(filePath string, viper *viper.Viper) error {
	// 未指定配置文件
	if len(filePath) == 0 {
		return errFilePathIsEmpty
	}

	_, err := os.Stat(filePath)
	// 文件存在
	if err == nil || os.IsExist(err) {
		viper.SetConfigFile(filePath)
		return viper.ReadInConfig()
	}
	return err
}
