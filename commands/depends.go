package commands

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/chiyouhen/getopt"
	"github.com/fatih/color"
	"github.com/onodera-punpun/prt/pkgfile"
	"github.com/onodera-punpun/prt/ports"
	"github.com/onodera-punpun/prt/utils"
)

func depends(l string) {
	// Read out Pkgfile.
	f, err := ioutil.ReadFile(filepath.Join(l, "Pkgfile"))
	if err != nil {
		utils.Printe(err.Error())
		return
	}

	// Read out Pkgfile dependencies.
	dl, err := pkgfile.Depends(f, "Depends on")
	if err != nil {
		return
	}

	for _, p := range dl {
		// Continue if already checked.
		if utils.StringInList(p, cp) {
			continue
		}
		cp = append(cp, p)

		// Get port location.
		ll, err := ports.Loc(all, p)
		if err != nil {
			continue
		}
		l := ll[0]

		// Alias if needed.
		if !utils.StringInList("n", o) {
			l = ports.Alias(l)
		}

		// Continue port is already installed.
		if !utils.StringInList("a", o) {
			if utils.StringInList(filepath.Base(l), inst) {
				continue
			}
			// Core packages should always be installed.
			if filepath.Dir(l) == "core" {
				continue
			}
		}

		// Print tree indentation.
		if utils.StringInList("t", o) {
			if i > 0 {
				color.Set(c.DarkColor)
				fmt.Printf(strings.Repeat(c.IndentChar, i))
				color.Unset()
			}
			i++
		}

		// Finally print the port.
		fmt.Println(l)

		// Loop.
		depends(filepath.Join(c.PortDir, l))

		if utils.StringInList("t", o) {
			i--
		}
	}
}

// Depends lists dependencies recursivly.
func Depends(args []string) {
	// Define opts.
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

	for _, opt := range opts {
		switch opt[0] {
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
			o = append(o, "a")
		case "-n", "--no-alias":
			o = append(o, "n")
		case "-t", "--tree":
			o = append(o, "t")
		}
	}

	// Get all and all installed ports.
	all, err = ports.All()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if !utils.StringInList("a", o) {
		inst, err = ports.Inst()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}

	depends("./")
}
