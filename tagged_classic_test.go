package log

import (
	"testing"
)

func TestTaggedClassic(t *testing.T) {
	l := NewTaggedClassicLogger(Zap(), "test")
	l.With(UseColor(Blue),
		UseSubTag(NewFixStyleText("red!", Red, false)),
		UseSubTag(NewFixStyleText("green!", Green, true))).
		Debug("aabbcc")
	l.With(UseTag("oooo")).Debugf("string is=%s", "aabbcc")
}
