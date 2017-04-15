package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
)

// portsAlias aliases ports using the config.Alias values.
func portAlias(p string) string {
	for _, a := range config.Alias {
		if a[0] == p {
			p = a[1]
		}
	}

	return p
}

// allPorts lists all ports found in the PrtDir.
func allPorts() ([]string, error) {
	// TODO: Is there something more efficient than Glob?
	dl, err := filepath.Glob(path.Join(config.PrtDir, "/*/*/Pkgfile"))
	if err != nil {
		return []string{}, err
	}

	// Remove PrtDir from output.
	var p []string
	for _, d := range dl {
		p = append(p, portBaseLoc(path.Dir(d)))
	}

	return p, nil
}

// portBaseLoc removes the PrtDir from a string.
// TODO: Should I move this to port.go?
// TODO: Make it path.Base if config.PrtDir is not found.
func portBaseLoc(d string) string {
	return strings.Replace(d, config.PrtDir+"/", "", 1)
}

// portFullLoc adds the PrtDir to a string.
// TODO: Should I move this to port.go?
func portFullLoc(d string) string {
	return path.Join(config.PrtDir, d)
}

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

// portLoc tries to get the location of a port.
func portLoc(ports []string, n string) ([]string, error) {
	var l []string
	for _, p := range ports {
		if path.Base(p) == n {
			l = append(l, p)
		}
	}

	if len(l) == 0 {
		return []string{}, fmt.Errorf("loc %s: Could not find location", n)
	}

	// If there are multiple matches, sort using the config Order value.
	if len(l) > 1 {
		var i int
		for _, r := range config.Order {
			nl := path.Join(r, path.Base(l[i]))
			if stringInList(nl, ports) {
				l[i] = nl
				i++
			}

			// Break if everything has been ordered.
			if i == len(l) {
				break
			}
		}
	}

	return l, nil
}
