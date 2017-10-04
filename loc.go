package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/go2c/optparse"
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
		return fmt.Errorf("invaild argument, use -h for a list of arguments")
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
	all, err := ports()
	if err != nil {
		return err
	}

	var c []string
	var i int
	for _, n := range vals {
		// Continue if already checked.
		if stringInList(n, c) {
			continue
		}

		// Add to checked ports.
		c = append(c, n)

		// Get port location.
		pl, err := location(n, all)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Port not found in the ports tree!")
			continue
		}
		if !*argd {
			pl = []port{pl[0]}
		}

		var op string
		for _, p := range pl {
			// Alias if needed.
			if !*argn {
				p.alias()
			}

			// Print duplicate indentation.
			if *argd {
				// Reset indentation on new port
				if p.getPortDir() != op {
					i = 0
				}
				op = p.getPortDir()

				if i > 0 {
					color.Set(config.DarkColor)
					fmt.Printf(strings.Repeat(config.IndentChar, i))
					color.Unset()
				}
				i++
			}

			// Finally print the port.
			fmt.Println(p.getBaseDir())
		}
	}

	return nil
}
