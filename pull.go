package main

import (
	"fmt"
	"os"

	"github.com/chiyouhen/getopt"
	"github.com/fatih/color"
)


func Pull(args []string) {
	// Define opts
	shortopts := "h"
	longopts := []string{
		"--help",
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
			fmt.Println("Usage: prt pull [arguments] [repos]")
			fmt.Println("")
			fmt.Println("arguments:")
			fmt.Println("  -h,   --help            print help and exit")
			os.Exit(0)
		}
	}

	for i, repo := range Config.Pull {
		color.Set(color.FgYellow, color.Bold)
		fmt.Println(i)
		fmt.Println("Pulling in repo " + i + "/" + string(len(Config.Pull)) + ", " + repo.Name + ".")
		color.Unset()
	}
}
