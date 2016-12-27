package commands

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/chiyouhen/getopt"
	"github.com/fatih/color"
	"github.com/onodera-punpun/prt/ports"
	"github.com/onodera-punpun/prt/utils"
)

// Prov searches ports for files.
func Prov(args []string) {
	// Define opts.
	shortopts := "hi"
	longopts := []string{
		"--help",
		"--installed",
	}

	// Read out opts.
	opts, vals, err := getopt.Getopt(args, shortopts, longopts)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Invaild argument, use -h for a list of arguments!")
		os.Exit(1)
	}

	for _, opt := range opts {
		switch opt[0] {
		case "-h", "--help":
			fmt.Println("Usage: prt print [arguments] [queries]")
			fmt.Println("")
			fmt.Println("arguments:")
			fmt.Println("  -i,   --installed       search in installed ports only")
			fmt.Println("  -h,   --help            print help and exit")
			os.Exit(0)
		case "-i", "--installed":
			o = append(o, "i")
		}
	}

	// This command needs a value.
	if len(vals) == 0 {
		fmt.Fprintln(os.Stderr, "Please specify a query!")
		os.Exit(1)
	}

	for _, v := range vals {
		// Evaluate regex.
		r, err := regexp.Compile(v)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		// TODO: Use Alias and Loc here to always display repo info?
		if utils.StringInList("i", o) {
			// Read out pkg db.
			// TODO: Should I use filepath stuff here?
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
				color.Set(c.DarkColor)
				fmt.Print(c.IndentChar)
				color.Unset()
				fmt.Println(l[1])

				on = l[0]
			}

			db.Close()
		} else {
			// Get all ports.
			all, err = ports.All()
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				continue
			}

			for _, p := range all {
				// Read out Pkgfile.
				f, err := os.Open(filepath.Join(c.PortDir, p, ".footprint"))
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
					color.Set(c.DarkColor)
					fmt.Print(c.IndentChar)
					color.Unset()
					fmt.Println(strings.Split(l, "\t")[2])
				}

				f.Close()
			}
		}
	}
}
