package stringTool

import (
	"fmt"
	"testing"
)

func TestIsLetter(t *testing.T) {
	s := "1"
	flag := IsLetter(s)
	fmt.Println(fmt.Sprintf("TestIsLetterOrDigit(\"h\") Success.Got %s", flag))

	panic(fmt.Errorf("s"))
}

func TestIsLetterOrDigit(t *testing.T) {
	s := "1hellow(orld"
	flag := IsLetterOrDigit(s)
	fmt.Println(fmt.Sprintf("TestIsLetterOrDigit(\"hello world\") Success.Got %s", flag))

	panic(fmt.Errorf("s"))
}
