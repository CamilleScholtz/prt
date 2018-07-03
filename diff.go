package main

import (
	"fmt"
	"os"

	"github.com/go2c/optparse"
	"github.com/onodera-punpun/prt/ports"
)

// DiffCommand lists outdated packages.
func diffCommand(input []string) error {
	// Define valid arguments.
	o := optparse.New()
	argn := o.Bool("no-alias", 'n', false)
	argv := o.Bool("version", 'v', false)
	argh := o.Bool("help", 'h', false)

	// Parse arguments.
	_, err := o.Parse(input)
	if err != nil {
		return fmt.Errorf("invaild argument, use `-h` for a list of arguments")
	}

	// Print help.
	if *argh {
		fmt.Println("Usage: prt diff [arguments]")
		fmt.Println("")
		fmt.Println("arguments:")
		fmt.Println("  -n,   --no-alias        disable aliasing")
		fmt.Println("  -v,   --version         print with version info")
		fmt.Println("  -h,   --help            print help and exit")

		return nil
	}

	// Get all ports.
	all, err := ports.All()
	if err != nil {
		return err
	}

	// Get installed ports.
	var db ports.Database
	if err := db.Parse(); err != nil {
		return err
	}

	for _, n := range db.Packages {
		pl, err := ports.Locate(all, n.Name)
		if err != nil {
			continue
		}
		p := pl[0]

		// Alias if needed.
		if !*argn {
			p.Alias()
		}

		// Read out Pkgfile.
		if err := p.Pkgfile.Parse(); err != nil {
			fmt.Fprintf(os.Stderr, "%s%s\n", warning(config.WarningChar), err)
			continue
		}

		// Get available version and release from Pkgfile.
		v := p.Pkgfile.Version + "-" + p.Pkgfile.Release

		// Print if installed and available version don't match.
		if v != n.Version {
			fmt.Print(p.Pkgfile.Name)

			// Print version information if needed.
			if *argv {
				fmt.Printf(" %s%s%s", n.Version, dark(" -> "), v)
			}

			fmt.Println()
		}
	}

	return nil
}
