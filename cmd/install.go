package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/chiyouhen/getopt"
	"github.com/fatih/color"
	"github.com/onodera-punpun/prt/pkg"
	"github.com/onodera-punpun/prt/pkgfile"
	"github.com/onodera-punpun/prt/ports"
	"github.com/onodera-punpun/prt/utils"
)

// TODO: Make this return something instead of filling a list.
func install(l string) {
	// Read out Pkgfile.
	f, err := ioutil.ReadFile(path.Join(l, "Pkgfile"))
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
		if utils.StringInList(path.Base(l), inst) {
			continue
		}
		// Core packages should always be installed.
		if path.Dir(l) == "core" {
			continue
		}

		toInst = append(toInst, l)

		// Loop.
		install(ports.FullLoc(l))
	}
}

// Install builds and installs packages.
func Install(args []string) {
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
			fmt.Println("Usage: prt install [arguments]")
			fmt.Println("")
			fmt.Println("arguments:")
			fmt.Println("  -v,   --verbose         enable verbose output")
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

	// Get ports to build.
	install("./")
	wd, err := os.Getwd()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Add current working dir to ports to install.
	if strings.Contains(wd, c.PortDir) {
		toInst = append(toInst, ports.BaseLoc(wd))
	} else {
		// Read out Pkgfile.
		f, err := ioutil.ReadFile("./Pkgfile")
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		d, err := pkgfile.Var(f, "name")
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		toInst = append(toInst, d)
	}

	t := len(toInst)
	for i, p := range toInst {
		// Set location.
		var l string
		if strings.Contains(p, "/") {
			l = ports.FullLoc(p)
		} else {
			l = wd
		}

		fmt.Printf("Installing port %d/%d, ", i+1, t)
		color.Set(c.LightColor)
		fmt.Printf(p)
		color.Unset()
		fmt.Println(".")

		_, err = os.Stat(path.Join(l, "pre-install"))
		if err == nil {
			utils.Printi("Running pre-install")
			err = pkg.PreInstall(l, utils.StringInList("v", o))
			if err != nil {
				utils.Printe(err.Error())
				os.Exit(1)
			}
		}

		utils.Printi("Downloading sources")
		err = pkg.Download(l, utils.StringInList("v", o))
		if err != nil {
			utils.Printe(err.Error())
			os.Exit(1)
		}

		utils.Printi("Unpacking sources")
		err = pkg.Unpack(l, utils.StringInList("v", o))
		if err != nil {
			utils.Printe(err.Error())
			os.Exit(1)
		}

		utils.Printi("Building package")
		err = pkg.Build(l, utils.StringInList("v", o))
		if err != nil {
			utils.Printe(err.Error())
			os.Exit(1)
		}

		utils.Printi("Installing package")
		err = pkg.Install(l, utils.StringInList("v", o))
		if err != nil {
			utils.Printe(err.Error())
			os.Exit(1)
		}

		_, err = os.Stat(path.Join(l, "post-install"))
		if err == nil {
			utils.Printi("Running post-install")
			err = pkg.PostInstall(l, utils.StringInList("v", o))
			if err != nil {
				utils.Printe(err.Error())
				os.Exit(1)
			}
		}
	}
}
