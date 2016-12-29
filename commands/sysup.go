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
			fmt.Println("Usage: prt sysup [arguments]")
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

	// Get version of installed ports.
	instv, err = ports.InstVers()
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
			err = pkgmk.PreInstall(l, v)
			if err != nil {
				utils.Printe(err.Error())
				os.Exit(1)
			}
		}

		utils.Printi("Downloading sources")
		err = pkgmk.Download(l, v)
		if err != nil {
			utils.Printe(err.Error())
			continue
		}

		utils.Printi("Unpacking sources")
		err = pkgmk.Unpack(l, v)
		if err != nil {
			utils.Printe(err.Error())
			continue
		}

		utils.Printi("Building package")
		err = pkgmk.Build(l, v)
		if err != nil {
			utils.Printe(err.Error())
			continue
		}

		utils.Printi("Updating package")
		err = pkgmk.Update(l, v)
		if err != nil {
			utils.Printe(err.Error())
			continue
		}

		_, err = os.Stat(path.Join(l, "post-install"))
		if err == nil {
			utils.Printi("Running post-install")
			err = pkgmk.PostInstall(l, v)
			if err != nil {
				utils.Printe(err.Error())
				os.Exit(1)
			}
		}
	}
}
