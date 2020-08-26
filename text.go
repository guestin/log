package log

import "fmt"

type RichText interface {
	Format(c Color, bold bool) string
}

type customRichText struct {
	implFunc func(c Color, bold bool) string
}

func NewCustomRichText(implFunc func(c Color, bold bool) string) RichText {
	return &customRichText{
		implFunc: implFunc,
	}
}

func (this *customRichText) Format(c Color, bold bool) string {
	return this.implFunc(c, bold)
}

type RichString string

func (this RichString) Format(c Color, bold bool) string {
	const normalFormat = "\x1b[%dm%s\x1b[0m"
	const boldFormat = "\x1b[1;%dm%s\x1b[0m"
	var curFmt string
	if !bold {
		curFmt = normalFormat
	} else {
		curFmt = boldFormat
	}
	return fmt.Sprintf(curFmt, uint8(c), this)
}

func (this RichString) logTagFormat(c Color, sp string) string {
	return fmt.Sprintf("[%s]%s", this.Format(c, false), sp)
}

type ColorConfig struct {
	Debug    Color
	Info     Color
	Warn     Color
	Error    Color
	Fatal    Color
	Panic    Color
	_padding [2]Color
}
