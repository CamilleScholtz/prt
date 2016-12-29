package pkgmk

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/onodera-punpun/prt/config"
	"github.com/onodera-punpun/prt/ports"
)

// Load config.
var c = config.Load()

// trErr translates pkgmk error codes to error strings.
func trErr(i int, f, p string) error {
	switch i {
	default:
		return fmt.Errorf("pkgmk %s %s: Something went wrong", f, p)
	case 2:
		return fmt.Errorf("pkgmk %s %s: Invalid Pkgfile", f, p)
	case 3:
		return fmt.Errorf("pkgmk %s %s: Directory missing or missing read/write permission", f, p)
	case 4:
		return fmt.Errorf("pkgmk %s %s: Could not download source", f, p)
	case 5:
		return fmt.Errorf("pkgmk %s %s: Could not unpack source", f, p)
	case 6:
		return fmt.Errorf("pkgmk %s %s: Md5sum verification failed", f, p)
	case 7:
		return fmt.Errorf("pkgmk %s %s: Footprint check failed", f, p)
	case 8:
		return fmt.Errorf("pkgmk %s %s: Error while running build()", f, p)
	case 10:
		return fmt.Errorf("pkgmk %s %s: Signature verification failed", f, p)
	}
}

// Build builds a port.
func Build(l string, v bool) error {
	cmd := exec.Command("/usr/share/prt/pkgmk", "-bo")
	cmd.Dir = l
	if v {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}

	err := cmd.Run()
	if err != nil {
		i, _ := strconv.Atoi(strings.Split(err.Error(), " ")[2])
		return trErr(i, "build", ports.BaseLoc(l))
	}

	return nil
}

// Download downloads a port sources.
func Download(l string, v bool) error {
	cmd := exec.Command("/usr/share/prt/pkgmk", "-do")
	cmd.Dir = l
	if v {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}

	err := cmd.Run()
	if err != nil {
		i, _ := strconv.Atoi(strings.Split(err.Error(), " ")[2])
		return trErr(i, "download", ports.BaseLoc(l))
	}

	return nil
}

// Install installs a port.
func Install(l string, v bool) error {
	cmd := exec.Command("/usr/share/prt/pkgmk", "-io")
	cmd.Dir = l
	if v {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}

	err := cmd.Run()
	if err != nil {
		i, _ := strconv.Atoi(strings.Split(err.Error(), " ")[2])
		return trErr(i, "install", ports.BaseLoc(l))
	}

	return nil
}

// PostInstall runs a pre-install scripts.
func PostInstall(l string, v bool) error {
	cmd := exec.Command("bash", "./post-install")
	cmd.Dir = l
	if v {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("pkgmk post-install %s: Something went wrong", ports.BaseLoc(l))
	}

	return nil

	return nil
}

// PreInstall runs a pre-install scripts.
func PreInstall(l string, v bool) error {
	cmd := exec.Command("bash", "./pre-install")
	cmd.Dir = l
	if v {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("pkgmk pre-install %s: Something went wrong", ports.BaseLoc(l))

	}

	return nil

	return nil
}

// Unpack unpacks a port sources.
func Unpack(l string, v bool) error {
	cmd := exec.Command("/usr/share/prt/pkgmk", "-eo")
	cmd.Dir = l
	if v {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}

	err := cmd.Run()
	if err != nil {
		i, _ := strconv.Atoi(strings.Split(err.Error(), " ")[2])
		return trErr(i, "unpack", ports.BaseLoc(l))
	}

	return nil
}

// Update updates a port.
func Update(l string, v bool) error {
	cmd := exec.Command("/usr/share/prt/pkgmk", "-uo")
	cmd.Dir = l
	if v {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}

	err := cmd.Run()
	if err != nil {
		i, _ := strconv.Atoi(strings.Split(err.Error(), " ")[2])
		return trErr(i, "update", ports.BaseLoc(l))
	}

	return nil
}
