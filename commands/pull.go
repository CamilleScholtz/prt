package commands

import (
	"fmt"
	"os"

	"github.com/chiyouhen/getopt"
	"github.com/fatih/color"
	"github.com/onodera-punpun/prt/config"
	"github.com/onodera-punpun/prt/git"
	"github.com/onodera-punpun/prt/utils"
)

// Pull pulls in ports
func Pull(args []string) {
	// Define opts
	shortopts := "h"
	longopts := []string{
		"--help",
	}

	// Read out opts
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

	// Count total repos that need to be pulled
	var t int
	if len(vals) == 0 {
		t = len(config.Struct.Pull)
	} else {
		t = len(vals)
	}

	for name, repo := range config.Struct.Pull {
		// Skip repos if needed
		if len(vals) != 0 {
			if !utils.StringInList(name, vals) {
				return
			}
		}
		i++

		// Print some info
		fmt.Printf("Pulling in repo %d/%d, ", i, t)
		color.Set(config.Struct.LightColor)
		fmt.Printf(name)
		color.Unset()
		fmt.Println(".")

		loc := config.Struct.PortDir + "/" + name

		// Check if location exists, clone if it doesn't
		_, err := os.Stat(loc)
		if err != nil {
			err := git.Clone(repo.URL, repo.Branch, loc)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			}
			return
		}

		err = git.Checkout(repo.Branch, loc)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
		err = git.Fetch(loc)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}

		// Output changes
		// TODO: Does this actually output anything?
		diff := git.Diff(repo.Branch, loc)
		for _, l := range diff {
			color.Set(config.Struct.DarkColor)
			fmt.Print(config.Struct.IndentChar)
			color.Unset()
			fmt.Println(l)
		}

		err = git.Clean(loc)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
		err = git.Reset(repo.Branch, loc)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
	}
}
