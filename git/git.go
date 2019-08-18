package git

import (
	"bytes"
	"fmt"
	"os/exec"
	"sort"
	"strings"
)

type Repo struct {
	Location string
	URL      string
	Branch   string
}

// Checkout checks out a repo.
func (r Repo) Checkout() error {
	cmd := exec.Command("git", "checkout", r.Branch)
	cmd.Dir = r.Location

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("git checkout %s: Something went wrong", r.Location)
	}

	return nil
}

// Clean cleans a repo.
func (r Repo) Clean() error {
	cmd := exec.Command("git", "clean", "-f")
	cmd.Dir = r.Location

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("git clean %s: Something went wrong", r.Location)
	}

	return nil
}

// Clone clones a repo.
func (r Repo) Clone() error {
	cmd := exec.Command("git", "clone", "--depth", "1", "-b", r.Branch, r.URL,
		r.Location)

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("git clone %s: Something went wrong", r.URL)
	}

	return nil
}

// Diff checks a repo for differences.
func (r Repo) Diff() ([]string, error) {
	cmd := exec.Command("git", "diff", "--name-status", "--diff-filter",
		"ACDM", "origin/"+r.Branch)
	cmd.Dir = r.Location
	bb := new(bytes.Buffer)
	cmd.Stdout = bb

	if err := cmd.Run(); err != nil {
		return []string{}, fmt.Errorf("git diff %s: Something went wrong", r.
			Location)
	}

	d := bb.String()
	if len(d) < 1 {
		return []string{}, nil
	}

	// Make output pretty.
	// TODO: This prints Deleted when it should be Added.
	// TODO: Enable Renamed.
	d = strings.Replace(d, "A\t", "Added ", -1)
	d = strings.Replace(d, "C\t", "Copied ", -1)
	d = strings.Replace(d, "D\t", "Deleted ", -1)
	d = strings.Replace(d, "M\t", "Modiefied ", -1)
	//d = strings.Replace(d, "R\t", "Renamed ", -1)
	dl := strings.Split(d, "\n")
	sort.Strings(dl)

	return dl[1:], nil
}

// Fetch fetches a repo.
func (r Repo) Fetch() error {
	cmd := exec.Command("git", "fetch", "--depth", "1")
	cmd.Dir = r.Location

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("git fetch %s: Something went wrong", r.Location)
	}

	return nil
}

// Reset resets a repo.
func (r Repo) Reset() error {
	cmd := exec.Command("git", "reset", "--hard", "origin/"+r.Branch)
	cmd.Dir = r.Location

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("git reset %s: Something went wrong", r.Location)
	}

	return nil
}
