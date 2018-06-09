package main

import (
	"fmt"
	"os"
	"path"
	"unicode"

	"github.com/fatih/color"
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

// TODO: Use toTitle or Title.
func capitalize(s string) string {
	a := []rune(s)
	a[0] = unicode.ToUpper(a[0])
	return string(a)
}

// printe prints a string with an error character prefix.
func printe(s string) {
	color.Set(config.ErrorColor)
	fmt.Printf(config.ErrorChar)
	color.Unset()
	fmt.Fprintln(os.Stderr, s)
}

// printi prints a string with an indent character prefix.
func printi(s string) {
	color.Set(config.DarkColor)
	fmt.Printf(config.IndentChar)
	color.Unset()
	fmt.Println(s)
}

// reverseList reverses a list
/*func reverseList(l []string) []string {
	var nl []string
	for i := range l {
		nl = append(nl, l[len(l)-1-i])
	}

	return nl
}*/
