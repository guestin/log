package log

import (
	"fmt"
	"go.uber.org/zap"
)

type Opt func(*taggedLogCore)

func DefaultColorConfig() ColorConfig {
	return defaultColorCfg
}

var defaultColorCfg = ColorConfig{
	Debug: Magenta,
	Info:  Blue,
	Warn:  Yellow,
	Error: Red,
	Fatal: Red,
	Panic: Red,
}

type taggedLogCore struct {
	logger    *zap.Logger
	tagf      func() RichText
	afterTagf func(string) string
	colorCfg  ColorConfig
}

func (this *taggedLogCore) applyOpts(opt ...Opt) {
	for _, op := range opt {
		op(this)
	}
}

func (this *taggedLogCore) clone() *taggedLogCore {
	cloned := *this
	return &cloned
}

func defaultAfterTag(tag string) string {
	return fmt.Sprintf("[%s]  ", tag)
}

func defaultTagFormatOption(tag string) func() RichText {
	return func() RichText {
		return NewCustomRichText(func(c Color, bold bool) string {
			return fmt.Sprintf("%s", RichString(tag).Format(c, bold))
		})
	}
}
