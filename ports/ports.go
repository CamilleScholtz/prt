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

func Alias(port string) string {
	for _, alias := range config.Struct.Alias {
		if alias[0] == port {
			port = alias[1]
		}
	}

	return port
}

func All() ([]string, error) {
	// TODO: Is there something more efficient than Glob?
	dirs, err := filepath.Glob(config.Struct.PortDir + "/*/*/Pkgfile")
	if err != nil {
		return []string{}, fmt.Errorf("Could not read '" + config.Struct.PortDir + "/*/*/Pkgfile'!")
	}

	var ports []string
	for _, port := range dirs {
		path := strings.Split(filepath.Dir(port), "/")
		ports = append(ports, strings.Join(path[len(path)-2:], "/"))
	}

	return ports, nil
}

func Inst() ([]string, error) {
	db, err := os.Open("/var/lib/pkg/db")
	if err != nil {
		return []string{}, fmt.Errorf("Could not read '/var/lib/pkg/db'!")
	}

	defer db.Close()
	scanner := bufio.NewScanner(db)

	var match bool
	var ports []string
	for scanner.Scan() {
		if match {
			ports = append(ports, scanner.Text())
			match = false
		} else if scanner.Text() == "" {
			match = true
		}
	}

	return ports, nil
}

func Loc(ports []string, name string) ([]string, error) {
	var locs []string
	for _, port := range ports {
		if strings.Split(port, "/")[1] == name {
			locs = append(locs, port)
		}
	}

	if len(locs) == 0 {
		return []string{}, fmt.Errorf("Could not find location for '" + name + "'!")
	}

	// If there are multiple matches, sort using the config Order value
	if len(locs) > 1 {
		var i int
		for _, repo := range config.Struct.Order {
			newLoc := repo + "/" + filepath.Base(locs[i])
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

func InstVer(name string) (string, error) {
	db, err := os.Open("/var/lib/pkg/db")
	if err != nil {
		return "", fmt.Errorf("Could not read '/var/lib/pkg/db'!")
	}

	defer db.Close()
	scanner := bufio.NewScanner(db)

	var match bool
	var ver string
	for scanner.Scan() {
		if match {
			ver = scanner.Text()
			break
		} else if scanner.Text() == name {
			match = true
		}
	}

	if len(ver) == 0 {
		return "", fmt.Errorf("Could not find installed version of '" + name + "'!")
	}

	return ver, nil
}