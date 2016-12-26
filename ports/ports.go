package ports

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"

	"github.com/onodera-punpun/prt/config"
	"github.com/onodera-punpun/prt/utils"
)

// Load config
var c = config.Load()

// Alias aliases ports using the config values
func Alias(port string) string {
	for _, alias := range c.Alias {
		if alias[0] == port {
			port = alias[1]
		}
	}

	return port
}

// All lists all ports found in the PortDir
func All() ([]string, error) {
	// TODO: Is there something more efficient than Glob?
	dirs, err := filepath.Glob(filepath.Join(c.PortDir, "/*/*/Pkgfile"))
	if err != nil {
		return []string{}, fmt.Errorf("Could not read '" + filepath.Join(c.PortDir, "/*/*/Pkgfile") + "'!")
	}

	// TODO: Use string replace or delete here
	var ports []string
	for _, loc := range dirs {
		path := filepath.Dir(loc)
		repo := filepath.Base(filepath.Dir(path))
		port := filepath.Base(path)
		ports = append(ports, filepath.Join(repo, port))
	}

	return ports, nil
}

// Inst lists all installed ports
func Inst() ([]string, error) {
	// TODO: Use filepath stuff here?
	db, err := os.Open("/var/lib/pkg/db")
	if err != nil {
		return []string{}, fmt.Errorf("Could not read '/var/lib/pkg/db'!")
	}
	defer db.Close()
	s := bufio.NewScanner(db)

	var blank bool
	var ports []string
	for s.Scan() {
		if blank {
			ports = append(ports, s.Text())
			blank = false
		} else if s.Text() == "" {
			blank = true
		}
	}

	return ports, nil
}

// InstVers list all installed versions, this should follow the same order as Inst()
func InstVers() ([]string, error) {
	// TODO: Use filepath stuff here?
	db, err := os.Open("/var/lib/pkg/db")
	if err != nil {
		return []string{}, fmt.Errorf("Could not read '/var/lib/pkg/db'!")
	}
	defer db.Close()
	s := bufio.NewScanner(db)

	var blank, name bool
	var vers []string
	for s.Scan() {
		if blank {
			blank, name = false, true
		} else if name {
			vers = append(vers, s.Text())
			name = false
		} else if s.Text() == "" {
			blank = true
		}
	}

	return vers, nil
}

// Loc tries to get the location of a port
func Loc(ports []string, name string) ([]string, error) {
	var locs []string
	for _, port := range ports {
		if filepath.Base(port) == name {
			locs = append(locs, port)
		}
	}

	if len(locs) == 0 {
		return []string{}, fmt.Errorf("Could not find location for '" + name + "'!")
	}

	// If there are multiple matches, sort using the config Order value
	if len(locs) > 1 {
		var i int
		for _, repo := range c.Order {
			newLoc := filepath.Join(repo, filepath.Base(locs[i]))
			if utils.StringInList(newLoc, ports) {
				locs[i] = newLoc
				i++
			}

			// Break if everything has been ordered
			if i == len(locs) {
				break
			}
		}
	}

	return locs, nil
}
