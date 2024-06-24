package logger

import (
	"fmt"
	"sync/atomic"
)

var defLogger atomic.Pointer[Logger]

func init() {
	def, err := InitLogger(defaultConfig)
	if err != nil {
		panic(fmt.Errorf("init default logger failed: %v", err))
	}
	defLogger.Store(&def)
}

// InitDefaultLogger 手动初始化默认日志
func InitDefaultLogger(cfg LogConfig) (Logger, error) {
	def, err := InitLogger(defaultConfig)
	if err != nil {
		return nil, err
	}
	defLogger.Store(&def)
	return def, nil
}

func Default() Logger {
	return (*defLogger.Load()).(Logger)
}

func With(args ...any) Logger {
	return Default().With(args...)
}

func Debug(v ...any) {
	Default().Debug(v...)
}

func Info(v ...any) {
	Default().Info(v...)
}

func Warn(v ...any) {
	Default().Warn(v...)
}

func Error(v ...any) {
	Default().Error(v...)
}

func Fatal(v ...any) {
	Default().Fatal(v...)
}

func Panic(v ...any) {
	Default().Panic(v...)
}

func Debugf(format string, v ...any) {
	Default().Debugf(format, v...)
}

func Infof(format string, v ...interface{}) {
	Default().Infof(format, v...)
}

func Warnf(format string, v ...interface{}) {
	Default().Warnf(format, v...)
}

func Errorf(format string, v ...interface{}) {
	Default().Errorf(format, v...)
}

func Fatalf(format string, v ...interface{}) {
	Default().Fatalf(format, v...)
}

func Panicf(format string, v ...interface{}) {
	Default().Panicf(format, v...)
}

func Debugw(msg string, keysAndValues ...any) {
	Default().Debugw(msg, keysAndValues...)
}

func Infow(msg string, keysAndValues ...any) {
	Default().Infow(msg, keysAndValues...)
}

func Warnw(msg string, keysAndValues ...any) {
	Default().Warnw(msg, keysAndValues...)
}

func Errorw(msg string, keysAndValues ...any) {
	Default().Errorw(msg, keysAndValues...)
}

func Fatalw(msg string, keysAndValues ...any) {
	Default().Fatalw(msg, keysAndValues...)
}

func Panicw(msg string, keysAndValues ...any) {
	Default().Panicw(msg, keysAndValues...)
}

func Flush() error {
	return Default().Flush()
}
