package chapterOne

import (
	"strings"
)

func IsUnique(input string) bool {
	length := len(input)
	var sb strings.Builder

	for i := 0; i < length; i++ {
		sb.WriteByte(input[i])
		if strings.Count(input[i+1:], sb.String()) == 1 {
			return false
		}
		sb.Reset()
	}

	return true
}

func CheckPermutation(str1, str2 string) bool {
	return false
}
