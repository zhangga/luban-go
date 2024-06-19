package main

import (
	"fmt"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"os"
)

type CommandOptions struct {
	SchemaCollector string `json:"schema_collector"`
	Conf            string `json:"conf"`
}

const FlagNameConfigPath = "config"

// loadCommandOptions 读取启动参数
func loadCommandOptions(filePath string, options *CommandOptions) error {
	var err error
	// 从配置文件读取
	if err = loadOptionsByFile(filePath); err != nil {
		return err
	}

	// 反序列化配置
	if err = viper.Unmarshal(options); err != nil {
		return err
	}

	return nil
}

// loadOptionsByFile 从配置文件中读取
func loadOptionsByFile(filePath string) error {
	// 未指定配置文件
	if len(filePath) == 0 {
		return nil
	}

	// 1. 读取配置文件. 如: configs/config.toml
	_, err := os.Stat(filePath)
	// 文件存在
	if err == nil || os.IsExist(err) {
		viper.SetConfigFile(filePath)
		return viper.ReadInConfig()
	}

	// 2. 手动指定了配置文件路径，但是不存在
	if flag := pflag.Lookup(FlagNameConfigPath); flag != nil && flag.Changed {
		return fmt.Errorf("cannot found config file ----->: --confpath=%s", filePath)
	}
	// 4. 没指定配置文件
	return nil
}
