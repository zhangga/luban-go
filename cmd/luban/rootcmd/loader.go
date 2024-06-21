package rootcmd

import (
	"errors"
	"fmt"
	"github.com/zhangga/luban/core/options"
	"os"
)

const FlagNameConfigPath = "config"

// loadCommandOptions 读取启动参数
func loadCommandOptions(filePath string, opts *options.CommandOptions) error {
	var err error
	// 从配置文件读取
	if err = loadOptionsByFile(filePath); err != nil {
		return err
	}

	// 反序列化配置
	if err = rootViper.Unmarshal(opts); err != nil {
		return err
	}

	// 检查必要的参数是否设置
	if len(opts.Conf) == 0 {
		return errors.New("must set boostrap args, can use --conf=$luban_conf_file_path")
	}
	if len(opts.Target) == 0 {
		return errors.New("must set boostrap args, can use -t=$target_name")
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
		rootViper.SetConfigFile(filePath)
		return rootViper.ReadInConfig()
	}

	// 2. 手动指定了配置文件路径，但是不存在
	if rootViper.IsSet(FlagNameConfigPath) {
		return fmt.Errorf("cannot found config file: --%s=%s", FlagNameConfigPath, filePath)
	}
	// 3. 没指定配置文件
	return nil
}
