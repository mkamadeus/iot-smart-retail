package utils

import "strings"

// SplitByCharacter splits string by a particular character and returns
// an array of strings
func SplitByCharacter(s string, c rune) []string {
	return strings.FieldsFunc(s, func(r rune) bool {
		return r == c
	})
}
