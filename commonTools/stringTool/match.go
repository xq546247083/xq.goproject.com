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

// IsImg 是否图片
func IsImg(val string) bool {
	pattern := ".*(.jpg|.png|.jpeg)$"

	result, err := regexp.MatchString(pattern, val)
	if err != nil {
		return false
	}

	return result
}

// IsLetter 单个字符是否字母(val必须为单个字符)
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
	pattern := "^[A-Za-z0-9]+$"

	result, err := regexp.MatchString(pattern, val)
	if err != nil {
		return false
	}

	return result
}
