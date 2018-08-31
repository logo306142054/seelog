package seelog_test

import (
	"testing"
	"github.com/xmge/seelog"
	"time"
)

func TestDebug(t *testing.T) {
	seelog.SetPort(8080)
	for {
		time.Sleep(1 * time.Second)
		seelog.Debug(time.Now().UTC())
	}
}