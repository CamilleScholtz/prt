package cmd

import (
	"fmt"
	"os"

	"github.com/chiyouhen/getopt"
	"github.com/fatih/color"
	"github.com/onodera-punpun/prt/config"
	"github.com/onodera-punpun/prt/pkg"
	"github.com/onodera-punpun/prt/utils"
)

// Uninstall uninstalls packages.
func Uninstall(args []string) {
	// Load config.
	var conf = config.Load()

	// Define allowed opts.
	shortopts := "h"
	longopts := []string{
		"--help",
	}

	// Read out opts.
	opts, vals, err := getopt.Getopt(args, shortopts, longopts)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Invaild argument, use -h for a list of arguments!")
		os.Exit(1)
	}

	for _, o := range opts {
		switch o[0] {
		case "-h", "--help":
			fmt.Println("Usage: prt uninstall [arguments] [packages]")
			fmt.Println("")
			fmt.Println("arguments:")
			fmt.Println("  -h,   --help            print help and exit")
			os.Exit(0)
		}
	}

	// This command needs a value.
	if len(vals) == 0 {
		fmt.Fprintln(os.Stderr, "Please specify a package!")
		os.Exit(1)
	}

	t := len(vals)
	for i, p := range vals {
		fmt.Printf("Uninstalling package %d/%d, ", i+1, t)
		color.Set(conf.LightColor)
		fmt.Printf(p)
		color.Unset()
		fmt.Println(".")

		if err := pkg.Uninstall(p); err != nil {
			utils.Printe(err.Error())
			continue
		}
	}
}
