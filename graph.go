package main

import (
	"fmt"
	"os"

	"github.com/go2c/optparse"
	"github.com/paetzke/godot/godot"
)

// graph generates a dependency grap.
func graph(input []string) {
	// Define valid arguments.
	o := optparse.New()
	argn := o.Bool("no-alias", 'n', false)
	argh := o.Bool("help", 'h', false)

	// Parse arguments.
	vals, err := o.Parse(input)
	if err != nil {
		fmt.Fprintln(os.Stderr,
			"Invaild argument, use -h for a list of arguments!")
		os.Exit(1)
	}

	// Print help.
	if *argh {
		fmt.Println("Usage: prt graph [arguments] [location]")
		fmt.Println("")
		fmt.Println("arguments:")
		fmt.Println("  -n,   --no-alias        disable aliasing")
		fmt.Println("  -h,   --help            print help and exit")
		os.Exit(0)
	}

	// Get all ports.
	all, err := ports()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	var p port
	p.Location = "."
	if err := p.parsePkgfile(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	p.getDepends(!*argn, all)

	var f string
	if len(vals) > 0 {
		f = vals[0]
	} else {
		f = p.Pkgfile.Name + ".svg"
	}

	dot, err := godot.NewDotter("svg", "digraph", f)
	defer dot.Close()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	op := p.getBaseDir()
	pl := p.Depends
	var recursive func()
	recursive = func() {
		for _, p := range pl {
			dot.SetLink(op, p.getBaseDir())
			dot.SetLabel(p.getBaseDir(), p.getBaseDir())

			if len(p.Depends) > 0 {
				op = p.getBaseDir()
				pl = p.Depends
				recursive()
			}
		}
	}
	recursive()

	if err := dot.Close(); err != nil {
		fmt.Fprintln(os.Stderr, "Could not generate graph!")
		os.Exit(1)
	}
}
