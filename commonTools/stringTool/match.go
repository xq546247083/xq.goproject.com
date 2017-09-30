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
