package randomTool

import (
	"testing"
)

func TestGetRandomStr(t *testing.T) {
	t.Errorf(GetRandomStr(6))
}
