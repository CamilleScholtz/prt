package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/go2c/optparse"
)

// Diff lists outdated packages.
func diff(input []string) {
	// Define valid arguments.
	o := optparse.New()
	argn := o.Bool("no-alias", 'n', false)
	argv := o.Bool("version", 'v', false)
	argh := o.Bool("help", 'h', false)

	// Parse arguments.
	_, err := o.Parse(input)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Invaild argument, use -h for a list of arguments!")
		os.Exit(1)
	}

	// Print help.
	if *argh {
		fmt.Println("Usage: prt diff [arguments]")
		fmt.Println("")
		fmt.Println("arguments:")
		fmt.Println("  -n,   --no-alias        disable aliasing")
		fmt.Println("  -v,   --version         print with version info")
		fmt.Println("  -h,   --help            print help and exit")
		os.Exit(0)
	}

	// Get all ports.
	all, err := ports()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Get installed ports.
	db, err := parseDatabase()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	for i, n := range db.Name {
		pl, err := location(n, all)
		if err != nil {
			continue
		}
		p := pl[0]

		// Alias if needed.
		if !*argn {
			p.alias()
		}

		// Read out Pkgfile.
		if err := p.parsePkgfile(); err != nil {
			printe(err.Error())
			continue
		}

		// Get available version and release from Pkgfile.
		v := p.Pkgfile.Version + "-" + p.Pkgfile.Release

		// Print if installed and available version don't match.
		if v != db.Version[i] {
			fmt.Print(p.Pkgfile.Name)

			// Print version information if needed.
			if *argv {
				fmt.Print(" " + db.Version[i])

				color.Set(config.DarkColor)
				fmt.Print(" -> ")
				color.Unset()

				fmt.Print(v)
			}

			fmt.Println()
		}
	}
}
