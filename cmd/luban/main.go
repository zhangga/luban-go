package main

import (
	"github.com/spf13/viper"
	"log"

	"github.com/spf13/cobra"
	"github.com/zhangga/luban/pkg/version"
)

var rootCmd = &cobra.Command{
	Use:     "luban",
	Short:   "luban game config tool",
	Version: version.Version,
	Run:     run,
}

// 初始化cobra命令
func init() {
	rootCmd.AddCommand(version.Command())
}

var (
	// configPath 启动配置文件
	configPath string
	// options 启动配置
	options CommandOptions
)

// 绑定启动参数
func init() {
	rootCmd.Flags().StringVar(&configPath, FlagNameConfigPath, "configs/config.toml", "luban command file path")

	// 启动参数
	rootCmd.Flags().StringVar(&options.Conf, "conf", "", "luban.conf file path")

	// 日志参数
	rootCmd.Flags().StringVar(&options.SchemaCollector, "schema_collector", "", "schema collector")
}

// 绑定命令行参数
func init() {
	// 启用环境变量
	viper.AutomaticEnv()
	// 绑定命令行参数
	if err := viper.BindPFlags(rootCmd.Flags()); err != nil {
		panic(err)
	}
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
