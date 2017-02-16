package main

import (
	"fmt"
	"os"
	"path"
	"sort"

	"github.com/go2c/optparse"
)

// list lists ports.
func list(args []string) {
	// Define valid arguments.
	o := optparse.New()
	argi := o.Bool("installed", 'i', false)
	argr := o.Bool("repo", 'r', false)
	argv := o.Bool("version", 'v', false)
	argh := o.Bool("help", 'h', false)

	// Parse arguments.
	_, err := o.Parse(args)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Invaild argument, use -h for a list of arguments!")
		os.Exit(1)
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
		os.Exit(0)
	}

	// Get all ports.
	all, err := allPorts()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	var instv []string
	if *argi {
		// Get installed ports.
		inst, err := instPorts()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		// Get installed port versions if needed.
		if *argv {
			instv, err = instVersPorts()
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
		}

		// Get port locations if needed.
		if *argr {
			for i, p := range inst {
				ll, err := portLoc(all, p)
				if err != nil {
					fmt.Fprintln(os.Stderr, err)
					continue
				}
				inst[i] = ll[0]
			}
		}

		// We want pretty output, so sort.
		sort.Strings(inst)

		// I'm using all in the the following for loop, so alias inst to all.
		all = inst
	}

	// TODO: Use l instead of p here?
	for i, p := range all {
		if *argv {
			var v string
			if *argi {
				// Get installed version.
				v = instv[i]
			} else {
				// Read out Pkgfile.
				f, err := readPkgfile(portFullLoc(p))
				if err != nil {
					fmt.Fprintln(os.Stderr, err)
					continue
				}

				// Get available version from Pkgfile.
				v, err = f.variable("version")
				if err != nil {
					fmt.Fprintln(os.Stderr, err)
					continue
				}
			}

			// Merge port and version.
			p += " " + v
		}

		// Remove repo if needed.
		if !*argr && !*argi {
			p = path.Base(p)
		}

		fmt.Println(p)
	}
}
