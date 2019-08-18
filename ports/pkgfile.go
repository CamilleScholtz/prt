package ports

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"path"
	"strings"

	"mvdan.cc/sh/v3/shell"
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
// force the use of a bash interpreter to get the `source` BASH array of a
// `Pkgfile`, this is relatively slow.
func (f *Pkgfile) Parse(source ...bool) error {
	r, err := os.Open(path.Join(f.Location.Full(), "Pkgfile"))
	if err != nil {
		return fmt.Errorf("could not open `%s/Pkgfile`", f.Location.Full())
	}
	defer r.Close()
	s := bufio.NewScanner(r)

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
					v, err := shell.SourceFile(context.TODO(), path.Join(f.
						Location.Full(), "Pkgfile"))
					if err != nil {
						return err
					}

					f.Source = v["source"].List
				}

				// Since `source` should be the last meaningfull value in a
				// `Pkgfile`, we will stop walking.
				return nil
			}
		}
	}

	return nil
}

// Validate checks if all required variables and functions are in a `Pkgfile`
// file.
// TODO: Check for `build()` function
func (f *Pkgfile) Validate() error {
	if f.Name == "" {
		return fmt.Errorf("pkgfile %s: Name variable is empty", f.Location.
			Base())
	}
	if f.Version == "" {
		return fmt.Errorf("pkgfile %s: Version variable is empty", f.Location.
			Base())
	}
	if f.Release == "" {
		return fmt.Errorf("pkgfile %s: Release variable is empty", f.Location.
			Base())
	}

	return nil
}
