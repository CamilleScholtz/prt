package main

import (
	"fmt"
	"strings"

	"github.com/go2c/optparse"
)

// info prints port information.
func info(input []string) error {
	// Enable all arguments if the user hasn't specified any.
	var b bool
	if len(input) == 0 {
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
	args := o.Bool("source", 's', b)
	argh := o.Bool("help", 'h', false)

	// Parse arguments.
	_, err := o.Parse(input)
	if err != nil {
		return fmt.Errorf("invaild argument, use `-h` for a list of arguments")
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
		fmt.Println("  -s,   --source          print sources")
		fmt.Println("  -h,   --help            print help and exit")

		return nil
	}

	// Read out Pkgfile.
	p := newPort(".")
	if err := p.parsePkgfile(true); err != nil {
		return err
	}

	// Print info from Pkgfile.
	if *argd {
		fmt.Println("Description: " + p.Pkgfile.Description)
	}
	if *argu {
		fmt.Println("URL: " + p.Pkgfile.URL)
	}
	if *argm {
		fmt.Println("Maintainer: " + p.Pkgfile.Maintainer)
	}
	if *arge {
		fmt.Println("Depends on: " + strings.Join(p.Pkgfile.Depends, ", "))
	}
	if *argo {
		fmt.Println("Nice to have: " + strings.Join(p.Pkgfile.Optional, ", "))
	}
	if *argv {
		fmt.Println("Version: " + p.Pkgfile.Version)
	}
	if *argr {
		fmt.Println("Release: " + p.Pkgfile.Release)
	}
	if *args {
		fmt.Println("Source: " + strings.Join(p.Pkgfile.Source, ", "))
	}

	return nil
}
