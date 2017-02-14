package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
)

type pkgfile struct {
	content []byte
}

// comment reads a comment from a Pkgfile.
func (p pkgfile) comment(v string) (string, error) {
	r := regexp.MustCompile("(?m)^# " + v + ":[[:blank:]]*(.*)")
	m := r.FindSubmatch(p.content)

	if len(m) == 0 {
		return "", fmt.Errorf("pkgfile %s: No such Pkgfile comment", v)
	}

	return string(m[1]), nil
}

// readPkgfile uhh... Reads a a Pkgfile.
func readPkgfile(l string) (pkgfile, error) {
	p, err := ioutil.ReadFile(l)
	if err != nil {
		return pkgfile{}, err
	}

	return pkgfile{p}, nil
}

// variable reads a variable from a Pkgfile.
func (p pkgfile) variable(v string) (string, error) {
	r := regexp.MustCompile("\n" + v + "=(.*)")
	m := r.FindSubmatch(p.content)

	if len(m) == 0 {
		return "", fmt.Errorf("pkgfile %s: No such Pkgfile variable", v)
	}

	return string(m[1]), nil
}
