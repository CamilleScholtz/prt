package main

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/fatih/color"
	"github.com/go2c/optparse"
)

// install builds and installs packages.
func install(args []string) {
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
		fmt.Println("Usage: prt install [arguments] [ports to skip]")
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

	// Recursive loop that adds dependencies to instMe.
	instMeMap := make(map[int][]string)
	var c []string
	var i int
	var recursive func(l string)
	recursive = func(l string) {
		// Get dependencies from Pkgfile.
		if err := initPkgfile(l, []string{"Depends"}); err != nil {
			return
		}

		// Get location and dependencies for each port in dependency list.
		for _, p := range pkgfile.Depends {
			// Get port location.
			ll, err := portLoc(all, p)
			if err != nil {
				continue
			}
			l := ll[0]

			// Alias ports.
			l = portAlias(l)

			// Don't add ports to instMe if in vals.
			if stringInList(l, vals) {
				continue
			}

			// Continue if port is already installed.
			if stringInList(path.Base(l), inst) {
				continue
			}

			// Increment tree level.
			if !stringInList(p, c) {
				i++
			}

			// Prepend port to instMeMap.
			instMeMap[i] = append([]string{l}, instMeMap[i]...)

			// Continue if the port has already been checked.
			if stringInList(p, c) {
				// We will also "merge maps" here, here is a quick ASCII illustration
				// using the `prt depends -t` syntax of what this basically does:
				//
				// BEFORE:
				// port1
				// - port2
				// - - port3
				// port4
				// - port2 *
				//
				// AFTER:
				// port1
				// - port2
				// port4
				// - port2
				// - - port3
				//
				// We do this because without this "merge" port2 would be
				// complaining about how port3 isn't installed, since we iterrate
				// over the "list" from bottom to top.
				var n int
				for i := 0; i <= len(instMeMap); i++ {
					if stringInList(p, instMeMap[i]) {
						n = i
					}
				}
				instMeMap[n] = append(instMeMap[n], instMeMap[i]...)

				continue
			}

			// Append port to checked ports.
			c = append(c, p)

			// Loop.
			recursive(portFullLoc(l))

			// If we end up here, decrement tree level.
			i--
		}
	}
	recursive("./")

	// Convert InstMeMap to list (instMe).
	c = []string{}
	var instMe []string
	for i := len(instMeMap); i >= 0; i-- {
		for _, p := range instMeMap[i] {
			// Continue if the port has already been checked.
			if stringInList(p, c) {
				continue
			}

			// Append port to instMe.
			instMe = append(instMe, p)

			// Append port to checked ports.
			c = append(c, p)
		}
	}

	// Add current working dir to ports to instMe.
	wd, err := os.Getwd()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Strip of PortDir if needed.
	if strings.Contains(wd, config.PortDir) {
		instMe = append(instMe, portBaseLoc(wd))
	} else {
		// Get port name from Pkgfile.
		if err := initPkgfile(".", []string{"Name"}); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		// Add name to ports to install.
		instMe = append(instMe, pkgfile.Name)
	}

	// Actually install ports in this loop.
	t := len(instMe)
	for i, p := range instMe {
		// Set location.
		var l string
		if strings.Contains(p, "/") {
			l = portFullLoc(p)
		} else {
			l = wd
		}

		if stringInList(path.Base(p), inst) {
			fmt.Printf("Updating package %d/%d, ", i+1, t)
		} else {
			fmt.Printf("Installing package %d/%d, ", i+1, t)
		}
		color.Set(config.LightColor)
		fmt.Printf(p)
		color.Unset()
		fmt.Println(".")

		if _, err := os.Stat(path.Join(l, "pre-install")); err == nil {
			printi("Running pre-install")
			err = pkgPreInstall(l, *argv)
			if err != nil {
				printe(err.Error())
				os.Exit(1)
			}
		}

		printi("Downloading sources")
		if err := pkgDownload(l, *argv); err != nil {
			printe(err.Error())
			os.Exit(1)
		}

		printi("Unpacking sources")
		if err := pkgUnpack(l, *argv); err != nil {
			printe(err.Error())
			os.Exit(1)
		}

		printi("Building package")
		if stringInList(path.Base(p), inst) {
			if err := pkgBuild(l, true, *argv); err != nil {
				printe(err.Error())
				os.Exit(1)
			}
		} else {
			if err := pkgBuild(l, false, *argv); err != nil {
				printe(err.Error())
				os.Exit(1)
			}
		}

		if stringInList(path.Base(p), inst) {
			printi("Updating package")
			if err := pkgUpdate(l, *argv); err != nil {
				printe(err.Error())
				os.Exit(1)
			}
		} else {
			printi("Installing package")
			if err := pkgInstall(l, *argv); err != nil {
				printe(err.Error())
				os.Exit(1)
			}
		}

		if _, err = os.Stat(path.Join(l, "post-install")); err == nil {
			printi("Running post-install")
			err = pkgPostInstall(l, *argv)
			if err != nil {
				printe(err.Error())
				os.Exit(1)
			}
		}
	}
}
