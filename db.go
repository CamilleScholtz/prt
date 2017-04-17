package main

import (
	"bufio"
	"os"
)

// instPorts lists all installed ports.
func instPorts() ([]string, error) {
	db, err := os.Open("/var/lib/pkg/db")
	if err != nil {
		return []string{}, err
	}
	defer db.Close()
	s := bufio.NewScanner(db)

	// Check for versions.
	var b, f bool
	var p []string
	for s.Scan() {
		if b || !f {
			p = append(p, s.Text())
			b = false
		} else if s.Text() == "" {
			b = true
		}
		f = true
	}

	return p, nil
}

// instVersPorts list all installed versions, this should follow the same order
// as Inst().
func instVersPorts() ([]string, error) {
	db, err := os.Open("/var/lib/pkg/db")
	if err != nil {
		return []string{}, err
	}
	defer db.Close()
	s := bufio.NewScanner(db)

	// Check for versions.
	var b, f, n bool
	var v []string
	for s.Scan() {
		if b || !f {
			b, n = false, true
		} else if n {
			v = append(v, s.Text())
			n = false
		} else if s.Text() == "" {
			b = true
		}
		f = true
	}

	return v, nil
}
