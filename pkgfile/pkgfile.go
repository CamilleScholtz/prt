package pkgfile

import (
	"fmt"
	"regexp"
	"strings"
)

// Comment reads a comment from a (Pkg)file
func Comment(file []byte, value string) (string, error) {
	// We use (?m)^ here because there is always a comment on top of the file
	// so the use of (the slightly more optimal?) \n is impossible
	r := regexp.MustCompile("(?m)^# " + value + ":[[:blank:]]*(.*)")
	match := r.FindSubmatch(file)

	if len(match) == 0 {
		return "", fmt.Errorf("Could not read Pkgfile comment '" + value + "'!")
	}

	return string(match[1]), nil
}

// Depends reads the depends comment from a (Pkg)file
func Depends(file []byte, value string) ([]string, error) {
	r := regexp.MustCompile("\n# " + value + ":[[:blank:]]*(.*)")
	match := r.FindSubmatch(file)

	if len(match) == 0 {
		return []string{}, fmt.Errorf("Could not read Pkgfile comment '" + value + "'!")
	}

	// Some Pkgfiles use commas, remove them
	fix := strings.Replace(string(match[1]), ",", "", -1)

	return strings.Split(fix, " "), nil
}

// Var reads a variable from a (Pkg)file
func Var(file []byte, value string) (string, error) {
	r := regexp.MustCompile("\n" + value + "=([[:word:].-]*)")
	match := r.FindSubmatch(file)

	if len(match) == 0 {
		return "", fmt.Errorf("Could not read Pkgfile variable '" + value + "'!")
	}

	return string(match[1]), nil
}
