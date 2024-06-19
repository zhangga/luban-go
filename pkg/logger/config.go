package logger

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
)

// defaultConfig 默认配置
var defaultConfig = &Config{
	zapConfig: &zap.Config{
		Level:             zap.NewAtomicLevelAt(zap.DebugLevel),
		Development:       true,
		DisableCaller:     false,
		DisableStacktrace: false,
		Encoding:          "console",
		EncoderConfig:     zap.NewDevelopmentEncoderConfig(),
		OutputPaths:       []string{"stdout"},
		ErrorOutputPaths:  []string{"stderr"},
	},
	fileConfig: &lumberjack.Logger{
		Filename:   "logs/app.log", // 文件位置
		MaxSize:    100,            // megabytes，M 为单位，达到这个设置数后就进行日志切割
		MaxBackups: 3,              // 保留旧文件最大份数
		MaxAge:     30,             //days ， 旧文件最大保存天数
		Compress:   false,          // disabled by default，是否压缩日志归档，默认不压缩
	},
}

type Config struct {
	zapConfig  *zap.Config
	fileConfig *lumberjack.Logger
}
