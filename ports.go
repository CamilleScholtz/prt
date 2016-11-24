package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func GetPortAlias(port string) string {
	for _, alias := range Config.Alias {
		if alias[0] == port {
			port = alias[1]
		}
	}

	return port
}

// This function returns the port location
func GetPortLoc(name string) []string {
	var ports []string
	for _, port := range AllPorts {
		if strings.Split(port, "/")[1] == name {
			ports = append(ports, port)
		}
	}

	// If there are multiple matches, sort using the config Order value
	if len(ports) > 1 {
		for i, port := range ports {
			for _, repo := range Config.Order {
				if repo == filepath.Dir(port) {
					ports[i] = port
				}
			}
		}
	}

	return ports
}

// This functions lists all ports
func ListAllPorts() []string {
	// TODO: Is there something more efficient than Glob?
	dirs, err := filepath.Glob(Config.PortDir + "/*/*/Pkgfile")
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

// This functions lists installed ports
func ListInstPorts() []string {
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
