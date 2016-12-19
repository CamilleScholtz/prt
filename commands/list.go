package commands

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"

	"github.com/chiyouhen/getopt"
	"github.com/onodera-punpun/prt/config"
	"github.com/onodera-punpun/prt/pkgfile"
	"github.com/onodera-punpun/prt/ports"
	"github.com/onodera-punpun/prt/utils"
)

// List lists ports
func List(args []string) {
	// Define opts
	shortopts := "hirv"
	longopts := []string{
		"--help",
		"--installed",
		"--repo",
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
			fmt.Println("Usage: prt list [arguments]")
			fmt.Println("")
			fmt.Println("arguments:")
			fmt.Println("  -i,   --installed       list installed ports only")
			fmt.Println("  -r,   --repo            print with repo info")
			fmt.Println("  -v,   --version         print with version info")
			fmt.Println("  -h,   --help            print help and exit")
			os.Exit(0)
		case "-i", "--installed":
			o = append(o, "i")
		case "-r", "--repo":
			o = append(o, "r")
		case "-v", "--version":
			o = append(o, "v")
		}
	}

	allPorts, err = ports.All()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Only list installed ports
	if utils.StringInList("i", o) {
		instPorts, err = ports.Inst()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		instVers, err = ports.InstVer()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		// Get port locations
		if utils.StringInList("r", o) {
			for i, port := range instPorts {
				locs, err := ports.Loc(allPorts, port)
				if err != nil {
					fmt.Fprintln(os.Stderr, err)
					continue
				}
				instPorts[i] = locs[0]
			}
		}

		sort.Strings(instPorts)
		allPorts = instPorts
	}

	for i, port := range allPorts {
		if utils.StringInList("v", o) {
			var ver string
			if utils.StringInList("i", o) {
				ver = instVers[i]
			} else {
				// Read out Pkgfile
				f, err := ioutil.ReadFile(config.Struct.PortDir + "/" + port + "/Pkgfile")
				if err != nil {
					fmt.Fprintln(os.Stderr, "Could not read '"+config.Struct.PortDir+"/"+port+"/Pkgfile'!")
					continue
				}

				ver, err = pkgfile.Var(f, "version")
				if err != nil {
					fmt.Fprintln(os.Stderr, err)
					continue
				}

				// Remove repo if needed
				if utils.StringInList("r", o) {
					port = filepath.Base(port)
				}
			}

			port = port + " " + ver
		}

		fmt.Println(port)
	}
}
