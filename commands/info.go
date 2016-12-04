package commands

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/chiyouhen/getopt"
	"github.com/onodera-punpun/prt/pkgfile"
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
		text, err := pkgfile.Comment(f, "Description")
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		fmt.Println("Description: " + text)
	}
	if u {
		text, err := pkgfile.Comment(f, "URL")
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		fmt.Println("URL: " + text)
	}
	if m {
		text, err := pkgfile.Comment(f, "Maintainer")
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		fmt.Println("Maintainer: " + text)
	}
	if e {
		text, err := pkgfile.Comment(f, "Depends on")
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		fmt.Println("Depends on: " + text)
	}
	if o {
		text, err := pkgfile.Comment(f, "Nice to have|Optional")
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		fmt.Println("Nice to have: " + text)
	}
	if v {
		text, err := pkgfile.Comment(f, "version")
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		fmt.Println("Version: " + text)
	}
	if r {
		text, err := pkgfile.Comment(f, "release")
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		fmt.Println("Release: " + text)
	}
}
