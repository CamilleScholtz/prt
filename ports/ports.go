package ports

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/onodera-punpun/prt/config"
	"github.com/onodera-punpun/prt/utils"
)

// Load config.
var c = config.Load()

// Alias aliases ports using the config values.
func Alias(p string) string {
	for _, a := range c.Alias {
		if a[0] == p {
			p = a[1]
		}
	}

	return p
}

// All lists all ports found in the PortDir.
func All() ([]string, error) {
	// TODO: Is there something more efficient than Glob?
	dirs, err := filepath.Glob(filepath.Join(c.PortDir, "/*/*/Pkgfile"))
	if err != nil {
		return []string{}, err
	}

	// Remove PortDir from output.
	var p []string
	for _, d := range dirs {
		p = append(p, strings.Replace(filepath.Dir(d), c.PortDir+"/", "", 1))
	}

	return p, nil
}

// Inst lists all installed ports.
func Inst() ([]string, error) {
	// Read out pkg db.
	// TODO: Use filepath stuff here?
	db, err := os.Open("/var/lib/pkg/db")
	if err != nil {
		return []string{}, err
	}
	defer db.Close()
	s := bufio.NewScanner(db)

	// Check for versions.
	var b bool
	var p []string
	for s.Scan() {
		if b {
			p = append(p, s.Text())
			b = false
		} else if s.Text() == "" {
			b = true
		}
	}

	return p, nil
}

// InstVers list all installed versions, this should follow the same order as Inst().
func InstVers() ([]string, error) {
	// Read out pkg db.
	// TODO: Use filepath stuff here?
	db, err := os.Open("/var/lib/pkg/db")
	if err != nil {
		return []string{}, err
	}
	defer db.Close()
	s := bufio.NewScanner(db)

	// Check for versions.
	var b, n bool
	var v []string
	for s.Scan() {
		if b {
			b, n = false, true
		} else if n {
			v = append(v, s.Text())
			n = false
		} else if s.Text() == "" {
			b = true
		}
	}

	return v, nil
}

// Loc tries to get the location of a port.
func Loc(ports []string, n string) ([]string, error) {
	var l []string
	for _, p := range ports {
		if filepath.Base(p) == n {
			l = append(l, p)
		}
	}

	if len(l) == 0 {
		return []string{}, fmt.Errorf("loc %s: Could not find location", n)
	}

	// If there are multiple matches, sort using the config Order value.
	if len(l) > 1 {
		var i int
		for _, r := range c.Order {
			nl := filepath.Join(r, filepath.Base(l[i]))
			if utils.StringInList(nl, ports) {
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
