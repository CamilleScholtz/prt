package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/chiyouhen/getopt"
	"github.com/fatih/color"
	"github.com/onodera-punpun/prt/config"
	"github.com/onodera-punpun/prt/pkgfile"
	"github.com/onodera-punpun/prt/ports"
	"github.com/onodera-punpun/prt/utils"
)

// Depends lists dependencies recursively.
func Depends(args []string) {
	// Load config.
	var conf = config.Load()

	// Define allowed opts.
	shortopts := "hant"
	longopts := []string{
		"--help",
		"--no-alias",
		"--tree",
	}

	// Read out opts.
	opts, _, err := getopt.Getopt(args, shortopts, longopts)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Invaild argument, use -h for a list of arguments!")
		os.Exit(1)
	}

	type optStruct struct {
		a bool
		n bool
		t bool
	}

	var opt optStruct
	for _, o := range opts {
		switch o[0] {
		case "-h", "--help":
			fmt.Println("Usage: prt depends [arguments]")
			fmt.Println("")
			fmt.Println("arguments:")
			fmt.Println("  -a,   --all             also list installed dependencies")
			fmt.Println("  -n,   --no-alias        disable aliasing")
			fmt.Println("  -t,   --tree            list using tree view")
			fmt.Println("  -h,   --help            print help and exit")
			os.Exit(0)
		case "-a", "--all":
			opt.a = true
		case "-n", "--no-alias":
			opt.n = true
		case "-t", "--tree":
			opt.t = true
		}
	}

	// Get all ports.
	all, err := ports.All()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Get installed ports.
	var inst []string
	if !opt.a {
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
			// Continue if already checked.
			if utils.StringInList(p, c) {
				continue
			}
			// Add to checked ports.
			c = append(c, p)

			// Get port location.
			ll, err := ports.Loc(all, p)
			if err != nil {
				continue
			}
			l := ll[0]

			// Alias ports if needed.
			if !opt.n {
				l = ports.Alias(l)
			}

			// Continue port is already installed.
			if !opt.a {
				if utils.StringInList(path.Base(l), inst) {
					continue
				}

				// Core packages should always be installed.
				if path.Dir(l) == "core" {
					continue
				}
			}

			// Print tree indentation.
			if opt.t {
				if i > 0 {
					color.Set(conf.DarkColor)
					fmt.Printf(strings.Repeat(conf.IndentChar, i))
					color.Unset()
				}
				i++
			}

			// Finally print the port.
			fmt.Println(l)

			// Loop.
			recursive(ports.FullLoc(l))

			// If we end up here, remove one tree indentation level
			if opt.t {
				i--
			}
		}
	}

	recursive("./")
}
