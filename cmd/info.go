package cmd

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/go2c/optparse"
	"github.com/onodera-punpun/prt/pkgfile"
)

// Info prints port information.
func Info(args []string) {
	// Enable all arguments if the user hasn't specified any.
	var b bool
	if len(args) == 0 {
		b = true
	}

	// Define valid arguments.
	argd := optparse.Bool("description", 'd', b)
	argu := optparse.Bool("url", 'u', b)
	argm := optparse.Bool("maintainer", 'm', b)
	arge := optparse.Bool("depends", 'e', b)
	argo := optparse.Bool("optional", 'o', b)
	argv := optparse.Bool("version", 'v', b)
	argr := optparse.Bool("release", 'r', b)
	argh := optparse.Bool("help", 'h', false)

	// Parse arguments.
	_, err := optparse.Parse(args)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Invaild argument, use -h for a list of arguments!")
		os.Exit(1)
	}

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
	f, err := ioutil.ReadFile("./Pkgfile")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Print info.
	if *argd {
		s, _ := pkgfile.Comment(f, "Description")
		fmt.Println("Description: " + s)
	}
	if *argu {
		s, _ := pkgfile.Comment(f, "URL")
		fmt.Println("URL: " + s)
	}
	if *argm {
		s, _ := pkgfile.Comment(f, "Maintainer")
		fmt.Println("Maintainer: " + s)
	}
	if *arge {
		s, _ := pkgfile.Comment(f, "Depends on")
		fmt.Println("Depends on: " + s)
	}
	if *argo {
		s, _ := pkgfile.Comment(f, "Nice to have|Optional")
		fmt.Println("Nice to have: " + s)
	}
	if *argv {
		s, _ := pkgfile.Var(f, "version")
		fmt.Println("Version: " + s)
	}
	if *argr {
		s, _ := pkgfile.Var(f, "release")
		fmt.Println("Release: " + s)
	}
}
