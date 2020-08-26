package log

import "go.uber.org/zap"

type taggedZapLogger struct {
	taggedLogCore
}

func (this *taggedZapLogger) Debug(msg string, fields ...zap.Field) {
	panic("implement me")
}

func (this *taggedZapLogger) Info(msg string, fields ...zap.Field) {
	panic("implement me")
}

func (this *taggedZapLogger) Warn(msg string, fields ...zap.Field) {
	panic("implement me")
}

func (this *taggedZapLogger) Error(msg string, fields ...zap.Field) {
	panic("implement me")
}

func (this *taggedZapLogger) Fatal(msg string, fields ...zap.Field) {
	panic("implement me")
}

func (this *taggedZapLogger) Panic(msg string, fields ...zap.Field) {
	msgWithTag := msg
	this.logger.Panic(msgWithTag, fields...)
}

func (this *taggedZapLogger) With(opt ...Opt) ZapLog {
	cloned := this.clone()
	cloned.applyOpts(opt...)
	return &cloned
}

func (this *taggedZapLogger) clone() taggedZapLogger {
	return *this
}

func NewTaggedZapLogger(zapLog *zap.Logger, tag string, opts ...Opt) ZapLog {
	newLogger := zapLog.WithOptions(zap.AddCallerSkip(1))
	l := &taggedLogCore{
		logger:   newLogger,
		tagf:     nil,
		colorCfg: defaultColorCfg,
	}
	l.applyOpts(opts...)
	return &taggedZapLogger{
		*l,
	}
}
