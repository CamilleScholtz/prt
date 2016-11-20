package commands

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

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

	// Initialize options
	var description, url, maintainer, depends, optional, version, release bool

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
				description = true
			case "-u", "--url":
				url = true
			case "-m", "--maintainer":
				maintainer = true
			case "-e", "--depends":
				depends = true
			case "-o", "--optional":
				optional = true
			case "-v", "--version":
				version = true
			case "-r", "--release":
				release = true
			}
		}
	} else {
		description = true
		url = true
		maintainer = true
		depends = true
		optional = true
		version = true
		release = true
	}

	// Print stuff
	if description {
		fmt.Println("Description: " + utils.ReadComment(pkgfile, "Description"))
	}
	if url {
		fmt.Println("URL: " + utils.ReadComment(pkgfile, "URL"))
	}
	if maintainer {
		fmt.Println("Maintainer: " + utils.ReadComment(pkgfile, "Maintainer"))
	}
	if depends {
		fmt.Println("Depends on: " + strings.Join(utils.ReadDepends(pkgfile, "Depends on"), " "))
	}
	if optional {
		fmt.Println("Nice to have: " + strings.Join(utils.ReadDepends(pkgfile, "Nice to have|Optional"), " "))
	}
	if version {
		fmt.Println("Version: " + utils.ReadVar(pkgfile, "version"))
	}
	if release {
		fmt.Println("Release: " + utils.ReadVar(pkgfile, "release"))
	}
}
