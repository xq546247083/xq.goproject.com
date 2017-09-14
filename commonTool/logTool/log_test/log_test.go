package logToolTest

import (
	"xq.goproject.com/commonTool/logTool"
	"testing"
)

func TestLog(t *testing.T) {
	for i := 1; i <= 1000; i++ {
		logTool.WriteLog("a test log message", logTool.Info)
	}
}
