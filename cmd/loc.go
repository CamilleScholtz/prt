package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/chiyouhen/getopt"
	"github.com/fatih/color"
	"github.com/onodera-punpun/prt/ports"
	"github.com/onodera-punpun/prt/utils"
)

// Loc prints port locations
func Loc(args []string) {
	// Define opts.
	shortopts := "hdn"
	longopts := []string{
		"--help",
		"--no-alias",
		"--duplicate",
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
			fmt.Println("Usage: prt loc [arguments] [ports]")
			fmt.Println("")
			fmt.Println("arguments:")
			fmt.Println("  -d,   --duplicate       list duplicate ports as well")
			fmt.Println("  -n,   --no-alias        disable aliasing")
			fmt.Println("  -h,   --help            print help and exit")
			os.Exit(0)
		case "-d", "--duplicate":
			o = append(o, "d")
		case "-n", "--no-alias":
			o = append(o, "n")
		}
	}

	// This command needs a value.
	if len(vals) == 0 {
		fmt.Fprintln(os.Stderr, "Please specify a port!")
		os.Exit(1)
	}

	// Get all ports.
	all, err = ports.All()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	for _, p := range vals {
		// Continue if already checked.
		if utils.StringInList(p, cp) {
			continue
		}
		cp = append(cp, p)

		ll, err := ports.Loc(all, p)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}
		if !utils.StringInList("d", o) {
			ll = []string{ll[0]}
		}

		for _, l := range ll {
			// Alias if needed.
			if !utils.StringInList("a", o) {
				l = ports.Alias(l)
			}

			// Print duplicate indentation.
			if utils.StringInList("d", o) {
				if i > 0 {
					color.Set(c.DarkColor)
					fmt.Printf(strings.Repeat(c.IndentChar, i))
					color.Unset()
				}
				i++
			}

			// Finally print the port.
			fmt.Println(l)
		}
	}
}
