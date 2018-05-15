package logToolTest

import (
	"testing"

	"xq.goproject.com/commonTools/logTool"
)

func TestLog(t *testing.T) {
	for i := 1; i <= 1000; i++ {
		logTool.LogDebug("a test log message")
	}
}
