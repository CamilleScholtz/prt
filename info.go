package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/chiyouhen/getopt"
)

func Info(args []string) {
	// Initialize opt vars
	var d, u, m, e, o, v, r bool

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
				d = true
			case "-u", "--url":
				u = true
			case "-m", "--maintainer":
				m = true
			case "-e", "--depends":
				d = true
			case "-o", "--optional":
				o = true
			case "-v", "--version":
				v = true
			case "-r", "--release":
				r = true
			}
		}
	} else {
		d, u, m, e, o, v, r = true, true, true, true, true, true, true
	}

	// Read out Pkgfile
	f, err := ioutil.ReadFile("./Pkgfile")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Could not read Pkgfile!")
		os.Exit(1)
	}

	// Print stuff
	if d {
		fmt.Println("Description: " + ReadComment(f, "Description"))
	}
	if u {
		fmt.Println("URL: " + ReadComment(f, "URL"))
	}
	if m {
		fmt.Println("Maintainer: " + ReadComment(f, "Maintainer"))
	}
	if e {
		fmt.Println("Depends on: " + ReadComment(f, "Depends on"))
	}
	if o {
		fmt.Println("Nice to have: " + ReadComment(f, "Nice to have|Optional"))
	}
	if v {
		fmt.Println("Version: " + ReadVar(f, "version"))
	}
	if r {
		fmt.Println("Release: " + ReadVar(f, "release"))
	}
}
