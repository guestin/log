package log

import (
	"fmt"
	"go.uber.org/zap"
)

type taggedClassicLogger struct {
	taggedLogCore
	sugar *zap.SugaredLogger
}

func (this *taggedClassicLogger) Debug(args ...interface{}) {
	tag := this.makeTag(this.colorCfg.Debug, false)
	this.sugar.Debug(append([]interface{}{tag}, args...)...)
}

func (this *taggedClassicLogger) Debugf(template string, args ...interface{}) {
	tag := this.makeTag(this.colorCfg.Debug, false)
	this.sugar.Debugf(fmt.Sprintf(tag+"%s", template), args...)
}

func (this *taggedClassicLogger) Info(args ...interface{}) {
	tag := this.makeTag(this.colorCfg.Info, false)
	this.sugar.Info(append([]interface{}{tag}, args...)...)
}

func (this *taggedClassicLogger) Infof(template string, args ...interface{}) {
	tag := this.makeTag(this.colorCfg.Info, false)
	this.sugar.Infof(fmt.Sprintf(tag+"%s", template), args...)
}

func (this *taggedClassicLogger) Warn(args ...interface{}) {
	tag := this.makeTag(this.colorCfg.Warn, false)
	this.sugar.Warn(append([]interface{}{tag}, args...)...)
}

func (this *taggedClassicLogger) Warnf(template string, args ...interface{}) {
	tag := this.makeTag(this.colorCfg.Warn, false)
	this.sugar.Warnf(fmt.Sprintf(tag+"%s", template), args...)
}

func (this *taggedClassicLogger) Error(args ...interface{}) {
	tag := this.makeTag(this.colorCfg.Error, false)
	this.sugar.Error(append([]interface{}{tag}, args...)...)
}

func (this *taggedClassicLogger) Errorf(template string, args ...interface{}) {
	tag := this.makeTag(this.colorCfg.Error, false)
	this.sugar.Errorf(fmt.Sprintf(tag+"%s", template), args...)
}

func (this *taggedClassicLogger) Fatal(args ...interface{}) {
	tag := this.makeTag(this.colorCfg.Fatal, false)
	this.sugar.Fatal(append([]interface{}{tag}, args...)...)
}

func (this *taggedClassicLogger) Fatalf(template string, args ...interface{}) {
	tag := this.makeTag(this.colorCfg.Fatal, false)
	this.sugar.Fatalf(fmt.Sprintf(tag+"%s", template), args...)
}

func (this *taggedClassicLogger) Panic(args ...interface{}) {
	tag := this.makeTag(this.colorCfg.Panic, false)
	this.sugar.Panic(append([]interface{}{tag}, args...)...)
}

func (this *taggedClassicLogger) Panicf(template string, args ...interface{}) {
	tag := this.makeTag(this.colorCfg.Panic, false)
	this.sugar.Panicf(fmt.Sprintf(tag+"%s", template), args...)
}

func (this *taggedClassicLogger) clone() *taggedClassicLogger {
	cloned := *this
	return &cloned
}

func (this *taggedClassicLogger) With(opt ...Opt) ClassicLog {
	cloned := this.clone()
	cloned.applyOpts(opt...)
	return cloned
}

func (this *taggedClassicLogger) makeTag(c Color, b bool) string {
	return this.afterTagf(this.tagf().Format(c, b))
}

func NewTaggedClassicLogger(zapLog *zap.Logger, tag string, opts ...Opt) ClassicLog {
	newLogger := zapLog.WithOptions(zap.AddCallerSkip(1))
	out := &taggedClassicLogger{
		taggedLogCore: taggedLogCore{
			logger:    newLogger,
			tagf:      defaultTagFormatOption(tag),
			afterTagf: defaultAfterTag,
			colorCfg:  defaultColorCfg,
		},
		sugar: newLogger.Sugar(),
	}
	out.applyOpts(opts...)
	return out
}

func NewFixStyleText(tag string, c Color, b bool) RichText {
	return NewCustomRichText(func(_ Color, _ bool) string {
		return RichString(tag).Format(c, b)
	})
}
