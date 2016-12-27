package pkg

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/onodera-punpun/prt/config"
)

// Load config.
var c = config.Load()

// pkgmkErr translates pkgmk error codes to error strings.
func pkgmkErr(i int, p string) error {
	switch i {
	default:
		return fmt.Errorf("pkgmk %s: Something went wrong", p)
	case 2:
		return fmt.Errorf("pkgmk %s: Invalid Pkgfile", p)
	case 3:
		return fmt.Errorf("pkgmk %s: Directory missing or missing read/write permission", p)
	case 4:
		return fmt.Errorf("pkgmk %s: Could not download source", p)
	case 5:
		return fmt.Errorf("pkgmk %s: Could not unpack source", p)
	case 6:
		return fmt.Errorf("pkgmk %s: Md5sum verification failed", p)
	case 7:
		return fmt.Errorf("pkgmk %s: Footprint check failed", p)
	case 8:
		return fmt.Errorf("pkgmk %s: Error while running build()", p)
	case 10:
		return fmt.Errorf("pkgmk %s: Signature verification failed", p)
	}
}

// Build builds a port.
func Build(v bool) error {
	// TODO: I'm pretty sure the -f can cause some issues
	// what I want is this function to ONLY build a port
	// and -f can cause it to also update n' shit?
	cmd := exec.Command("/usr/share/prt/pkgmk", "-f")
	if v {
		cmd.Stderr = os.Stdout
	}

	err := cmd.Run()
	if err != nil {
		i, _ := strconv.Atoi(strings.Split(err.Error(), " ")[2])
		p, _ := os.Getwd()
		return pkgmkErr(i, p)
	}

	return nil
}

// Download downloads a port sources.
func Download(v bool) error {
	cmd := exec.Command("/usr/share/prt/pkgmk", "-do")
	if v {
		cmd.Stderr = os.Stdout
	}

	err := cmd.Run()
	if err != nil {
		i, _ := strconv.Atoi(strings.Split(err.Error(), " ")[2])
		p, _ := os.Getwd()
		return pkgmkErr(i, p)
	}

	return nil
}

// Install installs a port.
func Install(p string, v bool) error {
	// Get and fix location from config.
	loc := strings.Replace(c.PackageDir, "$REPO", filepath.Dir(p), -1)
	loc = strings.Replace(c.PackageDir, "$NAME", filepath.Base(p), -1)

	cmd := exec.Command("pkgadd", loc)
	if v {
		cmd.Stdout = os.Stdout
	}

	err := cmd.Run()
	if err != nil {
		i, _ := strconv.Atoi(strings.Split(err.Error(), " ")[2])
		p, _ := os.Getwd()
		return pkgmkErr(i, p)
	}

	return nil
}

// Unpack unpacks a port sources.
func Unpack(v bool) error {
	cmd := exec.Command("/usr/share/prt/pkgmk", "-eo")
	if v {
		cmd.Stderr = os.Stdout
	}

	err := cmd.Run()
	if err != nil {
		i, _ := strconv.Atoi(strings.Split(err.Error(), " ")[2])
		p, _ := os.Getwd()
		return pkgmkErr(i, p)
	}

	return nil
}

// Update updates a port.
func Update(v bool) error {
	cmd := exec.Command("pkgadd", "-u", "TODO")
	if v {
		cmd.Stderr = os.Stdout
	}

	err := cmd.Run()
	if err != nil {
		i, _ := strconv.Atoi(strings.Split(err.Error(), " ")[2])
		p, _ := os.Getwd()
		return pkgmkErr(i, p)
	}

	return nil
}
