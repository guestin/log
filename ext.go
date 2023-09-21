package log

import (
	"go.uber.org/zap"
)

type ClassicLog interface {
	Debug(args ...interface{})
	Debugf(template string, args ...interface{})

	Info(args ...interface{})
	Infof(template string, args ...interface{})

	Warn(args ...interface{})
	Warnf(template string, args ...interface{})

	Error(args ...interface{})
	Errorf(template string, args ...interface{})

	Fatal(args ...interface{})
	Fatalf(template string, args ...interface{})

	Panic(args ...interface{})
	Panicf(template string, args ...interface{})

	With(opt ...Opt) ClassicLog
}

type ZapLog interface {
	Debug(msg string, fields ...zap.Field)
	Info(msg string, fields ...zap.Field)
	Warn(msg string, fields ...zap.Field)
	Error(msg string, fields ...zap.Field)
	Fatal(msg string, fields ...zap.Field)
	Panic(msg string, fields ...zap.Field)
	With(opt ...Opt) ZapLog
}

// Color thanks for go.uber.org/zap/internal/color.go
type Color uint8

const (
	Black Color = iota + 30
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	White
)
