package logger

type Logger interface {
	With(args ...any) Logger
	Debug(v ...any)
	Info(v ...any)
	Warn(v ...any)
	Error(v ...any)
	Fatal(v ...any)
	Panic(v ...any)
	Debugf(format string, v ...any)
	Infof(format string, v ...any)
	Warnf(format string, v ...any)
	Errorf(format string, v ...any)
	Fatalf(format string, v ...any)
	Panicf(format string, v ...any)
	Debugw(msg string, keysAndValues ...any)
	Infow(msg string, keysAndValues ...any)
	Warnw(msg string, keysAndValues ...any)
	Errorw(msg string, keysAndValues ...any)
	Fatalw(msg string, keysAndValues ...any)
	Panicw(msg string, keysAndValues ...any)
	Flush() error
}
