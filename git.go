package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"sort"
	"strings"
)

type git struct {
	Branch string
	Loc    string
	URL    string
}

// TODO: Use exec.LookPath maybe.
// TODO: Use a Go package for git stuff maybe.

// checkout checks out a repo.
func (g git) checkout() error {
	cmd := exec.Command("git", "checkout", g.Branch)
	cmd.Dir = g.Loc

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("git checkout %s: Something went wrong", g.Loc)
	}

	return nil
}

// clean cleans a repo.
func (g git) clean() error {
	cmd := exec.Command("git", "clean", "-f")
	cmd.Dir = g.Loc

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("git clean %s: Something went wrong", g.Loc)
	}

	return nil
}

// clone clones a repo.
func (g git) clone() error {
	cmd := exec.Command("git", "clone", "--depth", "1", "-b", g.Branch, g.URL,
		g.Loc)

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("git clone %s: Something went wrong", g.URL)
	}

	return nil
}

// diff checks a repo for differences.
func (g git) diff() ([]string, error) {
	cmd := exec.Command("git", "diff", "--name-status", "--diff-filter",
		"ACDMR", "origin/"+g.Branch)
	cmd.Dir = g.Loc
	bb := new(bytes.Buffer)
	cmd.Stdout = bb

	if err := cmd.Run(); err != nil {
		return []string{}, fmt.Errorf("git diff %s: Something went wrong",
			g.Loc)
	}

	d := bb.String()
	if len(d) < 1 {
		return []string{}, nil
	}

	// Make output pretty.
	d = strings.Replace(d, "A\t", "Adding ", -1)
	d = strings.Replace(d, "C\t", "Copying ", -1)
	d = strings.Replace(d, "D\t", "Deleting ", -1)
	d = strings.Replace(d, "M\t", "Editing ", -1)
	d = strings.Replace(d, "R\t", "Renaming ", -1)
	dl := strings.Split(d, "\n")
	sort.Strings(dl)

	return dl[1:], nil
}

// fetch fetches a repo.
func (g git) fetch() error {
	cmd := exec.Command("git", "fetch", "--depth", "1")
	cmd.Dir = g.Loc

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("git fetch %s: Something went wrong", g.Loc)
	}

	return nil
}

// reset resets a repo.
func (g git) reset() error {
	cmd := exec.Command("git", "reset", "--hard", "origin/"+g.Branch)
	cmd.Dir = g.Loc

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("git reset %s: Something went wrong", g.Loc)
	}

	return nil
}
