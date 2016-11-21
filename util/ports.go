package util

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/onodera-punpun/prt/config"
)

// This functions lists all ports
func ListAllPorts() []string {
	// TODO: Is there something more efficient than Glob?
	dirs, err := filepath.Glob(config.Struct.PrtDir + "/*/*")
	if err != nil {
		fmt.Println("Could not read ports.")
		os.Exit(1)
	}

	// Count seperators in PrtDir
	sep := strings.Count(config.Struct.PrtDir, "/")

	var ports []string
	for _, port := range dirs {
		// TODO: Test if these sep+ values actually are correct
		// TODO: Is there an easier way of basically getting the last 2 directories?
		ports = append(ports, strings.SplitAfterN(port, "/", sep+2)[sep+1])
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
func GetPortLoc(ports []string, port string) []string {
	regex := regexp.MustCompile("(?m)^.*/" + port + "$")

	return regex.FindAllString(strings.Join(ports, "\n"), -1)
}
