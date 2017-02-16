package main

import (
	"fmt"
	"os"
	"path"

	"github.com/fatih/color"
	"github.com/go2c/optparse"
)

// sysup updates outdated packages.
func sysup(args []string) {
	// Define valid arguments.
	o := optparse.New()
	argv := o.Bool("verbose", 'v', false)
	argh := o.Bool("help", 'h', false)

	// Parse arguments.
	vals, err := o.Parse(args)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Invaild argument, use -h for a list of arguments!")
		os.Exit(1)
	}

	// Print help.
	if *argh {
		fmt.Println("Usage: prt sysup [arguments] [ports to skip]")
		fmt.Println("")
		fmt.Println("arguments:")
		fmt.Println("  -v,   --verbose         enable verbose output")
		fmt.Println("  -h,   --help            print help and exit")
		os.Exit(0)
	}

	// Get all ports.
	all, err := allPorts()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Get installed ports.
	inst, err := instPorts()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Get installed port versions.
	instv, err := instVersPorts()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Get out of date ports.
	var instMe []string
	for i, p := range inst {
		// Get port location.
		ll, err := portLoc(all, p)
		if err != nil {
			continue
		}
		l := ll[0]

		// Alias.
		l = portAlias(l)

		// Don't add ports to instMe if in vals.
		if stringInList(l, vals) {
			continue
		}

		// Read out Pkgfile.
		f, err := readPkgfile(path.Join(portFullLoc(l), "Pkgfile"))
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		// Get available version.
		v, err := f.variable("version")
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}
		r, err := f.variable("release")
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}
		availv := v + "-" + r

		// Add to toInst if installed and available version don't match.
		if availv != instv[i] {
			instMe = append(instMe, l)
		}
	}

	t := len(instMe)
	for i, l := range instMe {
		// Read out Pkgfile.
		f, err := readPkgfile(path.Join(portFullLoc(l), "Pkgfile"))
		if err != nil {
			printe(err.Error())
			return
		}

		fmt.Printf("Updating package %d/%d, ", i+1, t)
		color.Set(config.LightColor)
		fmt.Printf(l)
		color.Unset()
		fmt.Println(".")

		if _, err := os.Stat(path.Join(l, "pre-install")); err == nil {
			printi("Running pre-install")
			if err = f.pre(*argv); err != nil {
				printe(err.Error())
				os.Exit(1)
			}
		}

		if err := f.download(*argv); err != nil {
			printe(err.Error())
			continue
		}

		if err := f.unpack(*argv); err != nil {
			printe(err.Error())
			continue
		}

		printi("Building package")
		if err := f.build(false, *argv); err != nil {
			printe(err.Error())
			continue
		}

		printi("Updating package")
		if err := f.update(*argv); err != nil {
			printe(err.Error())
			continue
		}

		if _, err := os.Stat(path.Join(l, "post-install")); err == nil {
			printi("Running post-install")
			if err := f.post(*argv); err != nil {
				printe(err.Error())
				os.Exit(1)
			}
		}
	}
}
