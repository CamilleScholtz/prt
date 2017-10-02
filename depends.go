package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/go2c/optparse"
)

// depends lists dependencies recursively.
func depends(input []string) {
	// Define valid arguments.
	o := optparse.New()
	arga := o.Bool("all", 'a', false)
	argn := o.Bool("no-alias", 'n', false)
	argt := o.Bool("tree", 't', false)
	argh := o.Bool("help", 'h', false)

	// Parse arguments.
	_, err := o.Parse(input)
	if err != nil {
		fmt.Fprintln(os.Stderr,
			"Invaild argument, use -h for a list of arguments!")
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
	all, err := ports()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Get installed ports.
	var db database
	if !*arga {
		db, err = parseDatabase()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}

	var p port
	p.Location = "."
	if err := p.parsePkgfile(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	p.depends(!*argn, all)

	var i int
	var c []string
	pl := p.Depends
	var recursive func()
	recursive = func() {
		for _, p := range pl {
			if !*arga {
				if !stringInList(p.Pkgfile.Name, db.Name) {
					if !stringInList(p.Pkgfile.Name, c) {
						if *argt {
							color.Set(config.DarkColor)
							fmt.Printf(strings.Repeat(config.IndentChar, i))
							color.Unset()
						}
						fmt.Println(p.getBaseDir())

						// Append to printed ports.
						c = append(c, p.Pkgfile.Name)
					}
				}
			} else {
				if *argt {
					color.Set(config.DarkColor)
					fmt.Printf(strings.Repeat(config.IndentChar, i))
					color.Unset()
				}
				fmt.Println(p.getBaseDir())
			}

			i++
			if len(p.Depends) > 0 {
				pl = p.Depends
				recursive()
			}
			i--
		}
	}
	recursive()
}
