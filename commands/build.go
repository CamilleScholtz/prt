package commands

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/chiyouhen/getopt"
	"github.com/fatih/color"
	"github.com/onodera-punpun/prt/pkg"
	"github.com/onodera-punpun/prt/pkgfile"
	"github.com/onodera-punpun/prt/ports"
	"github.com/onodera-punpun/prt/utils"
)

func build(l string) {
	// Read out Pkgfile.
	f, err := ioutil.ReadFile(filepath.Join(l, "Pkgfile"))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	// Read out Pkgfile dependencies.
	dl, err := pkgfile.Depends(f, "Depends on")
	if err != nil {
		return
	}

	for _, p := range dl {
		// Continue if already dependency has already been checked.
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
		if utils.StringInList(filepath.Base(l), inst) {
			continue
		}
		// Core packages should always be installed.
		if filepath.Dir(l) == "core" {
			continue
		}

		toInst = append(toInst, l)

		// Loop.
		build(filepath.Join(c.PortDir, l))
	}
}

// Build builds ports.
func Build(args []string) {
	// Define opts.
	shortopts := "hv"
	longopts := []string{
		"--help",
		"--verbose",
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
			fmt.Println("Usage: prt build [arguments]")
			fmt.Println("")
			fmt.Println("arguments:")
			fmt.Println("  -v,   --verbose         toggle verbose output")
			fmt.Println("  -h,   --help            print help and exit")
			os.Exit(0)
		case "-v", "--verbose":
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

	// So uhh... I know I can do this in the opts for loop above
	// but I like consitensy and I do it like this in all other commands.
	var v bool
	if utils.StringInList("v", o) {
		v = true
	}

	// Get ports to build.
	build("./")
	wd, err := os.Getwd()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	// TODO: Remove PortDir here.
	toInst := append(toInst, wd)

	t := len(toInst)
	for i, p := range toInst {
		os.Chdir(filepath.Join(c.PortDir, p))

		fmt.Printf("Installing port %d/%d, ", i+1, t)
		color.Set(c.LightColor)
		fmt.Printf(p)
		color.Unset()
		fmt.Println(".")

		utils.Printi("Downloading sources")
		err = pkg.Download(v)
		if err != nil {
			utils.Printe(err.Error())
			continue
		}

		utils.Printi("Unpacking sources")
		err = pkg.Unpack(v)
		if err != nil {
			utils.Printe(err.Error())
			continue
		}

		utils.Printi("Building package")
		err = pkg.Build(v)
		if err != nil {
			utils.Printe(err.Error())
			continue
		}

		utils.Printi("Installing package sources")
		err = pkg.Install("TODO", v)
		if err != nil {
			utils.Printe(err.Error())
			continue
		}
	}
}
