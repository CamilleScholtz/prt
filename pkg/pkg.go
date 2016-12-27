package pkg

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/onodera-punpun/prt/config"
)

// Load config.
var c = config.Load()

// Build builds a port.
func Build(stdout bool) error {
	// TODO: I'm pretty sure the -f can cause some issues
	// what I want is this function to ONLY build a port
	// and -f can cause it to also update n' shit?
	cmd := exec.Command("/usr/share/prt/pkgmk", "-f")
	if stdout {
		cmd.Stderr = os.Stdout
	}

	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

// Download downloads a port sources.
func Download(stdout bool) error {
	cmd := exec.Command("/usr/share/prt/pkgmk", "-do")
	if stdout {
		cmd.Stderr = os.Stdout
	}

	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

// Extract extracts a port sources.
func Extract(stdout bool) error {
	cmd := exec.Command("/usr/share/prt/pkgmk", "-eo")
	if stdout {
		cmd.Stderr = os.Stdout
	}

	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

// Install installs a port.
func Install(port string, stdout bool) error {
	// Get and fix location from config.
	loc := strings.Replace(c.PackageDir, "$REPO", filepath.Dir(port), -1)
	loc = strings.Replace(c.PackageDir, "$NAME", filepath.Base(port), -1)

	cmd := exec.Command("pkgadd", loc)
	if stdout {
		cmd.Stdout = os.Stdout
	}

	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

// Update updates a port.
func Update(stdout bool) error {
	cmd := exec.Command("pkgadd", "-u", "TODO")
	if stdout {
		cmd.Stderr = os.Stdout
	}

	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}
