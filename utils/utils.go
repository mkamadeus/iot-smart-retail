package utils

import (
	"reflect"
	"strings"
	"unsafe"
)

// SplitByCharacter splits string by a particular character and returns
// an array of strings
func SplitByCharacter(s string, c rune) []string {
	return strings.FieldsFunc(s, func(r rune) bool {
		return r == c
	})
}

// B2S converts bytes array to string without memory allocation
//
// Note: It may break if string and/or slice header changes in the future Go versions
func B2S(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// S2B converts string to bytes array without memory allocation
//
// Note: It may break if string and/or slice header changes in the future Go versions
func S2B(s string) (b []byte) {
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))

	bh.Len = sh.Len
	bh.Cap = sh.Len
	bh.Data = sh.Data
	return b
}
