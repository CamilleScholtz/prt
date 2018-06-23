package main

import (
	"path"
	"unicode"
)

type byBase []string

// Custom "by basename" sort.
func (s byBase) Len() int      { return len(s) }
func (s byBase) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s byBase) Less(i, j int) bool {
	if path.Base(s[i]) < path.Base(s[j]) {
		return true
	}
	if path.Base(s[i]) > path.Base(s[j]) {
		return false
	}
	return path.Base(s[i]) < path.Base(s[j])
}

func capitalize(s string) string {
	a := []rune(s)
	a[0] = unicode.ToUpper(a[0])
	return string(a)
}
