package pkgfile

import (
	"fmt"
	"regexp"
	"strings"
)

func Comment(file []byte, value string) (string, error) {
	regex := regexp.MustCompile("(?m)^# " + value + ":[\t\f ]*(.*)")
	match := regex.FindSubmatch(file)

	if len(match[1]) == 0 {
		return "", fmt.Errorf("Could not read Pkgfile comment '" + value + "'!")
	}

	return string(match[1]), nil
}

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

func Var(file []byte, value string) (string, error) {
	regex := regexp.MustCompile("(?m)^" + value + "=([a-z0-9-_+.]*)")
	match := regex.FindSubmatch(file)

	if len(match[1]) == 0 {
		return "", fmt.Errorf("Could not read Pkgfile variable '" + value + "'!")
	}

	return string(match[1]), nil
}
