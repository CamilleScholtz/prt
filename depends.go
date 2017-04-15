package main

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/fatih/color"
	"github.com/go2c/optparse"
)

// depends lists dependencies recursively.
func depends(args []string) {
	// Define valid arguments.
	o := optparse.New()
	arga := o.Bool("all", 'a', false)
	argn := o.Bool("no-alias", 'n', false)
	argt := o.Bool("tree", 't', false)
	argh := o.Bool("help", 'h', false)

	// Parse arguments.
	_, err := o.Parse(args)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Invaild argument, use -h for a list of arguments!")
		os.Exit(1)
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
		os.Exit(0)
	}

	// Get all ports.
	all, err := allPorts()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Get installed ports.
	var inst []string
	if !*arga {
		inst, err = instPorts()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}

	// Recursive loop that prints dependencies.
	var c []string
	var i int
	var recursive func(l string)
	recursive = func(l string) {
		p, err := decodePort(l, "Pkgfile")
		if err != nil {
			printe(err.Error())
			return
		}

		// Get location and dependencies for each port in dependency list.
		for _, p := range p.Pkgfile.Depends {
			// Get port location.
			ll, err := portLoc(all, p)
			if err != nil {
				continue
			}
			l := ll[0]

			// Alias ports if needed.
			if !*argn {
				l = portAlias(l)
			}

			// Continue if port is already installed.
			if !*arga {
				if stringInList(path.Base(l), inst) {
					continue
				}
			}

			// Print tree indentation.
			if *argt {
				// Print tree indentation character for each tree level.
				if i > 0 {
					color.Set(config.DarkColor)
					fmt.Printf(strings.Repeat(config.IndentChar, i))
					color.Unset()
				}

				// Increment tree level if we already checked this port before.
				if !stringInList(p, c) {
					i++
				}
			}

			// Finally print the port.
			fmt.Print(l)

			if stringInList(p, c) {
				// Print "seen before" star if the port has already been
				// checked.
				if *argt {
					color.Set(config.DarkColor)
					fmt.Print(" *")
					color.Unset()
				}

				// Also continue, since we don't want to print each "seen
				// before" port a 100 times.
				fmt.Println()
				continue
			}
			fmt.Println()

			// Append port to checked ports.
			c = append(c, p)

			// Loop.
			recursive(portFullLoc(l))

			// If we end up here, decrement tree level.
			if *argt {
				i--
			}
		}
	}
	recursive("./")
}
