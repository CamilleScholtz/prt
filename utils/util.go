package utils

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
	files, err := filepath.Glob("/usr/ports/*/*/Pkgfile")
	if err != nil {
		fmt.Println("Could not read ports.")
		os.Exit(1)
	}

	var ports []string
	for _, port := range files {
		ports = append(ports, strings.Replace(port, "/Pkgfile", "", 1))
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
func PortLoc(ports []string, port string) []string {
	regex := regexp.MustCompile("(?m)^.*/" + port + "$")

	return regex.FindAllString(strings.Join(ports, "\n"), -1)
}

// This function reads out Pkgfile comments
func ReadComment(file []byte, name string) string {
	regex := regexp.MustCompile("(?m)^# " + name + ":[\t\f ]*(.*)")
	match := regex.FindSubmatch(file)
	if len(match) < 1 {
		return ""
	}

	return string(match[1])
}

// This function reads out Pkgfile dependencies
// This is pretty much like ReadComment,
// but it also removes commas and makes a list out if it
func ReadDepends(file []byte, name string) []string {
	regex := regexp.MustCompile("(?m)^# Depends on:[\t\f ]*(.*)")
	match := regex.FindSubmatch(file)
	if len(match) < 1 {
		return []string{}
	}

	// Some Pkgfiles use commas, remove them
	fix := strings.Replace(string(match[1]), ",", "", -1)

	return strings.Split(fix, " ")
}

// This function reads out Pkgfile variables
func ReadVar(file []byte, name string) string {
	regex := regexp.MustCompile("(?m)^" + name + "=([a-z0-9-_+.]*)")
	match := regex.FindSubmatch(file)

	return string(match[1])
}

// This function checks if a string is in a list
func StringInList(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}

	return false
}
