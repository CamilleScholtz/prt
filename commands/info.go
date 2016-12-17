package commands

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/chiyouhen/getopt"
	"github.com/onodera-punpun/prt/pkgfile"
	"github.com/onodera-punpun/prt/utils"
)

// Info prints ports information
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

	// Read out Pkgfile
	f, err := ioutil.ReadFile("./Pkgfile")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Could not read Pkgfile!")
		os.Exit(1)
	}

	// Print stuff
	if utils.StringInList("d", o) {
		text, err := pkgfile.Comment(f, "Description")
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		fmt.Println("Description: " + text)
	}
	if utils.StringInList("u", o) {
		text, err := pkgfile.Comment(f, "URL")
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		fmt.Println("URL: " + text)
	}
	if utils.StringInList("m", o) {
		text, err := pkgfile.Comment(f, "Maintainer")
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		fmt.Println("Maintainer: " + text)
	}
	if utils.StringInList("e", o) {
		text, err := pkgfile.Comment(f, "Depends on")
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		fmt.Println("Depends on: " + text)
	}
	if utils.StringInList("o", o) {
		text, err := pkgfile.Comment(f, "Nice to have|Optional")
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		fmt.Println("Nice to have: " + text)
	}
	if utils.StringInList("v", o) {
		text, err := pkgfile.Comment(f, "version")
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		fmt.Println("Version: " + text)
	}
	if utils.StringInList("r", o) {
		text, err := pkgfile.Comment(f, "release")
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		fmt.Println("Release: " + text)
	}
}
