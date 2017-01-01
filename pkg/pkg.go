package pkg

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/onodera-punpun/prt/ports"
)

// trErr translates pkgmk error codes to error strings.
func trErr(i int, f, p string) error {
	switch i {
	default:
		return fmt.Errorf("pkg %s %s: Something went wrong", f, p)
	case 2:
		return fmt.Errorf("pkg %s %s: Invalid Pkgfile", f, p)
	case 3:
		return fmt.Errorf("pkg %s %s: Directory missing or missing read/write permission", f, p)
	case 4:
		return fmt.Errorf("pkg %s %s: Could not download source", f, p)
	case 5:
		return fmt.Errorf("pkg %s %s: Could not unpack source", f, p)
	case 6:
		return fmt.Errorf("pkg %s %s: Md5sum verification failed", f, p)
	case 7:
		return fmt.Errorf("pkg %s %s: Footprint check failed", f, p)
	case 8:
		return fmt.Errorf("pkg %s %s: Error while running build()", f, p)
	case 10:
		return fmt.Errorf("pkg %s %s: Signature verification failed", f, p)
	}
}

// Build builds a port.
func Build(l string, f, v bool) error {
	var cmd *exec.Cmd
	if f {
		cmd = exec.Command("/usr/share/prt/pkgmk", "-bo", "-f")
	} else {
		cmd = exec.Command("/usr/share/prt/pkgmk", "-bo")
	}
	cmd.Dir = l
	if v {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}

	if err := cmd.Run(); err != nil {
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

	if err := cmd.Run(); err != nil {
		i, _ := strconv.Atoi(strings.Split(err.Error(), " ")[2])
		return trErr(i, "download", ports.BaseLoc(l))
	}

	return nil
}

// Install installs a package.
func Install(l string, v bool) error {
	cmd := exec.Command("/usr/share/prt/pkgmk", "-io")
	cmd.Dir = l
	if v {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}

	if err := cmd.Run(); err != nil {
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

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("pkg post-install %s: Something went wrong", ports.BaseLoc(l))
	}

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

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("pkg pre-install %s: Something went wrong", ports.BaseLoc(l))
	}

	return nil
}

// Uninstall uninstalls a package.
func Uninstall(p string) error {
	cmd := exec.Command("pkgrm", p)

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("pkg uninstall %s: Something went wrong", p)
	}

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

	if err := cmd.Run(); err != nil {
		i, _ := strconv.Atoi(strings.Split(err.Error(), " ")[2])
		return trErr(i, "unpack", ports.BaseLoc(l))
	}

	return nil
}

// Update updates a package.
func Update(l string, v bool) error {
	cmd := exec.Command("/usr/share/prt/pkgmk", "-uo")
	cmd.Dir = l
	if v {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}

	if err := cmd.Run(); err != nil {
		i, _ := strconv.Atoi(strings.Split(err.Error(), " ")[2])
		return trErr(i, "update", ports.BaseLoc(l))
	}

	return nil
}
