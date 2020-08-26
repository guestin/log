package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
)

type writerWrapper struct {
	Level  zapcore.Level
	Fields []zapcore.Field
	logger *zap.Logger
}

func NewLoggerWrapper(zapLogger *zap.Logger, level zapcore.Level, fields ...zapcore.Field) io.Writer {
	return &writerWrapper{
		Level:  level,
		Fields: fields,
		logger: zapLogger,
	}
}

func (this *writerWrapper) Write(p []byte) (n int, err error) {
	if ce := this.logger.Check(this.Level, string(p)); ce != nil {
		ce.Write(this.Fields...)
	}
	// always return success
	return len(p), nil
}
