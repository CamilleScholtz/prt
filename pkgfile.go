package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os/exec"
	"path"
	"regexp"
)

type pkgfile struct {
	Content []byte
	Loc     string
}

// comment reads a comment from a Pkgfile.
func (p pkgfile) comment(v string) (string, error) {
	r := regexp.MustCompile("(?m)^# " + v + ":[[:blank:]]*(.*)")
	m := r.FindSubmatch(p.Content)

	if len(m) == 0 {
		return "", fmt.Errorf("pkgfile %s: No such comment", v)
	}

	return string(m[1]), nil
}

// readPkgfile uhh... Reads a a Pkgfile.
func readPkgfile(l string) (pkgfile, error) {
	p, err := ioutil.ReadFile(l)
	if err != nil {
		return pkgfile{}, err
	}

	return pkgfile{p, path.Dir(l)}, nil
}

// variable reads a variable from a Pkgfile.
func (p pkgfile) variable(v string) (string, error) {
	r := regexp.MustCompile("\n" + v + "=(.*)")
	m := r.FindSubmatch(p.Content)

	if len(m) == 0 {
		return "", fmt.Errorf("pkgfile %s: No such variable", v)
	}

	return string(m[1]), nil
}

// variableSource reads a variable from a Pkgfile, it's like the variable
// function, but actually uses bash source. This is a lot slower but also
// more precise, because it completes variables. This is especially (only?)
// useful for the source variable in Pkgfiles.
func (p pkgfile) variableSource(v string) (string, error) {
	cmd := exec.Command("bash", "-c", "source ./Pkgfile && echo $"+v)
	cmd.Dir = p.Loc
	var b bytes.Buffer
	cmd.Stdout = &b

	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("pkgfile %s: Could not source", v)
	}

	if len(b.String()) == 0 {
		return "", fmt.Errorf("pkgfile %s: No such variable", v)
	}

	return b.String(), nil
}
