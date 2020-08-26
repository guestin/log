package log

import (
	"go.uber.org/zap"
	"testing"
)

func TestTaggedZap(t *testing.T) {
	l := NewTaggedZapLogger(Zap(), "test_zap")
	l.With(UseColor(Blue)).Debug("aabbcc")
	l.With(UseTag("oooo")).Debug("string is=%s", zap.String("str", "aabbcc"))
}
