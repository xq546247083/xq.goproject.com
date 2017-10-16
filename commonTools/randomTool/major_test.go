package randomTool

import (
	"fmt"
	"strings"
	"testing"
)

func TestGetRandomStr(t *testing.T) {
	for {
		str1 := GetRandomStr(6)
		if strings.Index(str1, "Z") == 5 {
			fmt.Println(str1)
			break
		}
	}

	t.Errorf(GetRandomStr(6))
}
