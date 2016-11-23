package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/chiyouhen/getopt"
)

// Initialize opt variables
var D, U, M, E, O, V, R bool

func Info(args []string) {
	// Define opts
	shortopts := "hdumeovr"
	longopts := []string{
		"--help",
		"--description",
		"--url",
		"--maintainer",
		"--depends",
		"--optional",
		"--version",
		"--release",
	}

	// Read out opts
	opts, _, err := getopt.Getopt(args, shortopts, longopts)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Invaild argument, use -h for a list of arguments!")
		os.Exit(1)
	}

	// Read out Pkgfile
	pkgfile, err := ioutil.ReadFile("./Pkgfile")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Could not read Pkgfile!")
		os.Exit(1)
	}

	if len(opts) > 0 {
		for _, opt := range opts {
			switch opt[0] {
			case "-h", "--help":
				fmt.Println("Usage: prt info [arguments]")
				fmt.Println("")
				fmt.Println("arguments:")
				fmt.Println("  -d,   --description     print description")
				fmt.Println("  -u,   --url             print url")
				fmt.Println("  -m,   --maintainer      print maintainer")
				fmt.Println("  -e,   --depends         print dependencies")
				fmt.Println("  -o,   --optional        print optional dependencies")
				fmt.Println("  -v,   --version         print version")
				fmt.Println("  -r,   --release         print release")
				fmt.Println("  -h,   --help            print help and exit")
				os.Exit(0)
			case "-d", "--description":
				D = true
			case "-u", "--url":
				U = true
			case "-m", "--maintainer":
				M = true
			case "-e", "--depends":
				D = true
			case "-o", "--optional":
				O = true
			case "-v", "--version":
				V = true
			case "-r", "--release":
				R = true
			}
		}
	} else {
		D, U, M, E, O, V, R = true, true, true, true, true, true, true
	}

	// Print stuff
	if D {
		fmt.Println("Description: " + ReadComment(pkgfile, "Description"))
	}
	if U {
		fmt.Println("URL: " + ReadComment(pkgfile, "URL"))
	}
	if M {
		fmt.Println("Maintainer: " + ReadComment(pkgfile, "Maintainer"))
	}
	if E {
		fmt.Println("Depends on: " + ReadComment(pkgfile, "Depends on"))
	}
	if O {
		fmt.Println("Nice to have: " + ReadComment(pkgfile, "Nice to have|Optional"))
	}
	if V {
		fmt.Println("Version: " + ReadVar(pkgfile, "version"))
	}
	if R {
		fmt.Println("Release: " + ReadVar(pkgfile, "release"))
	}
}
