package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/chiyouhen/getopt"
	"github.com/fatih/color"
)

// Initialize variables
var all, alias, tree bool
var allPorts, checkPorts, instPorts []string
var i int

// This function prints dependencies recursivly
func recursive(path string) {
	// Read out Pkgfile
	pkgfile, err := ioutil.ReadFile(path + "/Pkgfile")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Could not read Pkgfile!")
		os.Exit(1)
	}

	// Read out Pkgfile dependencies
	deps := ReadDepends(pkgfile, "Depends on")

	var locs []string
	var loc string

	for _, dep := range deps {
		// Continue if already checked
		if StringInList(dep, checkPorts) {
			continue
		}
		checkPorts = append(checkPorts, dep)

		// Get port location
		locs = PortLoc(dep)
		if len(locs) < 1 {
			return
		}
		loc = locs[0]

		// Alias if needed
		if !alias {
			loc = PortAlias(loc)
		}

		// Continue if already installed
		if !all {
			if StringInList(filepath.Base(loc), instPorts) {
				continue
			}
		}

		// Print tree indentation
		if tree {
			if i > 0 {
				color.Set(color.FgBlack, color.Bold)
				fmt.Printf(strings.Repeat(Config.IndentChar, i))
				color.Unset()
			}
			i += 1
		}

		// Finally print the port :)
		fmt.Println(loc)

		// Loop
		recursive(Config.PortDir + "/" + loc)

		if tree {
			i -= 1
		}
	}
}

func Depends(args []string) {
	// Define opts
	shortopts := "hant"
	longopts := []string{
		"--help",
		"--no-alias",
		"--tree",
	}

	// Read out opts
	opts, _, err := getopt.Getopt(args, shortopts, longopts)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Invaild argument, use -h for a list of arguments!")
		os.Exit(1)
	}

	for _, opt := range opts {
		switch opt[0] {
		case "-h", "--help":
			fmt.Println("Usage: prt depends [arguments]")
			fmt.Println("")
			fmt.Println("arguments:")
			fmt.Println("  -a,   --all             also list installed dependencies")
			fmt.Println("  -n,   --no-alias        disable aliasing")
			fmt.Println("  -t,   --tree            list using tree view")
			fmt.Println("  -h,   --help            print help and exit")
			os.Exit(0)
		case "-a", "--all":
			all = true
		case "-n", "--no-alias":
			alias = true
		case "-t", "--tree":
			tree = true
		}
	}

	allPorts = AllPorts()
	if !all {
		instPorts = InstPorts()
	}

	recursive("./")
}
