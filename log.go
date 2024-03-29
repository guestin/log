package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"os"
	"path/filepath"
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

// Flush logs to underlying device
//
//goland:noinspection ALL
func Flush() {
	if gRootZapLogger != nil {
		_ = gRootZapLogger.Sync()
	}
	if gRootSugarLogger != nil {
		_ = gRootSugarLogger.Sync()
	}
}

type encoderFnType func(zapcore.EncoderConfig) zapcore.Encoder

type EncodeType uint8

func newZapCore(encType EncodeType, level zapcore.Level, writer io.Writer) zapcore.Core {
	var encoderFn encoderFnType = nil
	switch encType {
	case JsonEncoder:
		encoderFn = zapcore.NewJSONEncoder
	case ConsoleEncoder:
		encoderFn = zapcore.NewConsoleEncoder
	default:
		panic("invalid encoder type")
	}
	pathCache := new(sync.Map)
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
			r, ok := pathCache.Load(fullPath)
			if !ok {
				pathArr := strings.Split(fullPath, string(filepath.Separator))
				if len(pathArr) > 3 {
					pathArr = pathArr[len(pathArr)-3:]
				}
				r = filepath.Join(pathArr...)
				pathCache.Store(fullPath, r)
			}
			encoder.AppendString(r.(string))
		},
	}
	return zapcore.NewCore(
		encoderFn(encoderConfig),
		zapcore.AddSync(writer),
		zap.NewAtomicLevelAt(level))
}

// EasyInitConsoleLogger init logger with some helpful default options.
// usually used in docker container
func EasyInitConsoleLogger(logLevel zapcore.Level,
	stacktraceLevel zapcore.Level,
	options ...zap.Option) (*zap.Logger, *zap.SugaredLogger) {
	options = append([]zap.Option{
		zap.AddCaller(),
		zap.AddStacktrace(stacktraceLevel),
		zap.ErrorOutput(zapcore.AddSync(os.Stderr))}, options...)
	return InitLog(ConsoleEncoder, logLevel, os.Stdout, options...)
}

// InitLog warning: if you don't understand what 'the option' means , use 'EasyInitConsoleLogger' instead
func InitLog(encoder EncodeType,
	logLevel zapcore.Level,
	writer io.Writer, options ...zap.Option) (*zap.Logger, *zap.SugaredLogger) {
	return initOnce(newZapCore(encoder, logLevel, writer), options...)
}

type MultiCfg struct {
	Encoder  EncodeType
	LogLevel zapcore.Level
	Writer   io.Writer
}

// InitMultiTargetLog init multi core logger
func InitMultiTargetLog(target []*MultiCfg, options ...zap.Option) (*zap.Logger, *zap.SugaredLogger) {
	cores := make([]zapcore.Core, 0)
	for i := range target {
		cores = append(cores, newZapCore(target[i].Encoder, target[i].LogLevel, target[i].Writer))
	}
	return initOnce(zapcore.NewTee(cores...), options...)
}

func initOnce(core zapcore.Core, options ...zap.Option) (*zap.Logger, *zap.SugaredLogger) {
	gInitOnce.Do(func() {
		gRootZapLogger = zap.New(core, options...)
		gRootSugarLogger = gRootZapLogger.Sugar()
	})
	return gRootZapLogger, gRootSugarLogger
}
