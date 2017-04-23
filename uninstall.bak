package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/go2c/optparse"
)

// uninstall uninstalls packages.
func uninstall(args []string) {
	// Define valid arguments.
	o := optparse.New()
	argh := o.Bool("help", 'h', false)

	// Parse arguments.
	vals, err := o.Parse(args)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Invaild argument, use -h for a list of arguments!")
		os.Exit(1)
	}

	// Print help.
	if *argh {
		fmt.Println("Usage: prt uninstall [arguments] [packages]")
		fmt.Println("")
		fmt.Println("arguments:")
		fmt.Println("  -h,   --help            print help and exit")
		os.Exit(0)
	}

	// This command needs a value.
	if len(vals) == 0 {
		fmt.Fprintln(os.Stderr, "Please specify a package!")
		os.Exit(1)
	}

	t := len(vals)
	for i, p := range vals {
		fmt.Printf("Uninstalling package %d/%d, ", i+1, t)
		color.Set(config.LightColor)
		fmt.Printf(p)
		color.Unset()
		fmt.Println(".")

		if err := pkgUninstall(p); err != nil {
			printe(err.Error())
			continue
		}
	}
}
