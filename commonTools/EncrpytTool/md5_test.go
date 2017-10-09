package EncrpytTool

import (
	"testing"
)

var (
	ExpectedLowerString = "5eb63bbbe01eeed093cb22bb8f5acdc3"
)

func TestMd5String(t *testing.T) {
	s := "hello world"
	result := Md5String(s)
	if result != ExpectedLowerString {
		t.Errorf("Md5String(\"hello world\") failed.Got %s, expected %s", result, ExpectedLowerString)
	}
}

func TestMd5Bytes(t *testing.T) {
	s := "hello world"
	b := []byte(s)
	result := Md5Bytes(b)
	if result != ExpectedLowerString {
		t.Errorf("Md5String(\"hello world\") failed.Got %s, expected %s", result, ExpectedLowerString)
	}
}
