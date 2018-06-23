package main

import (
	"fmt"
	"strings"

	"github.com/go2c/optparse"
	"github.com/onodera-punpun/prt/packages"
	"github.com/onodera-punpun/prt/ports"
)

// depends lists dependencies recursively.
func depends(input []string) error {
	// Define valid arguments.
	o := optparse.New()
	arga := o.Bool("all", 'a', false)
	argn := o.Bool("no-alias", 'n', false)
	argt := o.Bool("tree", 't', false)
	argh := o.Bool("help", 'h', false)

	// Parse arguments.
	_, err := o.Parse(input)
	if err != nil {
		return fmt.Errorf("invaild argument, use `-h` for a list of arguments")
	}

	// Print help.
	if *argh {
		fmt.Println("Usage: prt depends [arguments]")
		fmt.Println("")
		fmt.Println("arguments:")
		fmt.Println("  -a,   --all             also list installed dependencies")
		fmt.Println("  -n,   --no-alias        disable aliasing")
		fmt.Println("  -t,   --tree            list using tree view")
		fmt.Println("  -h,   --help            print help and exit")

		return nil
	}

	// Get all ports.
	all, err := ports.All(config.PrtDir)
	if err != nil {
		return err
	}

	var db packages.Database
	if !*arga {
		// Get installed ports.
		if err := db.Parse(); err != nil {
			return err
		}
	}

	p := ports.New(".")
	if err := p.Pkgfile.Parse(); err != nil {
		return err
	}

	var a [][]ports.Location
	if !*argn {
		a = config.Alias
	}
	if err := p.ParseDepends(a, config.Order, all); err != nil {
		return err
	}

	dependsRecurse(&p, db, 0, *argt)

	return nil
}

var dependsCheck []*ports.Port
var dependsArrow bool

func dependsRecurse(p *ports.Port, db packages.Database, l int, t bool) {
outer:
	for _, d := range p.Depends {
		// Continue if installed.
		for _, n := range db.Packages {
			if d.Pkgfile.Name == n.Name {
				continue outer
			}
		}

		// Continue if already checked.
		for _, c := range dependsCheck {
			if d.Pkgfile.Name == c.Pkgfile.Name {
				if !t {
					continue outer
				}

				if len(d.Pkgfile.Depends) > 0 {
					if !dependsArrow {
						fmt.Printf("%s%s%s\n", dark(strings.Repeat(config.
							IndentChar, l)), d.Pkgfile.Name, dark(" ->"))

						continue outer
					}

					dependsArrow = !dependsArrow
				}
			}
		}
		dependsCheck = append(dependsCheck, d)

		if t {
			fmt.Print(dark(strings.Repeat(config.IndentChar, l)))
		}
		fmt.Println(d.Pkgfile.Name)

		dependsRecurse(d, db, l+1, t)
	}
}
