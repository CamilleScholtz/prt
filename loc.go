package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/chiyouhen/getopt"
	"github.com/fatih/color"
)

func Loc(args []string) {
	// Initialize opt vars
	var d, n bool

	// Define opts
	shortopts := "hdn"
	longopts := []string{
		"--help",
		"--no-alias",
		"--duplicate",
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
			fmt.Println("Usage: prt loc [arguments] [ports]")
			fmt.Println("")
			fmt.Println("arguments:")
			fmt.Println("  -d,   --duplicate       list duplicate ports as well")
			fmt.Println("  -n,   --no-alias        disable aliasing")
			fmt.Println("  -h,   --help            print help and exit")
			os.Exit(0)
		case "-d", "--duplicate":
			d = true
		case "-n", "--no-alias":
			n = true
		}
	}

	// TODO: Make this happen even when flags are given
	if len(os.Args[2:]) < 1 {
		fmt.Fprintln(os.Stderr, "Please specify a port!")
		os.Exit(1)
	}

	AllPorts = ListAllPorts()

	var checked []string
	var locs []string
	for _, port := range os.Args[2:] {
		// Continue if already checked
		if StringInList(port, checked) {
			continue
		} else {
			checked = append(checked, port)
		}

		iteration := 0

		// Get port location
		locs = GetPortLoc(port)
		if len(locs) < 1 {
			continue
		}
		if !d {
			locs = []string{locs[0]}
		}

		for _, loc := range locs {
			// Alias if needed
			if !n {
				loc = GetPortAlias(loc)
			}

			// Print duplicate indentation
			if d {
				if iteration > 0 {
					color.Set(color.FgBlack, color.Bold)
					fmt.Printf(strings.Repeat(Config.IndentChar, iteration))
					color.Unset()
				}
				iteration += 1
			}

			// Finally print the port :)
			fmt.Println(loc)
		}
	}
}
