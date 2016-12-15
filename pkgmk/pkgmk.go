package pkgmk

import "os/exec"

// Download TODO
func Download(loc string) error {
	args = []string{"checkout", repo.Branch}
	if err := exec.Command(cmd, args...).Run(); err != nil {
		return "Could not git checkout repo!"
	}

	return nil
}
