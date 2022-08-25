package routes

import (
	"fmt"
	"strings"
)

func IntToString(value int) string {
	var str string = fmt.Sprint(value)

	return str
}

func ToLower(value string) string {
	return strings.ToLower(value)
}

// https://zetcode.com/golang/string-format/
func formatable(number int) string { return "{" + fmt.Sprint(number) + "}" }
func FormatString(str string, value []string) string {
	for i := 0; i < len(value); i++ {
		str = strings.Replace(str, formatable(i), value[i], -1)
	}

	return str
}

func ArrayIncludes(array []string, value string) bool {
	for _, val := range array {
		if val == value {
			return true
		}
	}

	return false
}

func stringAppender(original string, addition []string) string {
	var str string = original

	for i := 0; i < len(addition); i++ {
		str = str + addition[i]
	}

	return str
}