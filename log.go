package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"strings"
	"sync"
)

const (
	JsonEncoder EncodeType = iota + 0xF0
	ConsoleEncoder
)

var gRootZapLogger *zap.Logger
var gRootSugarLogger *zap.SugaredLogger
var gInitOnce sync.Once

// flush logs to underlying device
func Flush() {
	if gRootZapLogger != nil {
		_ = gRootZapLogger.Sync()
	}
}

func Classic() *zap.SugaredLogger {
	return gRootSugarLogger
}
func Zap() *zap.Logger {
	return gRootZapLogger
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
			if strings.HasPrefix(fullPath, wd) {
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

// init logger with some helpful default options.
// usually used in docker container
func EasyInitConsoleLogger(logLevel zapcore.Level, stacktraceLevel zapcore.Level, options ...zap.Option) {
	options = append([]zap.Option{
		zap.AddCaller(),
		zap.AddStacktrace(stacktraceLevel),
		zap.ErrorOutput(zapcore.AddSync(os.Stderr))}, options...)
	InitLog(ConsoleEncoder, logLevel, options...)
}

// warning: if you doesn't understand what 'the option' means , use 'EasyInitConsoleLogger' instead
func InitLog(encoder EncodeType, logLevel zapcore.Level, options ...zap.Option) {
	gInitOnce.Do(func() {
		core := newZapCore(encoder, logLevel)
		gRootZapLogger = zap.New(core, options...)
		gRootSugarLogger = gRootZapLogger.Sugar()
	})
}
