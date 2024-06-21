package rootcmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/zhangga/luban/core"
	"github.com/zhangga/luban/core/options"
	"github.com/zhangga/luban/pkg/logger"
	"github.com/zhangga/luban/pkg/version"
)

var RootCmd = &cobra.Command{
	Use:     "luban",
	Short:   "luban game config tool",
	Version: version.Version,
	Run:     run,
}

// 初始化cobra命令
func init() {
	RootCmd.AddCommand(version.Command())
}

var (
	// rootViper 使用viper读取配置
	rootViper = viper.New()

	// configPath 启动配置文件
	configPath string
	// opts 启动配置
	opts options.CommandOptions
)

// 绑定启动参数
func init() {
	// 启动参数配置文件
	RootCmd.Flags().StringVar(&configPath, FlagNameConfigPath, "configs/config.full.toml", "luban boostrap config file path")

	// 命令行参数绑定
	RootCmd.Flags().StringVarP(&opts.SchemaCollector, "schema_collector", "s", "default", "schema collector name")
	RootCmd.Flags().StringVar(&opts.Conf, "conf", "", "luban.conf file path")
	RootCmd.Flags().StringVarP(&opts.Target, "target", "t", "", "target name")
	RootCmd.Flags().StringSliceVarP(&opts.CodeTargets, "code_targets", "c", nil, "code target name list")
	RootCmd.Flags().StringSliceVarP(&opts.DataTargets, "data_targets", "d", nil, "data target name list")
	RootCmd.Flags().StringVarP(&opts.Pipeline, "pipeline", "p", "default", "pipeline name")
	RootCmd.Flags().BoolVarP(&opts.ForceLoadTableDatas, "force_load_table_datas", "f", false, "force load table datas when not any dataTarget")
	RootCmd.Flags().StringSliceVarP(&opts.IncludeTags, "include_tags", "i", nil, "include tags")
	RootCmd.Flags().StringSliceVarP(&opts.ExcludeTags, "exclude_tags", "e", nil, "exclude tags")
	RootCmd.Flags().StringSliceVarP(&opts.OutputTables, "output_tables", "o", nil, "output tables")
	RootCmd.Flags().StringVar(&opts.TimeZone, "time_zone", "", "time zone")
	RootCmd.Flags().StringSliceVar(&opts.CustomTemplateDirs, "custom_template_dirs", nil, "custom template dirs")
	RootCmd.Flags().BoolVar(&opts.ValidationFailAsError, "validation_fail_as_error", false, "validation fail as error")
	RootCmd.Flags().StringSliceVarP(&opts.Xargs, "xargs", "x", nil, "args like -x a=1 -x b=2")
	RootCmd.Flags().BoolVarP(&opts.Verbose, "verbose", "v", false, "verbose")

	// 日志参数
	RootCmd.Flags().StringVarP(&opts.LogConfig, "log_config", "l", "", "log config file. [Log].xxx")
}

// 绑定命令行参数
func init() {
	// 启用环境变量
	rootViper.AutomaticEnv()
	// 绑定命令行参数
	if err := rootViper.BindPFlags(RootCmd.Flags()); err != nil {
		panic(err)
	}
}

func run(cmd *cobra.Command, args []string) {
	// 解析启动参数
	if err := loadCommandOptions(configPath, &opts); err != nil {
		panic(err)
	}

	// 初始化日志
	log, err := logger.InitLoggerByViper(rootViper, opts.LogConfig)
	if err != nil {
		panic(err)
	}
	defer log.Flush()

	log.Info(copyright)
	log.Infof("boostrap command options: %+v", opts)

	// 启动
	launcher := core.NewSimpleLauncher(log)
	launcher.Start(opts)
}
