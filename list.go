package main

import (
	"fmt"
	"sort"

	"github.com/go2c/optparse"
	"github.com/onodera-punpun/prt/packages"
	"github.com/onodera-punpun/prt/ports"
)

// list lists ports.
func list(input []string) error {
	// Define valid arguments.
	o := optparse.New()
	argi := o.Bool("installed", 'i', false)
	argr := o.Bool("repo", 'r', false)
	argv := o.Bool("version", 'v', false)
	argh := o.Bool("help", 'h', false)

	// Parse arguments.
	_, err := o.Parse(input)
	if err != nil {
		return fmt.Errorf("invaild argument, use `-h` for a list of arguments")
	}

	// Print help.
	if *argh {
		fmt.Println("Usage: prt list [arguments]")
		fmt.Println("")
		fmt.Println("arguments:")
		fmt.Println("  -i,   --installed       list installed ports only")
		fmt.Println("  -r,   --repo            list with repo info")
		fmt.Println("  -v,   --version         list with version info")
		fmt.Println("  -h,   --help            print help and exit")

		return nil
	}

	// Get all ports.
	all, err := ports.All(config.PrtDir)
	if err != nil {
		return err
	}

	var db packages.Database
	if *argi {
		// Get installed ports.
		if err := db.Parse(); err != nil {
			return err
		}

		// Get port locations.
		var pl []ports.Port
		for _, n := range db.Package {
			p, err := ports.Locate(n.Name, config.Order, all)
			if err != nil {
				continue
			}
			pl = append(pl, p[0])
		}

		// I'm using all in the the following for loop, so alias db to all.
		all = pl
	}

	var pl []string
	for i, p := range all {
		var s string

		if !*argr {
			s = p.Location.Port
		} else {
			s = p.Location.Base()
		}

		if *argv {
			if *argi {
				s += " " + db.Package[i].Version
			} else {
				if err := p.Pkgfile.Parse(); err != nil {
					printe(err.Error())
					continue
				}
				s += " " + p.Pkgfile.Version
			}
		}

		pl = append(pl, s)
	}

	sort.Strings(pl)
	for _, p := range pl {
		fmt.Println(p)
	}

	return nil
}
