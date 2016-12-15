package commands

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"unicode/utf8"

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
		"--no-version",
	}

	// Read out opts
	opts, _, err := getopt.Getopt(args, shortopts, longopts)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Invaild argument, use -h for a list of arguments!")
		os.Exit(1)
	}

	var alias, version bool
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
			alias = true
		case "-v", "--no-version":
			version = true
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

		// TODO
		if version {
			fmt.Println(ver)
		}

		if instVer != availVer {
			port = utils.TrimString(port, 24)
			fmt.Print(port)
			fmt.Printf(strings.Repeat(" ", 25-utf8.RuneCountInString(port)))

			instVer = utils.TrimString(instVer, 12)
			fmt.Print(instVer)
			fmt.Printf(strings.Repeat(" ", 13-utf8.RuneCountInString(instVer)))

			color.Set(color.FgBlack, color.Bold)
			fmt.Print("->")
			color.Unset()
			fmt.Print(" ")

			availVer = utils.TrimString(availVer, 12)
			fmt.Println(availVer)
		}
	}
}
