package git

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

// Checkout checks out a repo
func Checkout(branch, loc string) error {
	cmd := exec.Command("git", "checkout", branch)
	cmd.Dir = loc

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("Could not git checkout repo!")
	}

	return nil
}

// Clean cleans a repo
func Clean(loc string) error {
	cmd := exec.Command("git", "clean", "-f")
	cmd.Dir = loc

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("Could not git clean repo!")
	}

	return nil
}

// Clone clones a repo
func Clone(url, branch, loc string) error {
	cmd := exec.Command("git", "clone", "--depth", "1", "-b", branch, url)
	cmd.Dir = loc

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("Could not git clone repo!")
	}

	return nil
}

// Diff checks a repo for differences
func Diff(branch, loc string) ([]string, error) {
	cmd := exec.Command("diff", "--name-status", branch)
	cmd.Dir = loc

	b := new(bytes.Buffer)
	cmd.Stdout = b

	err := cmd.Run()
	if err != nil {
		return []string{}, fmt.Errorf("Could not git diff repo!")
	}

	// Make output pretty
	diff := b.String()
	diff = strings.Replace(diff, "A", "Adding", 1)
	diff = strings.Replace(diff, "C", "Copying", 1)
	diff = strings.Replace(diff, "D", "Deleting", 1)
	diff = strings.Replace(diff, "M", "Modifying", 1)
	diff = strings.Replace(diff, "R", "Renaming", 1)

	return strings.Split(diff, "\n"), nil
}

// Fetch fetches a repo
func Fetch(loc string) error {
	cmd := exec.Command("git", "fetch", "--depth", "1")
	cmd.Dir = loc

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("Could not git fetch repo!")
	}

	return nil
}

// Reset resets a repo
func Reset(branch, loc string) error {
	cmd := exec.Command("git", "reset", "--hard", "origin/"+branch)
	cmd.Dir = loc

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("Could not git reset repo!")
	}

	return nil
}
