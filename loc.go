package main

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/fatih/color"
	"github.com/go2c/optparse"
)

// loc prints port locations
func loc(args []string) {
	// Define valid arguments.
	o := optparse.New()
	argd := o.Bool("duplicate", 'd', false)
	argn := o.Bool("no-alias", 'n', false)
	argh := o.Bool("help", 'h', false)

	// Parse arguments.
	vals, err := o.Parse(args)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Invaild argument, use -h for a list of arguments!")
		os.Exit(1)
	}

	// Print help.
	if *argh {
		fmt.Println("Usage: prt loc [arguments] [ports]")
		fmt.Println("")
		fmt.Println("arguments:")
		fmt.Println("  -d,   --duplicate       list duplicate ports as well")
		fmt.Println("  -n,   --no-alias        disable aliasing")
		fmt.Println("  -h,   --help            print help and exit")
		os.Exit(0)
	}

	// This command needs a value.
	if len(vals) == 0 {
		fmt.Fprintln(os.Stderr, "Please specify a port!")
		os.Exit(1)
	}

	// Get all ports.
	all, err := allPorts()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	var c []string
	var i int
	for _, p := range vals {
		// Continue if already checked.
		if stringInList(p, c) {
			continue
		}
		// Add to checked ports.
		c = append(c, p)

		// Get port location.
		ll, err := portLoc(all, p)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}
		if !*argd {
			ll = []string{ll[0]}
		}

		var op string
		for _, l := range ll {
			// Alias if needed.
			if !*argn {
				l = portAlias(l)
			}

			// Print duplicate indentation.
			if *argd {
				// Reset indentation on new port
				if path.Base(l) != op {
					i = 0
				}
				op = path.Base(l)

				if i > 0 {
					color.Set(config.DarkColor)
					fmt.Printf(strings.Repeat(config.IndentChar, i))
					color.Unset()
				}
				i++
			}

			// Finally print the port.
			fmt.Println(l)
		}
	}
}
