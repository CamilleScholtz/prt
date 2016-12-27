package utils

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/onodera-punpun/prt/config"
)

// Load config.
var c = config.Load()

// Printe prints a string with an error character prefix.
func Printe(s string) {
	color.Set(c.ErrorColor)
	fmt.Printf(c.ErrorChar)
	color.Unset()
	fmt.Fprintln(os.Stderr, s)
}

// Printi prints a string with an indent character prefix.
func Printi(s string) {
	color.Set(c.DarkColor)
	fmt.Printf(c.IndentChar)
	color.Unset()
	fmt.Println(s)
}

// StringInList checks if a string is in a list.
func StringInList(s string, l []string) bool {
	for _, ls := range l {
		if ls == s {
			return true
		}
	}

	return false
}
