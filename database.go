// database.go contains functions that interact with the package database, this
// is a file called `db` found in `/var/lib/pkg`. These include functions such
// as parsing the database file.

package main

import (
	"bufio"
	"os"
)

// database is a struct type describing the pkg database file.
type database struct {
	Name    []string
	Version []string
}

// parseDatabase lists all installed ports.
func parseDatabase() (database, error) {
	f, err := os.Open("/var/lib/pkg/db")
	defer f.Close()
	if err != nil {
		return database{}, err
	}
	s := bufio.NewScanner(f)

	var db database
	var blank, file, name bool
	for s.Scan() {
		if blank || !file {
			db.Name = append(db.Name, s.Text())
			blank, name = false, true
		} else if name {
			db.Version = append(db.Version, s.Text())
			name = false
		} else if s.Text() == "" {
			blank = true
		}
		file = true
	}

	return db, nil
}
