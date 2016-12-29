package cmd

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/chiyouhen/getopt"
	"github.com/onodera-punpun/prt/pkgfile"
	"github.com/onodera-punpun/prt/utils"
)

// Info prints port information.
func Info(args []string) {
	// Define opts.
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

	// Read out opts.
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
				o = append(o, "a")
			case "-u", "--url":
				o = append(o, "u")
			case "-m", "--maintainer":
				o = append(o, "m")
			case "-e", "--depends":
				o = append(o, "e")
			case "-o", "--optional":
				o = append(o, "0")
			case "-v", "--version":
				o = append(o, "v")
			case "-r", "--release":
				o = append(o, "r")
			}
		}
	} else {
		o = []string{"a", "u", "m", "e", "o", "v", "r"}
	}

	// Read out Pkgfile.
	f, err := ioutil.ReadFile("./Pkgfile")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Print info.
	if utils.StringInList("d", o) {
		s, _ := pkgfile.Comment(f, "Description")
		fmt.Println("Description: " + s)
	}
	if utils.StringInList("u", o) {
		s, _ := pkgfile.Comment(f, "URL")
		fmt.Println("URL: " + s)
	}
	if utils.StringInList("m", o) {
		s, _ := pkgfile.Comment(f, "Maintainer")
		fmt.Println("Maintainer: " + s)
	}
	if utils.StringInList("e", o) {
		s, _ := pkgfile.Comment(f, "Depends on")
		fmt.Println("Depends on: " + s)
	}
	if utils.StringInList("o", o) {
		s, _ := pkgfile.Comment(f, "Nice to have|Optional")
		fmt.Println("Nice to have: " + s)
	}
	if utils.StringInList("v", o) {
		s, _ := pkgfile.Var(f, "version")
		fmt.Println("Version: " + s)
	}
	if utils.StringInList("r", o) {
		s, _ := pkgfile.Var(f, "release")
		fmt.Println("Release: " + s)
	}
}
