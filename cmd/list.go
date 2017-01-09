package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"sort"

	"github.com/go2c/optparse"
	"github.com/onodera-punpun/prt/pkgfile"
	"github.com/onodera-punpun/prt/ports"
)

// List lists ports.
func List(args []string) {
	// Define valid arguments.
	argi := optparse.Bool("installed", 'i', false)
	argr := optparse.Bool("repo", 'r', false)
	argv := optparse.Bool("version", 'v', false)
	argh := optparse.Bool("help", 'h', false)

	// Parse arguments.
	_, err := optparse.Parse(args)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Invaild argument, use -h for a list of arguments!")
		os.Exit(1)
	}

	if *argh {
		fmt.Println("Usage: prt list [arguments]")
		fmt.Println("")
		fmt.Println("arguments:")
		fmt.Println("  -i,   --installed       list installed ports only")
		fmt.Println("  -r,   --repo            list with repo info")
		fmt.Println("  -v,   --version         list with version info")
		fmt.Println("  -h,   --help            print help and exit")
		os.Exit(0)
	}

	// Get all ports.
	all, err := ports.All()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	var instv []string
	if *argi {
		// Get installed ports.
		inst, err := ports.Inst()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		// Get installed port versions if needed.
		if *argv {
			instv, err = ports.InstVers()
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
		}

		// Get port locations if needed.
		if *argr {
			for i, p := range inst {
				ll, err := ports.Loc(all, p)
				if err != nil {
					fmt.Fprintln(os.Stderr, err)
					continue
				}
				inst[i] = ll[0]
			}
		}

		// We want pretty output, so sort.
		sort.Strings(inst)

		// I'm using all in the the following for loop, so alias inst to all.
		all = inst
	}

	for i, p := range all {
		if *argv {
			var v string
			if *argi {
				// Get installed version.
				v = instv[i]
			} else {
				// Read out Pkgfile.
				f, err := ioutil.ReadFile(path.Join(ports.FullLoc(p), "Pkgfile"))
				if err != nil {
					fmt.Fprintln(os.Stderr, err)
					continue
				}

				// Get available version from Pkgfile.
				v, err = pkgfile.Var(f, "version")
				if err != nil {
					fmt.Fprintln(os.Stderr, err)
					continue
				}
			}

			// Merge port and version.
			p += " " + v
		}

		// Remove repo if needed.
		if !*argr && !*argi {
			p = path.Base(p)
		}

		fmt.Println(p)
	}
}
