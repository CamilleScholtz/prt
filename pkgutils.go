package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
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

// pkgBuild builds a port.
func pkgBuild(l string, f, v bool) error {
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
		return trErr(i, "build", portBaseLoc(l))
	}

	return nil
}

// pkgDownload downloads a port sources.
func pkgDownload(l string, v bool) error {
	cmd := exec.Command("/usr/share/prt/pkgmk", "-do")
	cmd.Dir = l
	if v {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}

	if err := cmd.Run(); err != nil {
		i, _ := strconv.Atoi(strings.Split(err.Error(), " ")[2])
		return trErr(i, "download", portBaseLoc(l))
	}

	return nil
}

// pkgInstall installs a package.
func pkgInstall(l string, v bool) error {
	cmd := exec.Command("/usr/share/prt/pkgmk", "-io")
	cmd.Dir = l
	if v {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}

	if err := cmd.Run(); err != nil {
		i, _ := strconv.Atoi(strings.Split(err.Error(), " ")[2])
		return trErr(i, "install", portBaseLoc(l))
	}

	return nil
}

// pkgPostInstall runs a pre-install scripts.
func pkgPostInstall(l string, v bool) error {
	cmd := exec.Command("bash", "./post-install")
	cmd.Dir = l
	if v {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("pkg post-install %s: Something went wrong", portBaseLoc(l))
	}

	return nil
}

// pkgPreInstall runs a pre-install scripts.
func pkgPreInstall(l string, v bool) error {
	cmd := exec.Command("bash", "./pre-install")
	cmd.Dir = l
	if v {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("pkg pre-install %s: Something went wrong", portBaseLoc(l))
	}

	return nil
}

// pkgUninstall uninstalls a package.
func pkgUninstall(p string) error {
	cmd := exec.Command("pkgrm", p)

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("pkg uninstall %s: Something went wrong", p)
	}

	return nil
}

// pkgUnpack unpacks a port sources.
func pkgUnpack(l string, v bool) error {
	cmd := exec.Command("/usr/share/prt/pkgmk", "-eo")
	cmd.Dir = l
	if v {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}

	if err := cmd.Run(); err != nil {
		i, _ := strconv.Atoi(strings.Split(err.Error(), " ")[2])
		return trErr(i, "unpack", portBaseLoc(l))
	}

	return nil
}

// pkgUpdate updates a package.
func pkgUpdate(l string, v bool) error {
	cmd := exec.Command("/usr/share/prt/pkgmk", "-uo")
	cmd.Dir = l
	if v {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}

	if err := cmd.Run(); err != nil {
		i, _ := strconv.Atoi(strings.Split(err.Error(), " ")[2])
		return trErr(i, "update", portBaseLoc(l))
	}

	return nil
}
