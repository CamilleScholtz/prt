package commands

import (
	"fmt"
	"os"

	"github.com/chiyouhen/getopt"
)

// Prov lists TODO
func Prov(args []string) {
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
			fmt.Println("Usage: prt prov [arguments] [queries]")
			fmt.Println("")
			fmt.Println("arguments:")
			fmt.Println("  -h,   --help            print help and exit")
			os.Exit(0)
		}
	}
}
