package log

import (
	"github.com/hashicorp/go-hclog"
	"testing"
)

func testA(l hclog.InterceptLogger) {
	l.Info("testA info")
}

func testB(l hclog.InterceptLogger) {
	l.Info("testB info")
}

func TestLogger(t *testing.T) {
	level := hclog.Level(1)
	l := NewLogger(".", "test", level, true, NsqConfig{})
	l.Error("error")
	l.Info("info")
	testA(l)
	testB(l)
}
