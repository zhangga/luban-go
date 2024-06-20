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
func loadConfig(vip *viper.Viper, filePath string, prefix string) (LogConfig, error) {
	if vip == nil {
		vip = viper.New()
	}
	// 从配置文件读取
	if err := loadConfigByFile(filePath, vip); err != nil {
		return LogConfig{}, err
	}

	var err error
	config := defaultConfig
	if len(prefix) > 0 {
		err = vip.Sub(prefix).UnmarshalExact(&config)
	} else {
		err = vip.UnmarshalExact(&config)
	}
	return config, err
}

// loadConfigByFile 从配置文件中读取
func loadConfigByFile(filePath string, vip *viper.Viper) error {
	// 未指定配置文件
	if len(filePath) == 0 {
		return nil
	}

	_, err := os.Stat(filePath)
	// 文件存在
	if err == nil || os.IsExist(err) {
		vip.SetConfigFile(filePath)
		return vip.ReadInConfig()
	}
	return err
}
