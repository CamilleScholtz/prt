package cmd

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/go2c/optparse"
	"github.com/onodera-punpun/prt/config"
	"github.com/onodera-punpun/prt/git"
	"github.com/onodera-punpun/prt/ports"
	"github.com/onodera-punpun/prt/utils"
)

// Pull pulls in ports.
func Pull(args []string) {
	// Load config.
	conf := config.Load()

	// Define valid arguments.
	argh := optparse.Bool("help", 'h', false)

	// Parse arguments.
	vals, err := optparse.Parse(args)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Invaild argument, use -h for a list of arguments!")
		os.Exit(1)
	}

	if *argh {
		fmt.Println("Usage: prt pull [arguments] [repos]")
		fmt.Println("")
		fmt.Println("arguments:")
		fmt.Println("  -h,   --help            print help and exit")
		os.Exit(0)
	}

	// Count total repos that need to be pulled.
	var t int
	if len(vals) == 0 {
		t = len(conf.Pull)
	} else {
		t = len(vals)
	}

	// TODO: Actually learn git and check if all these commands are needed.
	// TODO: Sort this?
	// also check if branch is needed for these commands.
	var i int
	for n, r := range conf.Pull {
		// Skip repos if needed.
		if len(vals) != 0 {
			if !utils.StringInList(n, vals) {
				continue
			}
		}
		i++

		fmt.Printf("Pulling in repo %d/%d, ", i, t)
		color.Set(conf.LightColor)
		fmt.Printf(n)
		color.Unset()
		fmt.Println(".")

		l := ports.FullLoc(n)

		// Check if location exists, clone if it doesn't.
		if _, err := os.Stat(l); err != nil {
			err := git.Clone(r.URL, r.Branch, l)
			if err != nil {
				utils.Printe(err.Error())
			}
			continue
		}

		if err := git.Checkout(r.Branch, l); err != nil {
			utils.Printe(err.Error())
			continue
		}
		if err := git.Fetch(l); err != nil {
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
			color.Set(conf.DarkColor)
			fmt.Print(conf.IndentChar)
			color.Unset()
			fmt.Println(d)
		}

		if err := git.Clean(l); err != nil {
			utils.Printe(err.Error())
			continue
		}
		if err := git.Reset(r.Branch, l); err != nil {
			utils.Printe(err.Error())
			continue
		}
	}
}
