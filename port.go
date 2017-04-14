package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"regexp"
)

type port struct {
	Footprint []byte
	Md5sum    []byte
	Pkgfile   []byte
	Signature []byte

	Loc string
}

// comment reads a comment from a Pkgfile.
func (p port) comment(v string) (string, error) {
	r := regexp.MustCompile("(?m)^# " + v + ":[[:blank:]]*(.*)")
	m := r.FindSubmatch(p.Pkgfile)

	if len(m) == 0 {
		return "", fmt.Errorf("pkgfile %s: No such comment", v)
	}

	return string(m[1]), nil
}

// function checks if for a function in a Pkgfile.
func (p port) function(f string) error {
	cmd := exec.Command("bash", "-c", "source ./Pkgfile && type -t "+f)
	cmd.Dir = p.Loc
	var b bytes.Buffer
	cmd.Stdout = &b

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("pkgfile %s: Could not source", f)
	}

	if b.String() != "function\n" {
		return fmt.Errorf("pkgfile %s: No such function", f)
	}

	return nil
}

// readFootprint reads a .footprint.
func readFootprint(l string) (port, error) {
	if _, err := os.Stat(path.Join(l, ".footprint")); err != nil {
		return port{}, nil
	}

	f, err := ioutil.ReadFile(path.Join(l, ".footprint"))
	if err != nil {
		return port{}, err
	}

	return port{f, nil, nil, nil, l}, nil
}

// readMd5sum reads a .md5sum.
func readMd5sum(l string) (port, error) {
	if _, err := os.Stat(path.Join(l, ".md5sum")); err != nil {
		return port{}, nil
	}

	m, err := ioutil.ReadFile(path.Join(l, ".md5sum"))
	if err != nil {
		return port{}, err
	}

	return port{nil, m, nil, nil, l}, nil
}

// readFootprint reads a Pkgfile.
func readPkgfile(l string) (port, error) {
	p, err := ioutil.ReadFile(path.Join(l, "Pkgfile"))
	if err != nil {
		return port{}, err
	}

	return port{nil, nil, p, nil, l}, nil
}

// readFootprint reads a .signatrure.
func readSignature(l string) (port, error) {
	if _, err := os.Stat(path.Join(l, ".signature")); err != nil {
		return port{}, nil
	}

	s, err := ioutil.ReadFile(path.Join(l, ".signature"))
	if err != nil {
		return port{}, err
	}

	return port{nil, nil, nil, s, l}, nil
}

// readPort reads a Pkfile, .footprint and .md5sum.
// TODO: Can I make this simpler somehow, just use return port
// instead of this return port{asdasd} stuff?
func readPort(l string) (port, error) {
	f, err := readFootprint(l)
	if err != nil {
		return port{}, err
	}
	m, err := readMd5sum(l)
	if err != nil {
		return port{}, err
	}
	p, err := readPkgfile(l)
	if err != nil {
		return port{}, err
	}
	s, err := readSignature(l)
	if err != nil {
		return port{}, err
	}

	return port{f.Footprint, m.Md5sum, p.Pkgfile, s.Signature, l}, nil
}

// variable reads a variable from a Pkgfile.
func (p port) variable(v string) (string, error) {
	r := regexp.MustCompile("\n" + v + "=(.*)")
	m := r.FindSubmatch(p.Pkgfile)

	if len(m) == 0 {
		return "", fmt.Errorf("pkgfile %s: No such variable", v)
	}

	return string(m[1]), nil
}

// variableSource reads a variable from a Pkgfile, it's like the variable
// function, but actually uses bash source. This is a lot slower but also more
// precise because it completes variables. This is especially (only?) seful for
// the source variable in Pkgfiles.
func (p port) variableSource(v string) (string, error) {
	cmd := exec.Command("bash", "-c", "source ./Pkgfile && echo ${"+v+"[@]}")
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
