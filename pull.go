package main

import (
	"fmt"
	"os"
	"path"

	"github.com/go2c/optparse"
	"github.com/onodera-punpun/go-utils/array"
	"github.com/onodera-punpun/prt/git"
)

// pullCommand pulls in ports.
func pullCommand(input []string) error {
	// Define valid arguments.
	o := optparse.New()
	argh := o.Bool("help", 'h', false)

	// Parse arguments.
	vals, err := o.Parse(input)
	if err != nil {
		return fmt.Errorf("invaild argument, use `-h` for a list of arguments")
	}

	// Print help.
	if *argh {
		fmt.Println("Usage: prt pull [arguments] [repos]")
		fmt.Println("")
		fmt.Println("arguments:")
		fmt.Println("  -h,   --help            print help and exit")

		return nil
	}

	// Count repos that need to be pulled.
	var t int
	if len(vals) == 0 {
		t = len(config.Repo)
	} else {
		t = len(vals)
	}

	// TODO: Actually learn git and check if all these commands are needed.
	var i int
	for n, r := range config.Repo {
		// Skip repos if needed.
		if len(vals) != 0 {
			if !array.ContainsString(vals, n) {
				continue
			}
		}
		i++

		fmt.Printf("Pulling in repo %d/%d, %s.\n", i, t, light(n))

		l := path.Join(config.PrtDir, n)
		g := git.Repo{
			Location: l,
			URL:      r.URL,
			Branch:   r.Branch,
		}

		// Check if location exists, clone if it doesn't.
		if _, err := os.Stat(l); err != nil {
			if err := g.Clone(); err != nil {
				fmt.Fprintf(os.Stderr, "%s%s\n", warning(config.WarningChar),
					err)
			}
			continue
		}

		if err := g.Checkout(); err != nil {
			fmt.Fprintf(os.Stderr, "%s%s\n", warning(config.WarningChar), err)
			continue
		}
		if err := g.Fetch(); err != nil {
			fmt.Fprintf(os.Stderr, "%s%s\n", warning(config.WarningChar), err)
			continue
		}

		// Print changes.
		dl, err := g.Diff()
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s%s\n", warning(config.WarningChar), err)
			continue
		}
		for _, d := range dl {
			fmt.Printf("%s%s\n", dark(config.IndentChar), d)
		}

		if err := g.Clean(); err != nil {
			fmt.Fprintf(os.Stderr, "%s%s\n", warning(config.WarningChar), err)
			continue
		}
		if err := g.Reset(); err != nil {
			fmt.Fprintf(os.Stderr, "%s%s\n", warning(config.WarningChar), err)
			continue
		}
	}

	return nil
}
