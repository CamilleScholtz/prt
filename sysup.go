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
	all, err := portAll()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Get installed ports.
	inst, err := portInst()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Get installed port versions.
	instv, err := portInstVers()
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

		// Get available version.
		if err := initPkgfile(portFullLoc(l), []string{"Version", "Release"}); err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}
		availv := pkgfile.Version + "-" + pkgfile.Release

		// Add to toInst if installed and available version don't match.
		if availv != instv[i] {
			instMe = append(instMe, l)
		}
	}

	t := len(instMe)
	for i, p := range instMe {
		// Set location.
		l := portFullLoc(p)

		fmt.Printf("Updating package %d/%d, ", i+1, t)
		color.Set(config.LightColor)
		fmt.Printf(p)
		color.Unset()
		fmt.Println(".")

		if _, err := os.Stat(path.Join(l, "pre-install")); err == nil {
			printi("Running pre-install")
			if err = pkgPreInstall(l, *argv); err != nil {
				printe(err.Error())
				os.Exit(1)
			}
		}

		printi("Downloading sources")
		if err := pkgDownload(l, *argv); err != nil {
			printe(err.Error())
			continue
		}

		printi("Unpacking sources")
		if err := pkgUnpack(l, *argv); err != nil {
			printe(err.Error())
			continue
		}

		printi("Building package")
		if err := pkgBuild(l, false, *argv); err != nil {
			printe(err.Error())
			continue
		}

		printi("Updating package")
		if err := pkgUpdate(l, *argv); err != nil {
			printe(err.Error())
			continue
		}

		if _, err := os.Stat(path.Join(l, "post-install")); err == nil {
			printi("Running post-install")
			if err := pkgPostInstall(l, *argv); err != nil {
				printe(err.Error())
				os.Exit(1)
			}
		}
	}
}
