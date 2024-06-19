package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	_ "unsafe"
)

//go:linkname buildEncoder go.uber.org/zap.(*Config).buildEncoder
func buildEncoder(cfg *zap.Config) (zapcore.Encoder, error)

//go:linkname openSinks go.uber.org/zap.(*Config).openSinks
func openSinks(cfg *zap.Config) (zapcore.WriteSyncer, zapcore.WriteSyncer, error)

//go:linkname buildOptions go.uber.org/zap.(*Config).buildOptions
func buildOptions(cfg *zap.Config, errSink zapcore.WriteSyncer) []zap.Option
