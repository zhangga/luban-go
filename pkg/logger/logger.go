package logger

import (
	"errors"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var ErrMissingLevel = errors.New("missing level")

func DefaultLogger() (Logger, error) {
	return InitLogger(defaultConfig)
}

func InitLoggerByFile(configFile string) (Logger, error) {
	var config *Config
	err := loadConfig[Config](configFile, config)
	if err != nil {
		return nil, err
	}
	return InitLogger(config)
}

func InitLogger(cfg *Config) (Logger, error) {
	if cfg == nil {
		cfg = defaultConfig
	}

	enc, err := buildEncoder(cfg.zapConfig)
	if err != nil {
		return nil, err
	}
	sink, errSink, err := openSinks(cfg.zapConfig)
	if err != nil {
		return nil, err
	}
	if cfg.zapConfig.Level == (zap.AtomicLevel{}) {
		return nil, ErrMissingLevel
	}

	cores := []zapcore.Core{
		zapcore.NewCore(enc, sink, cfg.zapConfig.Level),
	}
	var fileLog *lumberjack.Logger
	if cfg.fileConfig != nil {
		fileLog = &lumberjack.Logger{
			Filename:   cfg.fileConfig.Filename,
			MaxSize:    cfg.fileConfig.MaxSize,
			MaxBackups: cfg.fileConfig.MaxBackups,
			MaxAge:     cfg.fileConfig.MaxAge,
			Compress:   cfg.fileConfig.Compress,
		}
		cores = append(cores, zapcore.NewCore(enc, zapcore.AddSync(fileLog), cfg.zapConfig.Level))
	}

	zapLog := zap.New(zapcore.NewTee(cores...), buildOptions(cfg.zapConfig, errSink)...)

	l := &logger{
		SugaredLogger: zapLog.Sugar(),
		fileLogger:    fileLog,
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
