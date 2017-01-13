package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/fatih/color"
	"github.com/go2c/optparse"
	"github.com/onodera-punpun/prt/config"
	"github.com/onodera-punpun/prt/pkg"
	"github.com/onodera-punpun/prt/pkgfile"
	"github.com/onodera-punpun/prt/ports"
	"github.com/onodera-punpun/prt/utils"
)

// Install builds and installs packages.
func Install(args []string) {
	// Load config.
	conf := config.Load()

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

	if *argh {
		fmt.Println("Usage: prt install [arguments] [ports to skip]")
		fmt.Println("")
		fmt.Println("arguments:")
		fmt.Println("  -v,   --verbose         enable verbose output")
		fmt.Println("  -h,   --help            print help and exit")
		os.Exit(0)
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

	// Recursive loop that adds dependencies to instMe.
	instMeMap := make(map[int][]string)
	var c []string
	var i int
	var recursive func(l string)
	recursive = func(l string) {
		// Read out Pkgfile.
		f, err := ioutil.ReadFile(path.Join(l, "Pkgfile"))
		if err != nil {
			utils.Printe(err.Error())
			return
		}

		// Get dependencies from Pkgfile.
		dl, err := pkgfile.Depends(f, "Depends on")
		if err != nil {
			return
		}

		for _, p := range dl {
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

			// Continue if port is already installed.
			if utils.StringInList(path.Base(l), inst) {
				continue
			}

			// Increment tree level.
			if !utils.StringInList(p, c) {
				i++
			}

			// Finally prepend port to instMeMap.
			instMeMap[i] = append([]string{l}, instMeMap[i]...)

			// Continue if the port has already been checked.
			if utils.StringInList(p, c) {
				// Swap maps.
				var n int
				for i := 0; i <= len(instMeMap); i++ {
					if utils.StringInList(p, instMeMap[i]) {
						n = i
					}
				}
				t := instMeMap[n]
				instMeMap[n] = instMeMap[i]
				instMeMap[i] = t

				continue
			}

			// Append port to checked ports.
			c = append(c, p)

			// Loop.
			recursive(ports.FullLoc(l))

			// If we end up here, decrement tree level.
			i--
		}
	}
	recursive("./")

	// Convert InstMeMap to list
	c = []string{}
	var instMe []string
	for i := len(instMeMap) - 1; i >= 0; i-- {
		for _, p := range instMeMap[i] {
			// Continue if the port has already been checked.
			if utils.StringInList(p, c) {
				continue
			}

			// Finally prepend port to instMe.
			instMe = append(instMe, p)

			// Append port to checked ports.
			c = append(c, p)
		}
	}

	for _, k := range instMe {
		fmt.Println(k)
	}
	os.Exit(0)

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

		if _, err := os.Stat(path.Join(l, "pre-install")); err == nil {
			utils.Printi("Running pre-install")
			err = pkg.PreInstall(l, *argv)
			if err != nil {
				utils.Printe(err.Error())
				os.Exit(1)
			}
		}

		utils.Printi("Downloading sources")
		if err := pkg.Download(l, *argv); err != nil {
			utils.Printe(err.Error())
			os.Exit(1)
		}

		utils.Printi("Unpacking sources")
		if err := pkg.Unpack(l, *argv); err != nil {
			utils.Printe(err.Error())
			os.Exit(1)
		}

		utils.Printi("Building package")
		if utils.StringInList(path.Base(p), inst) {
			if err := pkg.Build(l, true, *argv); err != nil {
				utils.Printe(err.Error())
				os.Exit(1)
			}
		} else {
			if err := pkg.Build(l, false, *argv); err != nil {
				utils.Printe(err.Error())
				os.Exit(1)
			}
		}

		if utils.StringInList(path.Base(p), inst) {
			utils.Printi("Updating package")
			if err := pkg.Update(l, *argv); err != nil {
				utils.Printe(err.Error())
				os.Exit(1)
			}
		} else {
			utils.Printi("Installing package")
			if err := pkg.Install(l, *argv); err != nil {
				utils.Printe(err.Error())
				os.Exit(1)
			}
		}

		if _, err = os.Stat(path.Join(l, "post-install")); err == nil {
			utils.Printi("Running post-install")
			err = pkg.PostInstall(l, *argv)
			if err != nil {
				utils.Printe(err.Error())
				os.Exit(1)
			}
		}
	}
}
