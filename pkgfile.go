package main

import (
	"regexp"
	"strings"
)

// This function reads out Pkgfile comments
func ReadComment(file []byte, name string) string {
	regex := regexp.MustCompile("(?m)^# " + name + ":[\t\f ]*(.*)")
	match := regex.FindSubmatch(file)
	if len(match) < 1 {
		return ""
	}

	return string(match[1])
}

// This function reads out Pkgfile dependencies
// This is pretty much like ReadComment,
// but it also removes commas and makes a list out if it
func ReadDepends(file []byte, name string) []string {
	regex := regexp.MustCompile("(?m)^# Depends on:[\t\f ]*(.*)")
	match := regex.FindSubmatch(file)
	if len(match) < 1 {
		return []string{}
	}

	// Some Pkgfiles use commas, remove them
	fix := strings.Replace(string(match[1]), ",", "", -1)

	return strings.Split(fix, " ")
}

// This function reads out Pkgfile variables
func ReadVar(file []byte, name string) string {
	regex := regexp.MustCompile("(?m)^" + name + "=([a-z0-9-_+.]*)")
	match := regex.FindSubmatch(file)

	return string(match[1])
}
