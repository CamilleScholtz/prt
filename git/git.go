package git

import (
	"bytes"
	"fmt"
	"os/exec"
	"sort"
	"strings"
)

// Checkout checks out a repo.
func Checkout(b, l string) error {
	cmd := exec.Command("git", "checkout", b)
	cmd.Dir = l

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("git checkout %s: Something went wrong", l)
	}

	return nil
}

// Clean cleans a repo.
func Clean(l string) error {
	cmd := exec.Command("git", "clean", "-f")
	cmd.Dir = l

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("git clean %s: Something went wrong", l)
	}

	return nil
}

// Clone clones a repo.
func Clone(u, b, l string) error {
	cmd := exec.Command("git", "clone", "--depth", "1", "-b", b, u)
	cmd.Dir = l

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("git clone %s: Something went wrong", l)
	}

	return nil
}

// Diff checks a repo for differences.
func Diff(b, l string) ([]string, error) {
	cmd := exec.Command("git", "diff", "--name-status", "origin/"+b)
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

// Fetch fetches a repo.
func Fetch(l string) error {
	cmd := exec.Command("git", "fetch", "--depth", "1")
	cmd.Dir = l

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("git fetch %s: Something went wrong", l)
	}

	return nil
}

// Reset resets a repo.
func Reset(b, l string) error {
	cmd := exec.Command("git", "reset", "--hard", "origin/"+b)
	cmd.Dir = l

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("git reset %s: Something went wrong", l)
	}

	return nil
}
