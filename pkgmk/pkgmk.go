package pkgmk

import (
	"fmt"
	"os"
	"os/exec"
)

// Download TODO
func Download(loc string) error {
	os.Chdir(loc)

	cmd := "pkgmk"
	args := []string{"-d"}
	if err := exec.Command(cmd, args...).Run(); err != nil {
		return fmt.Errorf("Could not git checkout repo!")
	}

	return nil
}
