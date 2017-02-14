package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

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
func reverseList(l []string) []string {
	var nl []string
	for i := range l {
		nl = append(nl, l[len(l)-1-i])
	}

	return nl
}

// stringInList checks if a string is in a list.
func stringInList(s string, l []string) bool {
	for _, ls := range l {
		if ls == s {
			return true
		}
	}

	return false
}
