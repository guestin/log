package log

import "go.uber.org/zap/zapcore"

func UseColor(color Color) Opt {
	return func(l *taggedLogCore) {
		l.colorCfg.Debug = color
		l.colorCfg.Info = color
		l.colorCfg.Warn = color
		l.colorCfg.Error = color
		l.colorCfg.Fatal = color
		l.colorCfg.Panic = color
	}
}

func UseTag(tag string) Opt {
	return func(logger *taggedLogCore) {
		logger.tagf = func() RichText {
			return RichString(tag)
		}
	}
}

func UseSubTag(subTag RichText) Opt {
	return func(core *taggedLogCore) {
		oldTag := core.tagf()
		core.tagf = func() RichText {
			return NewCustomRichText(func(c Color, bold bool) string {
				return oldTag.Format(c, bold) + "->" + subTag.Format(c, bold)
			})
		}
	}
}

func UseLevelColor(level zapcore.Level, color Color) Opt {
	return func(l *taggedLogCore) {
		var target *Color
		switch level {
		case zapcore.DebugLevel:
			target = &l.colorCfg.Debug
		case zapcore.InfoLevel:
			target = &l.colorCfg.Info
		case zapcore.WarnLevel:
			target = &l.colorCfg.Warn
		case zapcore.ErrorLevel:
			target = &l.colorCfg.Error
		case zapcore.FatalLevel:
			target = &l.colorCfg.Fatal
		case zapcore.PanicLevel:
			target = &l.colorCfg.Panic
		default:
			panic("??? unknown level")
		}
		*target = color
	}
}

func UseAfterTagRender(f func(tag string) string) Opt {
	return func(c *taggedLogCore) {
		c.afterTagf = f
	}
}
