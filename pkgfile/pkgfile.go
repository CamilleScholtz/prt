package pkgfile

import (
	"fmt"
	"regexp"
)

// Comment reads a comment from a (Pkg)file.
func Comment(f []byte, v string) (string, error) {
	r := regexp.MustCompile("(?m)^# " + v + ":[[:blank:]]*(.*)")
	m := r.FindSubmatch(f)

	if len(m) == 0 {
		return "", fmt.Errorf("comment %s: No such Pkgfile comment", v)
	}

	return string(m[1]), nil
}

// Variable reads a variable from a (Pkg)file.
func Variable(f []byte, v string) (string, error) {
	r := regexp.MustCompile("\n" + v + "=(.*)")
	m := r.FindSubmatch(f)

	if len(m) == 0 {
		return "", fmt.Errorf("var %s: No such Pkgfile variable", v)
	}

	return string(m[1]), nil
}
