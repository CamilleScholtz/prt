package cmd

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/chiyouhen/getopt"
	"github.com/onodera-punpun/prt/pkgfile"
)

// Info prints port information.
func Info(args []string) {
	// Define allowed opts.
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

	type optStruct struct {
		d bool
		u bool
		m bool
		e bool
		o bool
		v bool
		r bool
	}

	var opt optStruct
	for _, o := range opts {
		switch o[0] {
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
			opt.d = true
		case "-u", "--url":
			opt.u = true
		case "-m", "--maintainer":
			opt.m = true
		case "-e", "--depends":
			opt.e = true
		case "-o", "--optional":
			opt.o = true
		case "-v", "--version":
			opt.v = true
		case "-r", "--release":
			opt.r = true
		}
	}

	// Enable all opts if the user hasn't specified any.
	if len(opts) == 0 {
		opt.d = true
		opt.u = true
		opt.m = true
		opt.e = true
		opt.o = true
		opt.v = true
		opt.r = true
	}

	// Read out Pkgfile.
	f, err := ioutil.ReadFile("./Pkgfile")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Print info.
	if opt.d {
		s, _ := pkgfile.Comment(f, "Description")
		fmt.Println("Description: " + s)
	}
	if opt.u {
		s, _ := pkgfile.Comment(f, "URL")
		fmt.Println("URL: " + s)
	}
	if opt.m {
		s, _ := pkgfile.Comment(f, "Maintainer")
		fmt.Println("Maintainer: " + s)
	}
	if opt.e {
		s, _ := pkgfile.Comment(f, "Depends on")
		fmt.Println("Depends on: " + s)
	}
	if opt.o {
		s, _ := pkgfile.Comment(f, "Nice to have|Optional")
		fmt.Println("Nice to have: " + s)
	}
	if opt.v {
		s, _ := pkgfile.Var(f, "version")
		fmt.Println("Version: " + s)
	}
	if opt.r {
		s, _ := pkgfile.Var(f, "release")
		fmt.Println("Release: " + s)
	}
}
