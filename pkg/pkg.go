package pkg

import (
	"fmt"
	"os"
	"os/exec"
)

// Build builds a port
func Build(stdout bool) error {
	// TODO: I'm pretty sure the -f can cause someissues
	// what I want is this function to ONLY build a port
	// and -f can cause it to also update n shit?
	cmd := exec.Command("pkgmk", "-f")
	if stdout {
		cmd.Stderr = os.Stdout
	}

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("Could not download sources!")
	}

	return nil
}

// Download downloads a port sources
func Download(stdout bool) error {
	cmd := exec.Command("pkgmk", "-do")
	if stdout {
		cmd.Stderr = os.Stdout
	}

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("Could not download sources!")
	}

	return nil
}

// Install installs a port
func Install(stdout bool) error {
	cmd := exec.Command("pkgadd", "TODO")
	if stdout {
		cmd.Stdout = os.Stdout
	}

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("Could not install package!")
	}

	return nil
}

// Update updates a port
func Update(stdout bool) error {
	cmd := exec.Command("pkgadd", "-u", "TODO")
	if stdout {
		cmd.Stderr = os.Stdout
	}

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("Could not update package!")
	}

	return nil
}
