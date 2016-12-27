package commands

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/chiyouhen/getopt"
	"github.com/fatih/color"
	"github.com/onodera-punpun/prt/pkg"
	"github.com/onodera-punpun/prt/pkgfile"
	"github.com/onodera-punpun/prt/ports"
	"github.com/onodera-punpun/prt/utils"
)

func build(path string) {
	// Read out Pkgfile
	f, err := ioutil.ReadFile(filepath.Join(path, "Pkgfile"))
	if err != nil {
		fmt.Fprintln(os.Stderr, "Could not read '"+filepath.Join(path, "Pkgfile")+"'!")
		return
	}

	// Read out Pkgfile dependencies
	deps, err := pkgfile.Depends(f, "Depends on")
	if err != nil {
		return
	}

	for _, dep := range deps {
		// Continue if already checked
		if utils.StringInList(dep, checkPorts) {
			continue
		}
		checkPorts = append(checkPorts, dep)

		// Get port location
		locs, err := ports.Loc(allPorts, dep)
		if err != nil {
			continue
		}
		loc := locs[0]

		// Alias if needed
		if !utils.StringInList("n", o) {
			loc = ports.Alias(loc)
		}

		// Continue if already installed
		if utils.StringInList(filepath.Base(loc), instPorts) {
			continue
		}
		// Core packages should always be installed
		if filepath.Dir(loc) == "core" {
			continue
		}

		toInst = append(toInst, loc)

		// Loop
		build(filepath.Join(c.PortDir, loc))
	}
}

// Build builds ports
func Build(args []string) {
	// Define opts
	shortopts := "hv"
	longopts := []string{
		"--help",
		"--verbose",
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
			fmt.Println("Usage: prt build [arguments]")
			fmt.Println("")
			fmt.Println("arguments:")
			fmt.Println("  -v,   --verbose         toggle verbose output")
			fmt.Println("  -h,   --help            print help and exit")
			os.Exit(0)
		case "-v", "--verbose":
			o = append(o, "v")
		}
	}

	allPorts, err = ports.All()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	instPorts, err = ports.Inst()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// So uhh... I know I can do this in the opts for loop above
	// but I like consitensy and I do it like this in all other commands
	if utils.StringInList("v", o) {
		v = true
	}

	build("./")
	toInst := append(toInst)

	t := len(toInst)
	for i, port := range toInst {
		fmt.Printf("Installing port %d/%d, ", i+1, t)
		color.Set(c.LightColor)
		fmt.Printf(port)
		color.Unset()
		fmt.Println(".")

		color.Set(c.DarkColor)
		fmt.Printf(c.IndentChar)
		color.Unset()
		fmt.Println("Downloading sources")
		err = pkg.Download(v)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		color.Set(c.DarkColor)
		fmt.Printf(c.IndentChar)
		color.Unset()
		fmt.Println("Extracting sources")
		err = pkg.Extract(v)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		color.Set(c.DarkColor)
		fmt.Printf(c.IndentChar)
		color.Unset()
		fmt.Println("Building package")
		err = pkg.Build(v)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		color.Set(c.DarkColor)
		fmt.Printf(c.IndentChar)
		color.Unset()
		fmt.Println("Installing package")
		err = pkg.Install("TODO", v)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}
	}
}
