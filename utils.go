package main

import (
	"unicode"
)

func capitalize(s string) string {
	a := []rune(s)
	a[0] = unicode.ToUpper(a[0])
	return string(a)
}
