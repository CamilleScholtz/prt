package cmd

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"regexp"
	"strings"

	"github.com/fatih/color"
	"github.com/go2c/optparse"
	"github.com/onodera-punpun/prt/config"
	"github.com/onodera-punpun/prt/ports"
)

// Prov searches ports for files.
func Prov(args []string) {
	// Load config.
	conf := config.Load()

	// Define valid arguments.
	argi := optparse.Bool("installed", 'i', false)
	argh := optparse.Bool("help", 'h', false)

	// Parse arguments.
	vals, err := optparse.Parse(args)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Invaild argument, use -h for a list of arguments!")
		os.Exit(1)
	}

	if *argh {
		fmt.Println("Usage: prt print [arguments] [queries]")
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
				color.Set(conf.DarkColor)
				fmt.Print(conf.IndentChar)
				color.Unset()
				fmt.Println(l[1])

				on = l[0]
			}

			db.Close()
		} else {
			// Get all ports.
			all, err := ports.All()
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				continue
			}

			for _, p := range all {
				// Read out Pkgfile.
				f, err := os.Open(path.Join(ports.FullLoc(p), ".footprint"))
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
					color.Set(conf.DarkColor)
					fmt.Print(conf.IndentChar)
					color.Unset()
					fmt.Println(strings.Split(l, "\t")[2])
				}

				f.Close()
			}
		}
	}
}
