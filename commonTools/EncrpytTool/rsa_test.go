package EncrpytTool

import (
	"fmt"
	"testing"
)

func TestRsaString(t *testing.T) {
	s := "hello world"
	result, err := RsaEncrypt([]byte(s))
	if err != nil {
		t.Errorf("RsaEncrypt(\"hello world\") failed.Got %s, expected %s", result, err)
	} else {
		fmt.Println(fmt.Sprintf("RsaEncrypt(\"hello world\") Success.Got %s", string(result)))
	}

	panic(fmt.Errorf("s"))
}

func TestRsaBytes(t *testing.T) {
	s := "hello world"
	result, _ := RsaEncrypt([]byte(s))

	result, err := RsaDecrypt(result)
	if err != nil {
		t.Errorf("RsaDecrypt(\"hello world\") failed.Got %s, expected %s", result, err)
	} else {
		fmt.Println(fmt.Sprintf("RsaDecrypt(\"hello world\") Success.Got %s", string(result)))
	}

	panic(fmt.Errorf("s"))
}
