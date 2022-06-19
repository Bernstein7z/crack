package chapterOne

import (
	"strings"
)

// IsUnique determines if a string has all unique characters.
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

// CheckPermutation decides if one of the two given strings is a permutation of the other.
func CheckPermutation(str1, str2 string) bool {
	// Check if both strings have the same length or equal
	if len(str1) != len(str2) || str1 == str2 {
		return false
	}

	// Check if both strings have the same char set
	len1 := len(str1)
	var sumStr1 uint8 = 0
	var sumStr2 uint8 = 0
	for i := 0; i < len1; i++ {
		sumStr1 += str1[i]
		sumStr2 += str2[i]
	}
	if sumStr1 != sumStr2 {
		return false
	}

	return true
}

// URLify replaces all spaces in a string with %20.
func URLify(rawURL string) string {
	return strings.ReplaceAll(strings.Trim(rawURL, " "), " ", "%20")
}
