package logger

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"strings"
)

// defaultConfig 默认配置
var defaultConfig = LogConfig{
	Level:             "debug",
	Development:       true,
	DisableConsole:    false,
	DisableCaller:     false,
	DisableStacktrace: false,
	Encoding:          "console",
	EncoderConfig:     "development",

	FileName:   "logs/app.log",
	MaxSize:    100,
	MaxBackups: 3,
	MaxAge:     30,
	Compress:   false,
}

type LogConfig struct {
	Level             string `json:"level" mapstructure:"level"`                     // 日志级别
	Development       bool   `json:"development" mapstructure:"development"`         // zap.Config 开发模式
	DisableConsole    bool   `json:"disable_console" mapstructure:"disable_console"` // 是否禁用控制台输出
	DisableCaller     bool   `json:"disable_caller" mapstructure:"disable_caller"`
	DisableStacktrace bool   `json:"disable_stacktrace" mapstructure:"disable_stacktrace"`
	Encoding          string `json:"encoding" mapstructure:"encoding"` // 编码格式: json, console
	EncoderConfig     string `json:"encoder_config" mapstructure:"encoder_config"`

	FileName   string `json:"filename" mapstructure:"filename"`       // 文件位置, 不设置则不输出到文件
	MaxSize    int    `json:"max_size" mapstructure:"max_size"`       // megabytes，M 为单位，达到这个设置数后就进行日志切割
	MaxBackups int    `json:"max_backups" mapstructure:"max_backups"` // 保留旧文件最大份数
	MaxAge     int    `json:"max_age" mapstructure:"max_age"`         //days ， 旧文件最大保存天数
	Compress   bool   `json:"compress" mapstructure:"compress"`       // disabled by default，是否压缩日志归档，默认不压缩
}

func (c LogConfig) build() (*zap.Config, *lumberjack.Logger) {
	var level zapcore.Level
	switch strings.ToLower(c.Level) {
	case "debug":
		level = zapcore.DebugLevel
	case "info":
		level = zapcore.InfoLevel
	case "warn":
		level = zapcore.WarnLevel
	case "error":
		level = zapcore.ErrorLevel
	case "dpanic":
		level = zapcore.DPanicLevel
	case "panic":
		level = zapcore.PanicLevel
	case "fatal":
		level = zapcore.FatalLevel
	default:
		level = zapcore.InfoLevel
	}

	var enc zapcore.EncoderConfig
	switch strings.ToLower(c.EncoderConfig) {
	case "development", "dev":
		enc = zap.NewDevelopmentEncoderConfig()
	case "production", "prod":
		enc = zap.NewProductionEncoderConfig()
	default:
		enc = zap.NewDevelopmentEncoderConfig()
	}

	var outputs, errOutPuts []string
	if !c.DisableConsole {
		outputs = append(outputs, "stdout")
		errOutPuts = append(errOutPuts, "stderr")
	}

	zapConfig := &zap.Config{
		Level:             zap.NewAtomicLevelAt(level),
		Development:       c.Development,
		DisableCaller:     c.DisableCaller,
		DisableStacktrace: c.DisableStacktrace,
		Encoding:          c.Encoding,
		EncoderConfig:     enc,
		OutputPaths:       outputs,
		ErrorOutputPaths:  errOutPuts,
	}

	if len(c.FileName) > 0 {
		fileConfig := &lumberjack.Logger{
			Filename:   c.FileName,
			MaxSize:    c.MaxSize,
			MaxBackups: c.MaxBackups,
			MaxAge:     c.MaxAge,
			Compress:   c.Compress,
		}
		return zapConfig, fileConfig
	}

	return zapConfig, nil
}
