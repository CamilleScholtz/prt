package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/chiyouhen/getopt"
	"github.com/fatih/color"
	"github.com/onodera-punpun/prt/config"
	"github.com/onodera-punpun/prt/pkg"
	"github.com/onodera-punpun/prt/pkgfile"
	"github.com/onodera-punpun/prt/ports"
	"github.com/onodera-punpun/prt/utils"
)

// Sysup updates outdated packages.
func Sysup(args []string) {
	// Load config.
	var conf = config.Load()

	// Define allowed opts.
	shortopts := "hsv"
	longopts := []string{
		"--help",
		"--skip",
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
			fmt.Println("Usage: prt sysup [arguments] [ports to skip]")
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

	// Get installed port versions.
	instv, err := ports.InstVers()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Get out of date ports.
	var instMe []string
	for i, p := range inst {
		// Get port location.
		ll, err := ports.Loc(all, p)
		if err != nil {
			continue
		}
		l := ll[0]

		// Alias.
		l = ports.Alias(l)

		// Don't add ports to instMe if in vals.
		if utils.StringInList(l, vals) {
			continue
		}

		// Read out Pkgfile.
		f, err := ioutil.ReadFile(path.Join(ports.FullLoc(l), "Pkgfile"))
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		// Get available version.
		v, err := pkgfile.Var(f, "version")
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}
		r, err := pkgfile.Var(f, "release")
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
	for i, p := range instMe {
		// Set location.
		l := ports.FullLoc(p)

		fmt.Printf("Updating package %d/%d, ", i+1, t)
		color.Set(conf.LightColor)
		fmt.Printf(p)
		color.Unset()
		fmt.Println(".")

		if _, err := os.Stat(path.Join(l, "README")); err == nil {
			utils.Printi("This port has a README")
		}

		if _, err := os.Stat(path.Join(l, "pre-install")); err == nil {
			utils.Printi("Running pre-install")
			if err = pkg.PreInstall(l, opt.v); err != nil {
				utils.Printe(err.Error())
				os.Exit(1)
			}
		}

		utils.Printi("Downloading sources")
		if err := pkg.Download(l, opt.v); err != nil {
			utils.Printe(err.Error())
			continue
		}

		utils.Printi("Unpacking sources")
		if err := pkg.Unpack(l, opt.v); err != nil {
			utils.Printe(err.Error())
			continue
		}

		utils.Printi("Building package")
		if err := pkg.Build(l, false, opt.v); err != nil {
			utils.Printe(err.Error())
			continue
		}

		utils.Printi("Updating package")
		if err := pkg.Update(l, opt.v); err != nil {
			utils.Printe(err.Error())
			continue
		}

		if _, err := os.Stat(path.Join(l, "post-install")); err == nil {
			utils.Printi("Running post-install")
			if err := pkg.PostInstall(l, opt.v); err != nil {
				utils.Printe(err.Error())
				os.Exit(1)
			}
		}
	}
}
