package pkgfile

import (
	"fmt"
	"regexp"
	"strings"
)

// Comment reads a comment from a (Pkg)file.
func Comment(f []byte, v string) (string, error) {
	// We use (?m)^ here because there is always a comment on top of the file
	// so the use of (the slightly more optimal?) \n is impossible.
	r := regexp.MustCompile("(?m)^# " + v + ":[[:blank:]]*(.*)")
	m := r.FindSubmatch(f)

	if len(m) == 0 {
		return "", fmt.Errorf("comment %s: No such Pkgfile comment", v)
	}

	return string(m[1]), nil
}

// Depends reads the depends comment from a (Pkg)file.
func Depends(f []byte, v string) ([]string, error) {
	r := regexp.MustCompile("\n# " + v + ":[[:blank:]]*(.*)")
	m := r.FindSubmatch(f)

	if len(m) == 0 {
		return []string{}, fmt.Errorf("depends %s: No such Pkgfile comment", v)
	}

	return strings.Split(strings.Replace(string(m[1]), ",", "", -1), " "), nil
}

// Var reads a variable from a (Pkg)file.
func Var(f []byte, v string) (string, error) {
	r := regexp.MustCompile("\n" + v + "=([[:word:].-]*)")
	m := r.FindSubmatch(f)

	if len(m) == 0 {
		return "", fmt.Errorf("var %s: No such Pkgfile variable", v)
	}

	return string(m[1]), nil
}
