package commands

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/chiyouhen/getopt"
	"github.com/fatih/color"
	"github.com/onodera-punpun/prt/config"
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

	var i, t int

	// Count total repos that need to be pulled
	if len(vals) != 0 {
		t = len(vals)
	} else {
		t = len(config.Struct.Pull)
	}

	for name, repo := range config.Struct.Pull {
		// Skip repos if needed
		if len(vals) != 0 {
			if !utils.StringInList(name, vals) {
				continue
			}
		}

		i++

		// Print some info
		fmt.Printf("Pulling in repo %d/%d, ", i, t)
		color.Set(color.FgYellow, color.Bold)
		fmt.Printf(name)
		color.Unset()
		fmt.Println(".")

		// Actually clone/pull port
		cmd := "git"
		loc := "./test" + "/" + name
		if _, err := os.Stat("./test" + "/" + name); err != nil {
			args = []string{"clone", "--depth", "1", "-b", repo.Branch, repo.URL, loc}
			if err := exec.Command(cmd, args...).Run(); err != nil {
				fmt.Fprintln(os.Stderr, "Could not git clone repo!")
				os.Exit(1)
			}
		} else {
			os.Chdir(loc)

			args = []string{"checkout", repo.Branch}
			if err := exec.Command(cmd, args...).Run(); err != nil {
				fmt.Fprintln(os.Stderr, "Could not git checkout repo!")
				os.Exit(1)
			}

			args = []string{"fetch", "--depth", "1"}
			if err := exec.Command(cmd, args...).Run(); err != nil {
				fmt.Fprintln(os.Stderr, "Could not git fetch repo!")
				os.Exit(1)
			}

			// TODO: Prettify this
			// TODO: Make this actually output something
			args = []string{"diff", "--pretty=format:", "--name-status", repo.Branch}
			info := exec.Command(cmd, args...)
			info.Stdout = os.Stdout
			info.Stderr = os.Stderr
			if err := info.Run(); err != nil {
				fmt.Fprintln(os.Stderr, "Could not git diff repo!")
				os.Exit(1)
			}

			args = []string{"clean", "-f"}
			if err := exec.Command(cmd, args...).Run(); err != nil {
				fmt.Fprintln(os.Stderr, "Could not git clean repo!")
				os.Exit(1)
			}

			args = []string{"reset", "--hard", "origin/" + repo.Branch}
			if err := exec.Command(cmd, args...).Run(); err != nil {
				fmt.Fprintln(os.Stderr, "Could not git reset repo!")
				os.Exit(1)
			}
		}
	}
}
