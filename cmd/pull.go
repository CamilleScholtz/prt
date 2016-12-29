package cmd

import (
	"fmt"
	"os"

	"github.com/chiyouhen/getopt"
	"github.com/fatih/color"
	"github.com/onodera-punpun/prt/git"
	"github.com/onodera-punpun/prt/ports"
	"github.com/onodera-punpun/prt/utils"
)

// Pull pulls in ports.
func Pull(args []string) {
	// Define opts.
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

	for _, opt := range opts {
		switch opt[0] {
		case "-h", "--help":
			fmt.Println("Usage: prt pull [arguments] [repos]")
			fmt.Println("")
			fmt.Println("arguments:")
			fmt.Println("  -h,   --help            print help and exit")
			os.Exit(0)
		}
	}

	// Count total repos that need to be pulled.
	var t int
	if len(vals) == 0 {
		t = len(c.Pull)
	} else {
		t = len(vals)
	}

	// TODO: Actually learn git and check if all these commands are needed.
	// TODO: Sort this?
	// also check if branch is needed for these commands.
	for n, r := range c.Pull {
		// Skip repos if needed.
		if len(vals) != 0 {
			if !utils.StringInList(n, vals) {
				continue
			}
		}
		i++

		fmt.Printf("Pulling in repo %d/%d, ", i, t)
		color.Set(c.LightColor)
		fmt.Printf(n)
		color.Unset()
		fmt.Println(".")

		l := ports.FullLoc(n)

		// Check if location exists, clone if it doesn't.
		_, err := os.Stat(l)
		if err != nil {
			err := git.Clone(r.URL, r.Branch, l)
			if err != nil {
				utils.Printe(err.Error())
			}
			continue
		}

		err = git.Checkout(r.Branch, l)
		if err != nil {
			utils.Printe(err.Error())
			continue
		}
		err = git.Fetch(l)
		if err != nil {
			utils.Printe(err.Error())
			continue
		}

		// Print changes.
		dl, err := git.Diff(r.Branch, l)
		if err != nil {
			utils.Printe(err.Error())
			continue
		}
		for _, d := range dl {
			color.Set(c.DarkColor)
			fmt.Print(c.IndentChar)
			color.Unset()
			fmt.Println(d)
		}

		err = git.Clean(l)
		if err != nil {
			utils.Printe(err.Error())
			continue
		}
		err = git.Reset(r.Branch, l)
		if err != nil {
			utils.Printe(err.Error())
			continue
		}
	}
}
