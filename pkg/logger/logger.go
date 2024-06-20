package logger

import (
	"errors"
	"github.com/natefinch/lumberjack"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var ErrMissingLevel = errors.New("missing level")

// DefaultLogger 使用默认配置创建的 Logger
func DefaultLogger() (Logger, error) {
	return InitLogger(defaultConfig)
}

// InitLoggerByFile 通过配置文件初始化 Logger
func InitLoggerByFile(configFile string) (Logger, error) {
	return InitLoggerByViper(nil, configFile)
}

// InitLoggerByViper 通过 Flag&配置文件 初始化 Logger
func InitLoggerByViper(viper *viper.Viper, configFile string) (Logger, error) {
	cfg, err := loadConfig(viper, configFile, "log")
	if err != nil {
		return nil, err
	}
	return InitLogger(cfg)
}

// InitLogger 初始化 Logger
func InitLogger(cfg LogConfig) (Logger, error) {
	zapConfig, fileConfig := cfg.build()
	enc, err := buildEncoder(zapConfig)
	if err != nil {
		return nil, err
	}
	sink, errSink, err := openSinks(zapConfig)
	if err != nil {
		return nil, err
	}
	if zapConfig.Level == (zap.AtomicLevel{}) {
		return nil, ErrMissingLevel
	}

	cores := []zapcore.Core{
		zapcore.NewCore(enc, sink, zapConfig.Level),
	}
	if fileConfig != nil {
		cores = append(cores, zapcore.NewCore(enc, zapcore.AddSync(fileConfig), zapConfig.Level))
	}

	zapLog := zap.New(zapcore.NewTee(cores...), buildOptions(zapConfig, errSink)...)

	l := &logger{
		SugaredLogger: zapLog.Sugar(),
		fileLogger:    fileConfig,
	}
	return l, nil
}

var _ Logger = (*logger)(nil)

type logger struct {
	*zap.SugaredLogger
	fileLogger *lumberjack.Logger
}

func (l logger) With(args ...any) Logger {
	return &logger{
		SugaredLogger: l.SugaredLogger.With(args...),
	}
}

func (l logger) Flush() error {
	var err error
	err = l.SugaredLogger.Sync()
	if l.fileLogger != nil {
		e := l.fileLogger.Close()
		if err == nil {
			err = e
		}
	}
	return err
}
