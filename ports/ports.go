package ports

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/onodera-punpun/prt/config"
)

func All() []string {
	// TODO: Is there something more efficient than Glob?
	dirs, err := filepath.Glob(config.Struct.PortDir + "/*/*/Pkgfile")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Could not read ports!")
		os.Exit(1)
	}

	var ports []string
	for _, port := range dirs {
		path := strings.Split(filepath.Dir(port), "/")
		ports = append(ports, strings.Join(path[len(path)-2:], "/"))
	}

	return ports
}

func Inst() []string {
	var ports []string
	if db, err := os.Open("/var/lib/pkg/db"); err == nil {
		// Make sure it gets closed
		defer db.Close()

		// Create a new scanner and read the db line by line
		scanner := bufio.NewScanner(db)

		var empty bool
		for scanner.Scan() {
			if empty {
				ports = append(ports, scanner.Text())
				empty = false
			} else if scanner.Text() == "" {
				empty = true
			}
		}
	} else {
		fmt.Fprintln(os.Stderr, "Could not read pkg db!")
	}

	return ports
}

func Alias(port string) string {
	for _, alias := range config.Struct.Alias {
		if alias[0] == port {
			port = alias[1]
		}
	}

	return port
}

func Loc(ports []string, name string) []string {
	var locs []string
	for _, port := range ports {
		if strings.Split(port, "/")[1] == name {
			locs = append(locs, port)
		}
	}

	// If there are multiple matches, sort using the config Order value
	if len(locs) > 1 {
		for i, loc := range locs {
			for _, repo := range config.Struct.Order {
				if repo == filepath.Dir(loc) {
					locs[i] = loc
				}
			}
		}
	}

	return locs
}
