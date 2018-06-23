package main

import (
	"fmt"
	"strings"

	"github.com/go2c/optparse"
	"github.com/onodera-punpun/go-utils/array"
	"github.com/onodera-punpun/prt/ports"
)

// loc prints port locations
func loc(input []string) error {
	// Define valid arguments.
	o := optparse.New()
	argd := o.Bool("duplicate", 'd', false)
	argn := o.Bool("no-alias", 'n', false)
	argh := o.Bool("help", 'h', false)

	// Parse arguments.
	vals, err := o.Parse(input)
	if err != nil {
		return fmt.Errorf("invaild argument, use `-h` for a list of arguments")
	}

	// Print help.
	if *argh {
		fmt.Println("Usage: prt loc [arguments] [ports]")
		fmt.Println("")
		fmt.Println("arguments:")
		fmt.Println("  -d,   --duplicate       list duplicate ports as well")
		fmt.Println("  -n,   --no-alias        disable aliasing")
		fmt.Println("  -h,   --help            print help and exit")

		return nil
	}

	// This command needs a value.
	if len(vals) == 0 {
		return fmt.Errorf("please specify a port")
	}

	// Get all ports.
	all, err := ports.All(config.PrtDir)
	if err != nil {
		return err
	}

	var check []string
	var i int
	for _, n := range vals {
		// Continue if already checked.
		if array.ContainsString(check, n) {
			continue
		}

		// Add to checked ports.
		check = append(check, n)

		// Get port location.
		pl, err := ports.Locate(n, config.Order, all)
		if err != nil {
			return err
		}
		if !*argd {
			pl = []ports.Port{pl[0]}
		}

		var op string
		for _, p := range pl {
			// Alias if needed.
			if !*argn {
				p.Alias(config.Alias)
			}

			// Print duplicate indentation.
			if *argd {
				// Reset indentation on new port
				if p.Location.Port != op {
					i = 0
				}
				op = p.Location.Port

				if i > 0 {
					fmt.Print(dark(strings.Repeat(config.IndentChar, i)))
				}
				i++
			}

			// Finally print the port.
			fmt.Println(p.Location.Base())
		}
	}

	return nil
}
