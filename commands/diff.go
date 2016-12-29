package commands

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/chiyouhen/getopt"
	"github.com/fatih/color"
	"github.com/onodera-punpun/prt/pkgfile"
	"github.com/onodera-punpun/prt/ports"
	"github.com/onodera-punpun/prt/utils"
)

// Diff lists outdated packages.
func Diff(args []string) {
	// Define opts.
	shortopts := "hnv"
	longopts := []string{
		"--help",
		"--no-alias",
		"--version",
	}

	// Read out opts.
	opts, _, err := getopt.Getopt(args, shortopts, longopts)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Invaild argument, use -h for a list of arguments!")
		os.Exit(1)
	}

	for _, opt := range opts {
		switch opt[0] {
		case "-h", "--help":
			fmt.Println("Usage: prt diff [arguments]")
			fmt.Println("")
			fmt.Println("arguments:")
			fmt.Println("  -n,   --no-alias        disable aliasing")
			fmt.Println("  -v,   --version         print with version info")
			fmt.Println("  -h,   --help            print help and exit")
			os.Exit(0)
		case "-n", "--no-alias":
			o = append(o, "n")
		case "-v", "--version":
			o = append(o, "v")
		}
	}

	// Get all and all installed ports.
	all, err = ports.All()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	inst, err = ports.Inst()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Get version of installed ports.
	instv, err = ports.InstVers()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	for i, p := range inst {
		// Get port location.
		ll, err := ports.Loc(all, p)
		if err != nil {
			continue
		}
		l := ll[0]

		// Alias if needed.
		if !utils.StringInList("a", o) {
			l = ports.Alias(l)
		}

		// Read out Pkgfile.
		f, err := ioutil.ReadFile(path.Join(ports.FullLoc(l), "Pkgfile"))
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		// Get available version.
		v, err := pkgfile.Var(f, "version")
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}
		r, err := pkgfile.Var(f, "release")
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}
		availv := v + "-" + r

		// Print if installed and available version don't match.
		if availv != instv[i] {
			fmt.Print(p)

			if utils.StringInList("v", o) {
				fmt.Print(" " + instv[i])

				color.Set(c.DarkColor)
				fmt.Print(" -> ")
				color.Unset()

				fmt.Print(availv)
			}

			fmt.Println()
		}
	}
}
