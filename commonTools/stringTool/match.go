package stringTool

import "regexp"

// IsEmail 是否邮箱
func IsEmail(val string) bool {
	pattern := "^[\\w-]+(\\.[\\w-]+)*@[\\w-]+(\\.[\\w-]+)+$"

	result, err := regexp.MatchString(pattern, val)
	if err != nil {
		return false
	}

	return result
}

// IsLetter 单个字符是否字母
func IsLetter(val string) bool {
	pattern := "[a-zA-Z]"

	result, err := regexp.MatchString(pattern, val)
	if err != nil {
		return false
	}

	return result
}

// IsLetterOrDigit 是否字母或者数字构成的字符串
func IsLetterOrDigit(val string) bool {
	pattern := "/^[0-9a-zA-Z]*$/g"

	result, err := regexp.MatchString(pattern, val)
	if err != nil {
		return false
	}

	return result
}
