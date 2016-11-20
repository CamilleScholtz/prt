package commands

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/chiyouhen/getopt"
	"github.com/onodera-punpun/prt/utils"
)

var Checked []string

// This function prints dependencies recursivly
func recursive(path string) {
	// Read out Pkgfile
	pkgfile, err := ioutil.ReadFile(path + "/Pkgfile")
	if err != nil {
		fmt.Println("Could not read Pkgfile.")
		os.Exit(1)
	}

	// Read out pkgfile dependencies
	var loc string
	var fix []string
	deps := utils.ReadDepends(pkgfile, "Depends on")
	if len(deps) < 1 {
		return
	}

	for _, port := range deps {
		// Continue if already checked
		if utils.StringInSlice(port, Checked) {
			continue
		} else {
			Checked = append(Checked, port)
		}

		fix = utils.PortLoc(port)
		if len(fix) < 1 {
			return
		} else {
			loc = fix[0]
		}
		fmt.Println(loc)

		// Loop
		recursive(loc)
	}
}

func Depends(args []string) {
	// Define opts
	shortopts := "hadt"
	longopts := []string{
		"--help",
		"--disable-alias",
		"--tree",
	}

	// Read out opts
	opts, _, err := getopt.Getopt(args, shortopts, longopts)
	if err != nil {
		fmt.Println("Invaild argument, use -h for a list of arguments.")
		os.Exit(1)
	}

	recursive("./")

	if len(opts) > 0 {
		for _, opt := range opts {
			switch opt[0] {
			case "-h", "--help":
				fmt.Println("Usage: prt depends [arguments]")
				fmt.Println("")
				fmt.Println("arguments:")
				fmt.Println("  -a,   --all             also list installed dependencies")
				fmt.Println("  -d,   --disable-alias   disable aliasing")
				fmt.Println("  -t,   --tree            list using tree view")
				fmt.Println("  -h,   --help            print help and exit")
				os.Exit(0)
			case "-a", "--all":
			case "-d", "--disable-alias":
			case "-t", "--tree":
			}
		}
	}
}
