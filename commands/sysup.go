package commands

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/chiyouhen/getopt"
	"github.com/fatih/color"
	"github.com/onodera-punpun/prt/pkgfile"
	"github.com/onodera-punpun/prt/pkgmk"
	"github.com/onodera-punpun/prt/ports"
	"github.com/onodera-punpun/prt/utils"
)

// Sysup updates outdated packages.
func Sysup(args []string) {
	// Define opts.
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

	var skip []string
	for i, opt := range opts {
		switch opt[0] {
		case "-h", "--help":
			fmt.Println("Usage: prt sysup [arguments]")
			fmt.Println("")
			fmt.Println("arguments:")
			fmt.Println("  -s,   --skip            skip port from updating")
			fmt.Println("  -v,   --verbose         enable verbose output")
			fmt.Println("  -h,   --help            print help and exit")
			os.Exit(0)
		case "-s", "--skip":
			o = append(o, "s")
			// TODO: This isn't 100% perfect.
			skip = append(skip, vals[i])
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

	// Get version of installed ports.
	instv, err = ports.InstVers()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Get out of date ports.
	for i, p := range inst {
		// Get port location.
		ll, err := ports.Loc(all, p)
		if err != nil {
			continue
		}
		l := ll[0]

		// Alias.
		l = ports.Alias(l)

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
			toInst = append(toInst, l)
		}
	}

	// Remove ports from toInst if needed.
	if utils.StringInList("s", o) {
		for _, val := range vals {
			fmt.Println(val)
		}
	}

	t := len(toInst)
	for i, p := range toInst {
		// Set location.
		l := ports.FullLoc(p)

		fmt.Printf("Updating port %d/%d, ", i+1, t)
		color.Set(c.LightColor)
		fmt.Printf(p)
		color.Unset()
		fmt.Println(".")

		_, err = os.Stat(path.Join(l, "pre-install"))
		if err == nil {
			utils.Printi("Running pre-install")
			err = pkgmk.PreInstall(l, utils.StringInList("v", o))
			if err != nil {
				utils.Printe(err.Error())
				os.Exit(1)
			}
		}

		utils.Printi("Downloading sources")
		err = pkgmk.Download(l, utils.StringInList("v", o))
		if err != nil {
			utils.Printe(err.Error())
			continue
		}

		utils.Printi("Unpacking sources")
		err = pkgmk.Unpack(l, utils.StringInList("v", o))
		if err != nil {
			utils.Printe(err.Error())
			continue
		}

		utils.Printi("Building package")
		err = pkgmk.Build(l, utils.StringInList("v", o))
		if err != nil {
			utils.Printe(err.Error())
			continue
		}

		utils.Printi("Updating package")
		err = pkgmk.Update(l, utils.StringInList("v", o))
		if err != nil {
			utils.Printe(err.Error())
			continue
		}

		_, err = os.Stat(path.Join(l, "post-install"))
		if err == nil {
			utils.Printi("Running post-install")
			err = pkgmk.PostInstall(l, utils.StringInList("v", o))
			if err != nil {
				utils.Printe(err.Error())
				os.Exit(1)
			}
		}
	}
}
