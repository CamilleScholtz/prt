package main

import (
	"fmt"
	"os"
	"path"

	"github.com/go2c/optparse"
	"github.com/onodera-punpun/go-utils/array"
	git "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
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

		g, err := git.PlainOpen(l)
		if err != nil && err != git.ErrRepositoryNotExists {
			fmt.Fprintf(os.Stderr, "%s%s\n", warning(config.WarningChar), err)
			continue
		}

		// Check if location exists, clone if it doesn't.
		if err == git.ErrRepositoryNotExists {
			g, err = git.PlainClone(l, false, &git.CloneOptions{
				URL:           r.URL,
				ReferenceName: plumbing.ReferenceName("refs/heads/" + r.Branch),
				SingleBranch:  true,
				Depth:         1,
				Progress:      nil,
			})
			if err != nil {
				fmt.Fprintf(os.Stderr, "%s%s\n", warning(config.WarningChar),
					err)
			}

			fmt.Printf("%sInitialized\n", dark(config.IndentChar))

			continue
		}

		if err := g.Fetch(&git.FetchOptions{
			Depth:    1,
			Progress: nil,
		}); err != nil && err != git.NoErrAlreadyUpToDate {
			fmt.Fprintf(os.Stderr, "%s%s\n", warning(config.WarningChar), err)
			continue
		}

		wt, err := g.Worktree()
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s%s\n", warning(config.WarningChar), err)
			continue
		}

		if err := wt.Clean(&git.CleanOptions{
			Dir: true,
		}); err != nil {
			fmt.Fprintf(os.Stderr, "%s%s\n", warning(config.WarningChar), err)
			continue
		}
		if err := wt.Reset(&git.ResetOptions{
			Mode: git.HardReset,
		}); err != nil {
			fmt.Fprintf(os.Stderr, "%s%s\n", warning(config.WarningChar), err)
			continue
		}

		sl, err := wt.Status()
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s%s\n", warning(config.WarningChar), err)
		}
		for _, s := range sl {
			fmt.Println("%s%s\n", dark(config.IndentChar), s)
		}
	}

	return nil
}
