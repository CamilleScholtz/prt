package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// This functions lists all ports
func ListAllPorts() []string {
	// TODO: Is there something more efficient than Glob?
	dirs, err := filepath.Glob(Config.PortDir + "/*/*/Pkgfile")
	if err != nil {
		fmt.Println("Could not read ports.")
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
// TODO: This could use some optimization
func ListInstPorts() []string {
	regex := regexp.MustCompile("(?m)^$\n(.*)")

	// Read out db
	db, err := ioutil.ReadFile("/var/lib/pkg/db")
	if err != nil {
		fmt.Println("Could not read pkg db.")
		os.Exit(1)
	}

	// Convert byte[][][] to string[]
	var ports []string
	for _, port := range regex.FindAllSubmatch(db, -1) {
		ports = append(ports, string(port[:][1]))
	}

	return ports
}

// This function returns the port location
func GetPortLoc(port string) []string {
	regex := regexp.MustCompile(".*/" + port + "$")

	var ports []string
	for _, port := range AllPorts {
		if regex.MatchString(port) {
			ports = append(ports, regex.FindString(port))
		}
	}

	// If there are multiple matches, sort using RepoOrder
	var repoPort []string
	if len(ports) > 1 {
		// Empty old port array
		oldPorts := ports
		ports = []string{}

		for _, port := range oldPorts {
			repoPort = strings.Split(port, "/")

			for _, repo := range Config.RepoOrder {
				if repo == repoPort[0] {
					ports = append(ports, strings.Join(repoPort, "/"))
				}
			}
		}
	}

	return ports
}
