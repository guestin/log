package log

import (
	"go.uber.org/zap"
	"testing"
)

func TestTaggedZap(t *testing.T) {
	l := NewTaggedZapLogger(Zap(), "test_zap")
	l.With(UseColor(Red)).Debug("aabbcc")
	l.With(UseColor(Blue),
		UseSubTag(NewFixStyleText("red!", Red, false)),
		UseSubTag(NewFixStyleText("green!", Green, true))).
		Debug("aabbcc")
	l.With(UseTag("oooo")).Debug("string is=", zap.String("str", "aabbcc"))
}
