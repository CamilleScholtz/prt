package ports

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"
)

// A Pkgfile is a type describing the `Pkgfile` file of a port. This file
// contains information about the package (such as `name`, `version`, et cetera)
// and the commands that should be executed in order to compile the package in
// question.
type Pkgfile struct {
	*Port

	// Comments with various information that isn't strictly needed in order to
	// build a package.
	Description string
	URL         string
	Maintainer  string

	// Comments with information about dependencies. These need some more
	// parsing because some `Pkgfile`s use commas to separate dependencies, and
	// others use spaces.
	Depends  []string
	Optional []string

	// BASH variables with various information that is required in order to
	// build a package.
	Name    string
	Version string
	Release string

	// A BASH array with the sources needed to build a package. We probably need
	// to parse this by actually using `source(1)` because `Pkgfile`s often use
	// BASH variables (such as `$name` or `$version`) and bashisms in the source
	// variable.
	Source []string
}

// Parse parses the `Pkgfile` file of a port and populates the various fields in
// the given `*Pkgfile`. Keep in mind that this does not expand BASH ariables by
// default. so `$version` will just be a literal string. Nor does this parse the
// `source` field of a `Pkgfile` because it often uses variables in the string
// and because it's simply too hard too parse.
//
// If you want to expand BASH variables pass a bool as a parameter. This will
// force the use of `source(1)` to get the `source` BASH array of a `Pkgfile`.
// Using `source(1)` is relatively slow.
func (f *Pkgfile) Parse(source ...bool) error {
	fr, err := os.Open(path.Join(f.Location.Full(), "Pkgfile"))
	defer fr.Close()
	if err != nil {
		return fmt.Errorf("could not open `%s/Pkgfile`", f.Location.Full())
	}
	s := bufio.NewScanner(fr)

	for s.Scan() {
		i := s.Text()

		if strings.HasPrefix(i, "#") {
			kv := strings.SplitN(i, ":", 2)

			switch kv[0] {
			case "# Description":
				f.Description = strings.TrimSpace(kv[1])
			case "# URL":
				f.URL = strings.TrimSpace(kv[1])
			case "# Maintainer":
				f.Maintainer = strings.TrimSpace(kv[1])
			case "# Depends on":
				f.Depends = strings.Fields(strings.Replace(strings.TrimSpace(
					kv[1]), ",", "", -1))
			case "# Optional":
			case "# Nice to have":
				f.Optional = strings.Fields(strings.Replace(strings.TrimSpace(
					kv[1]), ",", "", -1))
			}
		} else {
			kv := strings.SplitN(i, "=", 2)

			switch kv[0] {
			case "name":
				f.Name = strings.TrimSpace(kv[1])
			case "version":
				f.Version = strings.TrimSpace(kv[1])
			case "release":
				f.Release = strings.TrimSpace(kv[1])
			case "source":
				if len(source) > 0 {
					// TODO: Possibly use `mvdan.cc/sh/interp` for this.
					s, err := f.Expand("source")
					if err != nil {
						return err
					}
					f.Source = strings.Fields(s)
				}

				// Since `source` should be the last meaningfull value in a
				// `Pkgfile`, we will stop walking.
				return nil
			}
		}
	}

	return nil
}

// Global variables used by `RecursiveDepends()`.
// TODO: Not that clean, can I move this?
var check []string
var depends []Port

// RecursiveDepends is a function that calculates dependencies recursively. This
// function requires `Parse()` to be run on the `Pkgfile` in question
// beforehand.
func (f *Pkgfile) RecursiveDepends(aliases [][]Location, order []string,
	all []Port) ([]Port, error) {
	// Continue if already checked.
	if stringInStrings(f.Location.Port, check) {
		return depends, nil
	}

	for _, n := range f.Depends {
		pl, err := Locate(n, order, all)
		if err != nil {
			continue
		}
		d := pl[0]

		// Alias ports if needed.
		// TODO: Is this if needed?
		if len(aliases) != 0 {
			d.Alias(aliases)
		}

		// Read out Pkgfile.
		if err := d.Pkgfile.Parse(); err != nil {
			return []Port{}, err
		}

		// Append to `depends`.
		depends = append(depends, d)

		// Append port to checked ports.
		check = append(check, f.Location.Port)

		// Loop.
		depends[len(depends)-1].Pkgfile.RecursiveDepends(aliases, order, all)
	}

	return depends, nil
}

// Expand reads a variable from a `Pkgfile` file using `source(1)`. This is
// relatively slow but sometimes needed because it expands BASH variables. This
// is especially (only?) useful for the `source` BASH array in `Pkgfile` files.
func (f Pkgfile) Expand(key string) (string, error) {
	cmd := exec.Command("bash", "-c", "source ./Pkgfile && echo ${"+key+"[@]}")
	cmd.Dir = f.Location.Full()
	var b bytes.Buffer
	cmd.Stdout = &b

	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf(
			"could not source variable `%s` from `%s/Pkgfile`", key,
			f.Location.Full())
	}

	if len(b.String()) == 0 {
		return "", fmt.Errorf(
			"no variable with the name `%s` found in `%s/Pkgfile`", key,
			f.Location.Full())
	}

	return b.String(), nil
}
