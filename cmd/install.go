package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/chiyouhen/getopt"
	"github.com/fatih/color"
	"github.com/onodera-punpun/prt/config"
	"github.com/onodera-punpun/prt/pkg"
	"github.com/onodera-punpun/prt/pkgfile"
	"github.com/onodera-punpun/prt/ports"
	"github.com/onodera-punpun/prt/utils"
)

// Install builds and installs packages.
func Install(args []string) {
	// Load config.
	var conf = config.Load()

	// Define allowed opts.
	shortopts := "hv"
	longopts := []string{
		"--help",
		"--verbose",
	}

	// Read out opts.
	opts, vals, err := getopt.Getopt(args, shortopts, longopts)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Invaild argument, use -h for a list of arguments!")
		os.Exit(1)
	}

	type optStruct struct {
		v bool
	}

	var opt optStruct
	for _, o := range opts {
		switch o[0] {
		case "-h", "--help":
			fmt.Println("Usage: prt install [arguments] [ports to skip]")
			fmt.Println("")
			fmt.Println("arguments:")
			fmt.Println("  -v,   --verbose         enable verbose output")
			fmt.Println("  -h,   --help            print help and exit")
			os.Exit(0)
		case "-v", "--verbose":
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

	// Recursive loop that add dependencies to instMe.
	var instMe []string
	var c []string
	var recursive func(l string)
	recursive = func(l string) {
		// Read out Pkgfile.
		f, err := ioutil.ReadFile(path.Join(l, "Pkgfile"))
		if err != nil {
			utils.Printe(err.Error())
			return
		}

		// Get dependencies.
		dl, err := pkgfile.Depends(f, "Depends on")
		if err != nil {
			return
		}

		for _, p := range dl {
			// Continue if already checked.
			if utils.StringInList(p, c) {
				continue
			}
			// Add to checked ports.
			c = append(c, p)

			// Get port location.
			ll, err := ports.Loc(all, p)
			if err != nil {
				continue
			}
			l := ll[0]

			// Alias ports.
			l = ports.Alias(l)

			// Don't add ports to instMe if in vals.
			if utils.StringInList(l, vals) {
				continue
			}

			// Continue port is already installed.
			if utils.StringInList(path.Base(l), inst) {
				continue
			}

			instMe = append(instMe, l)

			// Loop.
			recursive(ports.FullLoc(l))
		}
	}
	recursive("./")

	// Reverse list
	instMe = utils.ReverseList(instMe)

	// Add current working dir to ports to install.
	wd, err := os.Getwd()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Strip of PortDir if needed.
	if strings.Contains(wd, conf.PortDir) {
		instMe = append(instMe, ports.BaseLoc(wd))
	} else {
		// Read out Pkgfile.
		f, err := ioutil.ReadFile("./Pkgfile")
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		// Get port name.
		d, err := pkgfile.Var(f, "name")
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		// Add name to ports to install.
		instMe = append(instMe, d)
	}

	t := len(instMe)
	for i, p := range instMe {
		// Set location.
		var l string
		if strings.Contains(p, "/") {
			l = ports.FullLoc(p)
		} else {
			l = wd
		}

		if utils.StringInList(path.Base(p), inst) {
			fmt.Printf("Updating package %d/%d, ", i+1, t)
		} else {
			fmt.Printf("Installing package %d/%d, ", i+1, t)
		}
		color.Set(conf.LightColor)
		fmt.Printf(p)
		color.Unset()
		fmt.Println(".")

		if _, err := os.Stat(path.Join(l, "README")); err == nil {
			utils.Printi("This port has a README")
		}

		if _, err := os.Stat(path.Join(l, "pre-install")); err == nil {
			utils.Printi("Running pre-install")
			err = pkg.PreInstall(l, opt.v)
			if err != nil {
				utils.Printe(err.Error())
				os.Exit(1)
			}
		}

		utils.Printi("Downloading sources")
		if err := pkg.Download(l, opt.v); err != nil {
			utils.Printe(err.Error())
			os.Exit(1)
		}

		utils.Printi("Unpacking sources")
		if err := pkg.Unpack(l, opt.v); err != nil {
			utils.Printe(err.Error())
			os.Exit(1)
		}

		utils.Printi("Building package")
		if utils.StringInList(path.Base(p), inst) {
			if err := pkg.Build(l, true, opt.v); err != nil {
				utils.Printe(err.Error())
				os.Exit(1)
			}
		} else {
			if err := pkg.Build(l, false, opt.v); err != nil {
				utils.Printe(err.Error())
				os.Exit(1)
			}
		}

		if utils.StringInList(path.Base(p), inst) {
			utils.Printi("Updating package")
			if err := pkg.Update(l, opt.v); err != nil {
				utils.Printe(err.Error())
				os.Exit(1)
			}
		} else {
			utils.Printi("Installing package")
			if err := pkg.Install(l, opt.v); err != nil {
				utils.Printe(err.Error())
				os.Exit(1)
			}
		}

		if _, err = os.Stat(path.Join(l, "post-install")); err == nil {
			utils.Printi("Running post-install")
			err = pkg.PostInstall(l, opt.v)
			if err != nil {
				utils.Printe(err.Error())
				os.Exit(1)
			}
		}
	}
}
