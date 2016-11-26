package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/chiyouhen/getopt"
	"github.com/fatih/color"
)

// Initialize variables
var all, alias, tree bool
var allPorts, checkPorts, instPorts []string
var i int

func Loc(args []string) {
	// Initialize variables
	var dup, alias bool
	var allPorts, checkPorts, instPorts []string

	// Define opts
	shortopts := "hdn"
	longopts := []string{
		"--help",
		"--no-alias",
		"--duplicate",
	}

	// Read out opts
	opts, vals, err := getopt.Getopt(args, shortopts, longopts)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Invaild argument, use -h for a list of arguments!")
		os.Exit(1)
	}

	for _, opt := range opts {
		switch opt[0] {
		case "-h", "--help":
			fmt.Println("Usage: prt loc [arguments] [ports]")
			fmt.Println("")
			fmt.Println("arguments:")
			fmt.Println("  -d,   --duplicate       list duplicate ports as well")
			fmt.Println("  -n,   --no-alias        disable aliasing")
			fmt.Println("  -h,   --help            print help and exit")
			os.Exit(0)
		case "-d", "--duplicate":
			dup = true
		case "-n", "--no-alias":
			alias = true
		}
	}

	if len(vals) == 1 {
		fmt.Fprintln(os.Stderr, "Please specify a port!")
		os.Exit(1)
	}

	allPorts = AllPorts()

	var locs []string
	for _, port := range args {
		// Continue if already checked
		if StringInList(port, checkPorts) {
			continue
		}
		checkPorts = append(checkPorts, port)

		i := 0

		// Get port location
		locs = PortLoc(port)
		if len(locs) < 1 {
			continue
		}
		if !dup {
			locs = []string{locs[0]}
		}

		for _, loc := range locs {
			// Alias if needed
			if !alias {
				loc = PortAlias(loc)
			}

			// Print duplicate indentation
			if dup {
				if i > 0 {
					color.Set(color.FgBlack, color.Bold)
					fmt.Printf(strings.Repeat(Config.IndentChar, i))
					color.Unset()
				}
				i += 1
			}

			// Finally print the port :)
			fmt.Println(loc)
		}
	}
}
