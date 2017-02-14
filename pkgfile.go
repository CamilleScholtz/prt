package main

import (
	"fmt"
	"io/ioutil"
	"path"
	"regexp"
	"strings"
)

// pkgfile is a stuct with all Pkgfile values.
var pkgfile struct {
	Description string
	URL         string
	Maintainer  string
	Depends     []string
	Optional    []string

	Name    string
	Version string
	Release string
	Source  []string
}

// comment reads a comment from a (Pkg)file.
func comment(f []byte, v string) (string, error) {
	r := regexp.MustCompile("(?m)^# " + v + ":[[:blank:]]*(.*)")
	m := r.FindSubmatch(f)

	if len(m) == 0 {
		return "", fmt.Errorf("pkgfile %s: No such Pkgfile comment", v)
	}

	return string(m[1]), nil
}

// variable reads a variable from a (Pkg)file.
func variable(f []byte, v string) (string, error) {
	r := regexp.MustCompile("\n" + v + "=(.*)")
	m := r.FindSubmatch(f)

	if len(m) == 0 {
		return "", fmt.Errorf("pkgfile %s: No such Pkgfile variable", v)
	}

	return string(m[1]), nil
}

// initPkgfile initializes the pkgfile struct.
func initPkgfile(l string, vl []string) error {
	// Read out Pkgfile.
	f, err := ioutil.ReadFile(path.Join(l, "Pkgfile"))
	if err != nil {
		return err
	}

	for _, v := range vl {
		switch v {
		case "Description":
			pkgfile.Description, err = comment(f, "Description")
			if err != nil {
				return err
			}
		case "URL":
			pkgfile.URL, err = comment(f, "URL")
			if err != nil {
				return err
			}
		case "Maintainer":
			pkgfile.Maintainer, err = comment(f, "Maintainer")
			if err != nil {
				return err
			}
		case "Depends":
			d, err := comment(f, "Depends on")
			if err != nil {
				return err
			}
			pkgfile.Depends = strings.Split(strings.Replace(d, ",", "", -1), " ")
		case "Optional":
			d, err := comment(f, "Nice to have|Optional")
			if err != nil {
				return err
			}
			pkgfile.Optional = strings.Split(strings.Replace(d, ",", "", -1), " ")
		case "Name":
			pkgfile.Name, err = variable(f, "name")
			if err != nil {
				return err
			}
		case "Version":
			pkgfile.Version, err = variable(f, "version")
			if err != nil {
				return err
			}
		case "Release":
			pkgfile.Release, err = variable(f, "release")
			if err != nil {
				return err
			}
		case "Source":
			// TODO
			pkgfile.Source = []string{}
		default:
			return fmt.Errorf("pkgfile %s, No such pkgfile variable or comment", v)
		}
	}

	return nil
}
