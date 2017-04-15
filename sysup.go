package main

import (
	"fmt"
	"os"

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

		p, err := decodePort(portFullLoc(l), "Pkgfile")
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		availv := p.Pkgfile.Version + "-" + p.Pkgfile.Release

		// Add to toInst if installed and available version don't match.
		if availv != instv[i] {
			instMe = append(instMe, l)
		}
	}

	// Actually update ports in this loop.
	t := len(instMe)
	for i, l := range instMe {
		p, err := decodePortStrict(portFullLoc(l), "Footprint", "Md5sum", "Pkgfile")
		if err != nil {
			printe(err.Error())
			return
		}

		fmt.Printf("Updating package %d/%d, ", i+1, t)
		color.Set(config.LightColor)
		fmt.Printf(l)
		color.Unset()
		fmt.Println(".")

		if err := p.pkgmk(inst, *argv); err != nil {
			printe(err.Error())
			continue
		}
	}
}
