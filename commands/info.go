package commands

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/chiyouhen/getopt"
	"github.com/onodera-punpun/prt/utils"
)

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
		fmt.Println("Invaild argument, use -h for a list of arguments.")
		os.Exit(1)
	}

	// Read out Pkgfile
	pkgfile, err := ioutil.ReadFile("./Pkgfile")
	if err != nil {
		fmt.Println("Could not read Pkgfile.")
		os.Exit(1)
	}

	// Initialize opt variables
	var d, u, m, e, o, v, r bool

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

	// Print stuff
	if d {
		fmt.Println("Description: " + utils.ReadComment(pkgfile, "Description"))
	}
	if u {
		fmt.Println("URL: " + utils.ReadComment(pkgfile, "URL"))
	}
	if m {
		fmt.Println("Maintainer: " + utils.ReadComment(pkgfile, "Maintainer"))
	}
	if e {
		fmt.Println("Depends on: " + utils.ReadComment(pkgfile, "Depends on"))
	}
	if o {
		fmt.Println("Nice to have: " + utils.ReadComment(pkgfile, "Nice to have|Optional"))
	}
	if v {
		fmt.Println("Version: " + utils.ReadVar(pkgfile, "version"))
	}
	if r {
		fmt.Println("Release: " + utils.ReadVar(pkgfile, "release"))
	}
}
