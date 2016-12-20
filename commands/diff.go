package commands

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/chiyouhen/getopt"
	"github.com/fatih/color"
	"github.com/onodera-punpun/prt/config"
	"github.com/onodera-punpun/prt/pkgfile"
	"github.com/onodera-punpun/prt/ports"
	"github.com/onodera-punpun/prt/utils"
)

// Diff lists outdated packages
func Diff(args []string) {
	// Define opts
	shortopts := "hnv"
	longopts := []string{
		"--help",
		"--no-alias",
		"--version",
	}

	// Read out opts
	opts, _, err := getopt.Getopt(args, shortopts, longopts)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Invaild argument, use -h for a list of arguments!")
		os.Exit(1)
	}

	for _, opt := range opts {
		switch opt[0] {
		case "-h", "--help":
			fmt.Println("Usage: prt diff [arguments]")
			fmt.Println("")
			fmt.Println("arguments:")
			fmt.Println("  -n,   --no-alias        disable aliasing")
			fmt.Println("  -v,   --no-version      print without version info")
			fmt.Println("  -h,   --help            print help and exit")
			os.Exit(0)
		case "-n", "--no-alias":
			o = append(o, "n")
		case "-v", "--version":
			o = append(o, "v")
		}
	}

	allPorts, err := ports.All()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	instPorts, err := ports.Inst()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	instVers, err := ports.InstVers()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	for i, port := range instPorts {
		// Get port location
		locs, err := ports.Loc(allPorts, port)
		if err != nil {
			continue
		}
		loc := locs[0]

		// Alias if needed
		if !utils.StringInList("a", o) {
			loc = ports.Alias(loc)
		}

		// Read out Pkgfile
		f, err := ioutil.ReadFile(config.Struct.PortDir + "/" + loc + "/Pkgfile")
		if err != nil {
			fmt.Fprintln(os.Stderr, "Could not read '"+config.Struct.PortDir+"/"+loc+"/Pkgfile!")
			continue
		}

		// Get available version
		ver, err := pkgfile.Var(f, "version")
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}
		rel, err := pkgfile.Var(f, "release")
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}
		availVer := ver + "-" + rel

		// Get installed version
		instVer := instVers[i]

		// Print if installed and available version don't match
		if availVer != instVer {
			fmt.Print(port)

			if utils.StringInList("v", o) {
				fmt.Print(instVer)

				color.Set(config.Struct.DarkColor)
				fmt.Print(" -> ")
				color.Unset()

				fmt.Print(availVer)
			}
			fmt.Println()
		}
	}
}
