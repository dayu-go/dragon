package zerolog

import (
	"os"
	"testing"

	"github.com/dayu-go/gkit/log"
)

// go test -v *.go -test.run=TestLogger
func TestLogger(t *testing.T) {
	l := NewLogger(WithOutput(os.Stderr), WithProductionMode(), WithLevel(log.LevelInfo))
	l.Log(log.LevelInfo, "k1", "v1", "k2", "v2")
	l.Log(log.LevelDebug, "k1", "v1", "k2", "v2")

	l = NewLogger(WithOutput(os.Stderr), WithDevelopmentMode())
	l.Log(log.LevelInfo, "k1", "v1", "k2", "v2")
	l.Log(log.LevelDebug, "k1", "v1", "k2", "v2")

	h := log.NewHelper(l)
	h.WithFields(map[string]interface{}{"kk1": "vv1"}).Infof("this is a msg")
	h.WithFields(map[string]interface{}{"kk1": "vv1"}).Infof("this is a %s msg", "format")
	h.WithFields(map[string]interface{}{"kk1": "vv1"}).Infow("a1", "b1", "a2", "b2")
}
