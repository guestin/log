package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"strings"
)

const (
	JsonEncoder EncodeType = iota + 0xF0
	ConsoleEncoder
)

var gZapLogger *zap.Logger
var gSugarLogger *zap.SugaredLogger

// flush logs to underlying device
func Flush() {
	if gZapLogger != nil {
		_ = gZapLogger.Sync()
	}
}

func Classic() *zap.SugaredLogger {
	return gSugarLogger
}
func Zap() *zap.Logger {
	return gZapLogger
}

type encoderFnType func(zapcore.EncoderConfig) zapcore.Encoder

type EncodeType uint8

func newZapCore(encType EncodeType, level zapcore.Level) zapcore.Core {
	var encoderFn encoderFnType = nil
	switch encType {
	case JsonEncoder:
		encoderFn = zapcore.NewJSONEncoder
	case ConsoleEncoder:
		encoderFn = zapcore.NewConsoleEncoder
	default:
		panic("invalid encoder type")
	}
	wd, _ := os.Getwd()
	skipLen := len(wd) + 1
	encoderConfig := zapcore.EncoderConfig{
		MessageKey:     "msg",
		LevelKey:       "lv",
		TimeKey:        "tm",
		NameKey:        "who",
		CallerKey:      "caller",
		StacktraceKey:  "trace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalColorLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.MillisDurationEncoder,
		EncodeCaller: func(caller zapcore.EntryCaller, encoder zapcore.PrimitiveArrayEncoder) {
			fullPath := caller.FullPath()
			if strings.HasPrefix(wd, fullPath) {
				encoder.AppendString(fullPath[skipLen:])
			} else {
				encoder.AppendString(caller.TrimmedPath())
			}
		},
	}
	return zapcore.NewCore(
		encoderFn(encoderConfig),
		zapcore.AddSync(os.Stdout),
		zap.NewAtomicLevelAt(level))
}

func init() {
	core := newZapCore(ConsoleEncoder, zapcore.DebugLevel)
	gZapLogger = zap.New(core,
		zap.AddCaller(),
		zap.AddStacktrace(zapcore.ErrorLevel),
		zap.ErrorOutput(zapcore.AddSync(os.Stderr)))
	gSugarLogger = gZapLogger.Sugar()
}
