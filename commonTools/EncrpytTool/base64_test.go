package EncrpytTool

import (
	"fmt"
	"testing"
)

func TestBase64String(t *testing.T) {
	s := "hello world"
	result := Base64Encrypt([]byte(s))
	fmt.Println(fmt.Sprintf("Base64Encrypt(\"hello world\") Success.Got %s", string(result)))

	panic(fmt.Errorf("s"))
}

func TestBase64Bytes(t *testing.T) {
	s := "hello world"
	result := Base64Encrypt([]byte(s))

	result, err := Base64Decrypt(result)
	if err != nil {
		t.Errorf("Base64Decrypt(\"hello world\") failed.Got %s, expected %s", result, err)
	} else {
		fmt.Println(fmt.Sprintf("Base64Decrypt(\"hello world\") Success.Got %s", string(result)))
	}

	panic(fmt.Errorf("s"))
}
