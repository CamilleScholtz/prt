package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/fatih/color"
	"github.com/go2c/optparse"
	"github.com/onodera-punpun/prt/config"
	"github.com/onodera-punpun/prt/pkgfile"
	"github.com/onodera-punpun/prt/ports"
	"github.com/onodera-punpun/prt/utils"
)

// Depends lists dependencies recursively.
func Depends(args []string) {
	// Load config.
	conf := config.Load()

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
	all, err := ports.All()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Get installed ports.
	var inst []string
	if !*arga {
		inst, err = ports.Inst()
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
		// Read out Pkgfile.
		f, err := ioutil.ReadFile(path.Join(l, "Pkgfile"))
		if err != nil {
			utils.Printe(err.Error())
			return
		}

		// Get dependencies.
		dl, err := pkgfile.Depends(f, "Depends on")
		if err != nil {
			return
		}

		for _, p := range dl {
			// Get port location.
			ll, err := ports.Loc(all, p)
			if err != nil {
				continue
			}
			l := ll[0]

			// Alias ports if needed.
			if !*argn {
				l = ports.Alias(l)
			}

			// Continue if port is already installed.
			if !*arga {
				if utils.StringInList(path.Base(l), inst) {
					continue
				}
			}

			// Print tree indentation.
			if *argt {
				if i > 0 {
					color.Set(conf.DarkColor)
					fmt.Printf(strings.Repeat(conf.IndentChar, i))
					color.Unset()
				}
				if !utils.StringInList(p, c) {
					i++
				}
			}

			// Finally print the port.
			fmt.Print(l)

			// Print "seen before" star if already checked.
			if utils.StringInList(p, c) {
				color.Set(conf.DarkColor)
				fmt.Println(" *")
				color.Unset()

				continue
			}
			fmt.Println()

			// Add to checked ports.
			c = append(c, p)

			// Loop.
			recursive(ports.FullLoc(l))

			// If we end up here, remove one tree indentation level
			if *argt {
				i--
			}
		}
	}
	recursive("./")
}
