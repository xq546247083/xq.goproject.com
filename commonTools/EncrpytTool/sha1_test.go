package EncrpytTool

import (
	"testing"
)

func TestSha1String(t *testing.T) {
	s := "hello world"
	result := Sha1String(s)
	if result != "2aae6c35c94fcfb415dbe95f408b9ce91ee846ed" {
		t.Errorf("Sha1String(\"hello world\") failed.Got %s, expected %s", result, "2aae6c35c94fcfb415dbe95f408b9ce91ee846ed")
	}
}

func TestSha1Bytes(t *testing.T) {
	s := "hello world"
	b := []byte(s)
	result := Sha1Bytes(b)
	if result != "2aae6c35c94fcfb415dbe95f408b9ce91ee846ed" {
		t.Errorf("Sha1Bytes(\"hello world\") failed.Got %s, expected %s", result, "2aae6c35c94fcfb415dbe95f408b9ce91ee846ed")
	}
}
