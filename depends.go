package main

import (
	"fmt"

	"github.com/go2c/optparse"
	"github.com/onodera-punpun/prt/packages"
	"github.com/onodera-punpun/prt/ports"
)

// depends lists dependencies recursively.
func depends(input []string) error {
	// Define valid arguments.
	o := optparse.New()
	arga := o.Bool("all", 'a', false)
	//argn := o.Bool("no-alias", 'n', false)
	//argt := o.Bool("tree", 't', false)
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

	dl, err := p.Pkgfile.RecursiveDepends(config.Alias, config.Order, all)
	if err != nil {
		return err
	}

	for _, d := range dl {
		fmt.Println(d.Pkgfile.Name)
	}

	return nil
}
