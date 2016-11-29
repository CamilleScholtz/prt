package pkgfile

import (
	"regexp"
	"strings"
)

func Comment(file []byte, name string) string {
	regex := regexp.MustCompile("(?m)^# " + name + ":[\t\f ]*(.*)")
	match := regex.FindSubmatch(file)
	if len(match) < 1 {
		return ""
	}

	return string(match[1])
}

func Depends(file []byte, name string) []string {
	regex := regexp.MustCompile("(?m)^# Depends on:[\t\f ]*(.*)")
	match := regex.FindSubmatch(file)
	if len(match) < 1 {
		return []string{}
	}

	// Some Pkgfiles use commas, remove them
	fix := strings.Replace(string(match[1]), ",", "", -1)

	return strings.Split(fix, " ")
}

func Var(file []byte, name string) string {
	regex := regexp.MustCompile("(?m)^" + name + "=([a-z0-9-_+.]*)")
	match := regex.FindSubmatch(file)

	return string(match[1])
}
