package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"sort"
	"strings"
)

// gitCheckout checks out a repo.
func gitCheckout(b, l string) error {
	cmd := exec.Command("git", "checkout", b)
	cmd.Dir = l

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("git checkout %s: Something went wrong", l)
	}

	return nil
}

// gitClean cleans a repo.
func gitClean(l string) error {
	cmd := exec.Command("git", "clean", "-f")
	cmd.Dir = l

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("git clean %s: Something went wrong", l)
	}

	return nil
}

// gitClone clones a repo.
func gitClone(u, b, l string) error {
	cmd := exec.Command("git", "clone", "--depth", "1", "-b", b, u, l)

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("git clone %s: Something went wrong", u)
	}

	return nil
}

// gitDiff checks a repo for differences.
func gitDiff(b, l string) ([]string, error) {
	cmd := exec.Command("git", "diff", "--name-status", "--diff-filter", "ACDMR", "origin/"+b)
	cmd.Dir = l
	bb := new(bytes.Buffer)
	cmd.Stdout = bb

	if err := cmd.Run(); err != nil {
		return []string{}, fmt.Errorf("git diff %s: Something went wrong", l)
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

// gitFetch fetches a repo.
func gitFetch(l string) error {
	cmd := exec.Command("git", "fetch", "--depth", "1")
	cmd.Dir = l

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("git fetch %s: Something went wrong", l)
	}

	return nil
}

// gitReset resets a repo.
func gitReset(b, l string) error {
	cmd := exec.Command("git", "reset", "--hard", "origin/"+b)
	cmd.Dir = l

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("git reset %s: Something went wrong", l)
	}

	return nil
}
