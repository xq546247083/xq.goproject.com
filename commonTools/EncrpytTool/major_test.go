package EncrpytTool

import (
	"testing"
)

func TestEncryptString(t *testing.T) {
	s := ""
	result := Encrypt(s)
	if result != "6fda14112d9151ebefc40a96c9b85be3" {
		t.Errorf("Md5String(\"hello world\") failed.Got %s, expected %s", result, "6fda14112d9151ebefc40a96c9b85be3")
	}
}
