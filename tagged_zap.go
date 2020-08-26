package log

import (
	"fmt"
	"go.uber.org/zap"
)

type taggedZapLogger struct {
	taggedLogCore
}

func (this *taggedZapLogger) Debug(msg string, fields ...zap.Field) {
	tag := this.makeTag(this.colorCfg.Debug, false)
	this.logger.Debug(fmt.Sprintf("%s%s", tag, msg), fields...)
}

func (this *taggedZapLogger) Info(msg string, fields ...zap.Field) {
	tag := this.makeTag(this.colorCfg.Debug, false)
	this.logger.Info(fmt.Sprintf("%s%s", tag, msg), fields...)
}

func (this *taggedZapLogger) Warn(msg string, fields ...zap.Field) {
	tag := this.makeTag(this.colorCfg.Debug, false)
	this.logger.Warn(fmt.Sprintf("%s%s", tag, msg), fields...)
}

func (this *taggedZapLogger) Error(msg string, fields ...zap.Field) {
	tag := this.makeTag(this.colorCfg.Debug, false)
	this.logger.Error(fmt.Sprintf("%s%s", tag, msg), fields...)
}

func (this *taggedZapLogger) Fatal(msg string, fields ...zap.Field) {
	tag := this.makeTag(this.colorCfg.Debug, false)
	this.logger.Fatal(fmt.Sprintf("%s%s", tag, msg), fields...)
}

func (this *taggedZapLogger) Panic(msg string, fields ...zap.Field) {
	tag := this.makeTag(this.colorCfg.Debug, false)
	this.logger.Panic(fmt.Sprintf("%s%s", tag, msg), fields...)
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
	out := &taggedZapLogger{
		taggedLogCore: taggedLogCore{
			logger:    newLogger,
			tagf:      defaultTagFormatOption(tag),
			afterTagf: defaultAfterTag,
			colorCfg:  defaultColorCfg,
		},
	}
	out.applyOpts(opts...)
	return out
}
