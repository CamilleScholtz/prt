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
)

func Diff(args []string) {
	// Define opts
	shortopts := "hn"
	longopts := []string{
		"--help",
		"--no-alias",
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
			fmt.Println("  -h,   --help            print help and exit")
			os.Exit(0)
		case "-n", "--no-alias":
			alias = true
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

	for _, port := range instPorts {
		// Get port location
		locs, err := ports.Loc(allPorts, port)
		if err != nil {
			continue
		}
		loc := locs[0]

		// Alias if needed
		if !alias {
			loc = ports.Alias(loc)
		}

		// Read out Pkgfile
		f, err := ioutil.ReadFile(config.Struct.PortDir + "/" + loc + "/Pkgfile")
		if err != nil {
			fmt.Fprintln(os.Stderr, "Could not read '"+config.Struct.PortDir+"/"+loc+"/Pkgfile!")
			continue
		}

		instVer, err := ports.InstVer(port)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}
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

		if instVer != availVer {
			fmt.Print(port + "\t\t" + instVer)
			color.Set(color.FgBlack, color.Bold)
			fmt.Print("\t-\t")
			color.Unset()
			fmt.Println(availVer)
		}
	}
}
