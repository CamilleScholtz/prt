package pkgfile

import (
	"fmt"
	"regexp"
	"strings"
)

// Comment reads a comment from a (Pkg)file
func Comment(file []byte, value string) (string, error) {
	regex := regexp.MustCompile("(?m)^# " + value + ":[\t\f ]*(.*)")
	match := regex.FindSubmatch(file)

	if len(match[1]) == 0 {
		return "", fmt.Errorf("Could not read Pkgfile comment '" + value + "'!")
	}

	return string(match[1]), nil
}

// Depends reads the depends comment from a (Pkg)file
func Depends(file []byte, value string) ([]string, error) {
	regex := regexp.MustCompile("(?m)^# " + value + ":[\t\f ]*(.*)")
	match := regex.FindSubmatch(file)

	if len(match) == 0 {
		return []string{}, fmt.Errorf("Could not read Pkgfile comment '" + value + "'!")
	}

	// Some Pkgfiles use commas, remove them
	fix := strings.Replace(string(match[1]), ",", "", -1)

	return strings.Split(fix, " "), nil
}

// Var reads a variable from a (Pkg)file
func Var(file []byte, value string) (string, error) {
	regex := regexp.MustCompile("(?m)^" + value + "=([a-z0-9-_+.]*)")
	match := regex.FindSubmatch(file)

	if len(match[1]) == 0 {
		return "", fmt.Errorf("Could not read Pkgfile variable '" + value + "'!")
	}

	return string(match[1]), nil
}
