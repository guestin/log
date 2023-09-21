package log

import (
	"go.uber.org/zap/zapcore"
	"testing"
)

var rlogger, _ = EasyInitConsoleLogger(zapcore.DebugLevel, zapcore.ErrorLevel)

func TestTaggedClassic(t *testing.T) {
	l := NewTaggedClassicLogger(rlogger, "test")
	l.With(UseColor(Blue),
		UseSubTag(NewFixStyleText("red!", Red, false)),
		UseSubTag(NewFixStyleText("green!", Green, true))).
		Debug("aabbcc")
	l.With(UseTag("oooo")).Debugf("string is=%s", "aabbcc")
}
