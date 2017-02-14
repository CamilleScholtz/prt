package main

import (
	"fmt"
	"os"

	"github.com/go2c/optparse"
)

// info prints port information.
func info(args []string) {
	// Enable all arguments if the user hasn't specified any.
	var b bool
	if len(args) == 0 {
		b = true
	}

	// Define valid arguments.
	o := optparse.New()
	argd := o.Bool("description", 'd', b)
	argu := o.Bool("url", 'u', b)
	argm := o.Bool("maintainer", 'm', b)
	arge := o.Bool("depends", 'e', b)
	argo := o.Bool("optional", 'o', b)
	argv := o.Bool("version", 'v', b)
	argr := o.Bool("release", 'r', b)
	argh := o.Bool("help", 'h', false)

	// Parse arguments.
	_, err := o.Parse(args)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Invaild argument, use -h for a list of arguments!")
		os.Exit(1)
	}

	// Print help.
	if *argh {
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
	}

	// Read out Pkgfile.
	f, err := readPkgfile("./Pkgfile")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Print info from Pkgfile..
	if *argd {
		s, _ := f.comment("Description")
		fmt.Println("Description: " + s)
	}
	if *argu {
		s, _ := f.comment("URL")
		fmt.Println("URL: " + s)
	}
	if *argm {
		s, _ := f.comment("Maintainer")
		fmt.Println("Maintainer: " + s)
	}
	if *arge {
		s, _ := f.comment("Depends on")
		fmt.Println("Depends on: " + s)
	}
	if *argo {
		s, _ := f.comment("Nice to have|Optional")
		fmt.Println("Nice to have: " + s)
	}
	if *argv {
		s, _ := f.variable("version")
		fmt.Println("Version: " + s)
	}
	if *argr {
		s, _ := f.variable("release")
		fmt.Println("Release: " + s)
	}
}
