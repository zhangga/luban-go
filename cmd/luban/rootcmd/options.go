package rootcmd

import (
	"errors"
	"fmt"
	"os"
)

// CommandOptions 启动参数
type CommandOptions struct {
	SchemaCollector       string   `json:"schema_collector" mapstructure:"schema_collector"`                 // schema collector name
	Conf                  string   `json:"conf" mapstructure:"conf"`                                         // luban conf file path
	Target                string   `json:"target" mapstructure:"target"`                                     // target name
	CodeTargets           []string `json:"code_targets" mapstructure:"code_targets"`                         // code target name list
	DataTargets           []string `json:"data_targets" mapstructure:"data_targets"`                         // data target name list
	Pipeline              string   `json:"pipeline" mapstructure:"pipeline"`                                 // pipeline name
	ForceLoadTableDatas   bool     `json:"force_load_table_datas" mapstructure:"force_load_table_datas"`     // force load table datas when not any dataTarget
	IncludeTags           []string `json:"include_tags" mapstructure:"include_tags"`                         // include tags
	ExcludeTags           []string `json:"exclude_tags" mapstructure:"exclude_tags"`                         // exclude tags
	OutputTables          []string `json:"output_tables" mapstructure:"output_tables"`                       // output tables
	TimeZone              string   `json:"time_zone" mapstructure:"time_zone"`                               // time zone
	CustomTemplateDirs    []string `json:"custom_template_dirs" mapstructure:"custom_template_dirs"`         // custom template dirs
	ValidationFailAsError bool     `json:"validation_fail_as_error" mapstructure:"validation_fail_as_error"` // validation fail as error
	Xargs                 []string `json:"xargs" mapstructure:"xargs"`                                       // args like -x a=1 -x b=2
	Verbose               bool     `json:"verbose" mapstructure:"verbose"`                                   // verbose
	LogConfig             string   `json:"log_config" mapstructure:"log_config"`                             // log config file
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
	if err = rootViper.Unmarshal(options); err != nil {
		return err
	}

	// 检查必要的参数是否设置
	if len(options.Conf) == 0 {
		return errors.New("must set boostrap args, can use --conf=$luban_conf_file_path")
	}
	if len(options.Target) == 0 {
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
