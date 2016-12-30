package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/chiyouhen/getopt"
	"github.com/fatih/color"
	"github.com/onodera-punpun/prt/config"
	"github.com/onodera-punpun/prt/pkgfile"
	"github.com/onodera-punpun/prt/ports"
)

// Diff lists outdated packages.
func Diff(args []string) {
	// Load config.
	var conf = config.Load()

	// Define allowed opts.
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

	type optStruct struct {
		n bool
		v bool
	}

	var opt optStruct
	for _, o := range opts {
		switch o[0] {
		case "-h", "--help":
			fmt.Println("Usage: prt diff [arguments]")
			fmt.Println("")
			fmt.Println("arguments:")
			fmt.Println("  -n,   --no-alias        disable aliasing")
			fmt.Println("  -v,   --version         print with version info")
			fmt.Println("  -h,   --help            print help and exit")
			os.Exit(0)
		case "-n", "--no-alias":
			opt.n = true
		case "-v", "--version":
			opt.v = true
		}
	}

	// Get all ports.
	all, err := ports.All()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Get installed ports.
	inst, err := ports.Inst()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Get installed port versions.
	instv, err := ports.InstVers()
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
		if opt.n {
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

			if opt.v {
				fmt.Print(" " + instv[i])

				color.Set(conf.DarkColor)
				fmt.Print(" -> ")
				color.Unset()

				fmt.Print(availv)
			}

			fmt.Println()
		}
	}
}
