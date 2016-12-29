package ports

import (
	"bufio"
	"fmt"
	"os"
	"path"
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
	dl, err := filepath.Glob(path.Join(c.PortDir, "/*/*/Pkgfile"))
	if err != nil {
		return []string{}, err
	}

	// Remove PortDir from output.
	var p []string
	for _, d := range dl {
		p = append(p, BaseLoc(path.Dir(d)))
	}

	return p, nil
}

// BaseLoc removes the PortDir from a string.
func BaseLoc(d string) string {
	return strings.Replace(d, c.PortDir+"/", "", 1)
}

// FullLoc adds the PortDir to a string.
func FullLoc(d string) string {
	return path.Join(c.PortDir, d)
}

// Inst lists all installed ports.
func Inst() ([]string, error) {
	// Read out pkg db.
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
		for _, r := range c.Order {
			nl := path.Join(r, path.Base(l[i]))
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
