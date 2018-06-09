package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/go2c/optparse"
)

// prov searches ports for files.
func prov(input []string) error {
	// Define valid arguments.
	o := optparse.New()
	argi := o.Bool("installed", 'i', false)
	argh := o.Bool("help", 'h', false)

	// Parse arguments.
	vals, err := o.Parse(input)
	if err != nil {
		return fmt.Errorf("invaild argument, use `-h` for a list of arguments")
	}

	// Print help.
	if *argh {
		fmt.Println("Usage: prt prov [arguments] [queries]")
		fmt.Println("")
		fmt.Println("arguments:")
		fmt.Println("  -i,   --installed       search in installed ports only")
		fmt.Println("  -h,   --help            print help and exit")

		return nil
	}

	// This command needs a value.
	if len(vals) == 0 {
		return fmt.Errorf("please specify a query")
	}

	for _, v := range vals {
		// TODO: Use Alias and Location here to always display repo info?
		// TODO: Can I somehow use db.go for this? I should right?
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
			var fl [][]string
			for s.Scan() {
				if b {
					n = s.Text()
					b = false
				} else if s.Text() == "" {
					b = true
				} else if strings.Contains(s.Text(), v) {
					fl = append(fl, []string{n, s.Text()})
				}
			}

			var on string
			for _, f := range fl {
				// Print port location.
				if on != f[0] {
					fmt.Println(f[0])
				}

				// Print matched files.
				color.Set(config.DarkColor)
				fmt.Print(config.IndentChar)
				color.Unset()
				fmt.Println(f[1])

				on = f[0]
			}

			db.Close()
		} else {
			// Get all ports.
			all, err := ports()
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				continue
			}

			for _, p := range all {
				if err := p.parseFootprint(); err != nil {
					printe(err.Error())
					continue
				}

				// Search for files.
				var fl []string
				for _, f := range p.Footprint.File {
					if strings.Contains(f, v) {
						fl = append(fl, f)
					}
				}

				// Print port location.
				if len(fl) > 0 {
					fmt.Println(p.Location.base())
				}

				// Print matched files.
				for _, f := range fl {
					color.Set(config.DarkColor)
					fmt.Print(config.IndentChar)
					color.Unset()
					fmt.Println(f)
				}
			}
		}
	}

	return nil
}
