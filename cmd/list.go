package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"sort"

	"github.com/chiyouhen/getopt"
	"github.com/onodera-punpun/prt/pkgfile"
	"github.com/onodera-punpun/prt/ports"
	"github.com/onodera-punpun/prt/utils"
)

// List lists ports.
func List(args []string) {
	// Define opts.
	shortopts := "hirv"
	longopts := []string{
		"--help",
		"--installed",
		"--repo",
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
			fmt.Println("Usage: prt list [arguments]")
			fmt.Println("")
			fmt.Println("arguments:")
			fmt.Println("  -i,   --installed       list installed ports only")
			fmt.Println("  -r,   --repo            list with repo info")
			fmt.Println("  -v,   --version         list with version info")
			fmt.Println("  -h,   --help            print help and exit")
			os.Exit(0)
		case "-i", "--installed":
			o = append(o, "i")
		case "-r", "--repo":
			o = append(o, "r")
		case "-v", "--version":
			o = append(o, "v")
		}
	}

	// Get all ports
	all, err = ports.All()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Only list installed ports.
	if utils.StringInList("i", o) {
		inst, err = ports.Inst()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		instv, err = ports.InstVers()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		// Get port locations
		if utils.StringInList("r", o) {
			for i, p := range inst {
				ll, err := ports.Loc(all, p)
				if err != nil {
					fmt.Fprintln(os.Stderr, err)
					continue
				}
				inst[i] = ll[0]
			}
		}

		sort.Strings(inst)
		all = inst
	}

	for i, p := range all {
		if utils.StringInList("v", o) {
			var v string
			if utils.StringInList("i", o) {
				v = instv[i]
			} else {
				// Read out Pkgfile.
				f, err := ioutil.ReadFile(path.Join(ports.FullLoc(p), "Pkgfile"))
				if err != nil {
					fmt.Fprintln(os.Stderr, err)
					continue
				}

				v, err = pkgfile.Var(f, "version")
				if err != nil {
					fmt.Fprintln(os.Stderr, err)
					continue
				}
			}

			p += " " + v
		}

		// Remove repo if needed.
		if !utils.StringInList("r", o) {
			p = path.Base(p)
		}

		fmt.Println(p)
	}
}
