package ports

import (
	"bufio"
	"os"
)

// Database is a type describing the package database file. This file lists all
// installed packages, the version of these packages, and the files these
// packages installed.
type Database struct {
	Packages []Package
}

// Parse parses the `db` file and populates the various fields in the given
// `*Database`.
func (f *Database) Parse() error {
	r, err := os.Open("/var/lib/pkg/db")
	if err != nil {
		return err
	}
	defer r.Close()
	s := bufio.NewScanner(r)

	var blank, file, name bool
	var p Package
	for s.Scan() {
		if blank || !file {
			p = Package{}
			p.Name = s.Text()
			blank, file, name = false, true, true
		} else if name {
			p.Version = s.Text()
			name = false
		} else if s.Text() == "" {
			f.Packages = append(f.Packages, p)
			blank = true
		} else {
			p.Files = append(p.Files, s.Text())
		}
	}

	return nil
}
