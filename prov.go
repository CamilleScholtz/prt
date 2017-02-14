package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"regexp"
	"strings"

	"github.com/fatih/color"
	"github.com/go2c/optparse"
)

// prov searches ports for files.
func prov(args []string) {
	// Define valid arguments.
	o := optparse.New()
	argi := o.Bool("installed", 'i', false)
	argh := o.Bool("help", 'h', false)

	// Parse arguments.
	vals, err := o.Parse(args)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Invaild argument, use -h for a list of arguments!")
		os.Exit(1)
	}

	// Print help.
	if *argh {
		fmt.Println("Usage: prt prov [arguments] [queries]")
		fmt.Println("")
		fmt.Println("arguments:")
		fmt.Println("  -i,   --installed       search in installed ports only")
		fmt.Println("  -h,   --help            print help and exit")
		os.Exit(0)
	}

	// This command needs a value.
	if len(vals) == 0 {
		fmt.Fprintln(os.Stderr, "Please specify a query!")
		os.Exit(1)
	}

	for _, v := range vals {
		r, err := regexp.Compile(v)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		// TODO: Use Alias and Loc here to always display repo info?
		if *argi {
			// Read out pkg db.
			db, err := os.Open("/var/lib/pkg/db")
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				continue
			}
			s := bufio.NewScanner(db)

			// Search for files.
			var b bool
			var n string
			var ll [][]string
			for s.Scan() {
				if b {
					n = s.Text()
					b = false
				} else if s.Text() == "" {
					b = true
				} else if r.MatchString(s.Text()) {
					ll = append(ll, []string{n, s.Text()})
				}
			}

			var on string
			for _, l := range ll {
				// Print port name.
				if on != l[0] {
					fmt.Println(l[0])
				}

				// Print matched files.
				color.Set(config.DarkColor)
				fmt.Print(config.IndentChar)
				color.Unset()
				fmt.Println(l[1])

				on = l[0]
			}

			db.Close()
		} else {
			// Get all ports.
			all, err := allPorts()
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				continue
			}

			for _, p := range all {
				// Read out Pkgfile.
				f, err := os.Open(path.Join(portFullLoc(p), ".footprint"))
				if err != nil {
					fmt.Fprintln(os.Stderr, err)
					continue
				}
				s := bufio.NewScanner(f)

				// Search for files.
				var ll []string
				for s.Scan() {
					if r.MatchString(s.Text()) {
						ll = append(ll, s.Text())
					}
				}

				// Print port name.
				if len(ll) > 0 {
					fmt.Println(p)
				}

				// Print matched files.
				for _, l := range ll {
					color.Set(config.DarkColor)
					fmt.Print(config.IndentChar)
					color.Unset()
					fmt.Println(strings.Split(l, "\t")[2])
				}

				f.Close()
			}
		}
	}
}
