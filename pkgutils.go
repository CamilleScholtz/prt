package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

type pkg struct {
	loc string
}

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

// build builds a port.
func (p pkg) build(f, v bool) error {
	var cmd *exec.Cmd
	if f {
		cmd = exec.Command("/usr/share/prt/pkgmk", "-bo", "-f")
	} else {
		cmd = exec.Command("/usr/share/prt/pkgmk", "-bo")
	}
	cmd.Dir = p.loc
	if v {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}

	if err := cmd.Run(); err != nil {
		i, _ := strconv.Atoi(strings.Split(err.Error(), " ")[2])
		return trErr(i, "build", portBaseLoc(p.loc))
	}

	return nil
}

// download downloads a port sources.
func (p pkg) download(v bool) error {
	cmd := exec.Command("/usr/share/prt/pkgmk", "-do")
	cmd.Dir = p.loc
	if v {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}

	if err := cmd.Run(); err != nil {
		i, _ := strconv.Atoi(strings.Split(err.Error(), " ")[2])
		return trErr(i, "download", portBaseLoc(p.loc))
	}

	return nil
}

// install installs a package.
func (p pkg) install(v bool) error {
	cmd := exec.Command("/usr/share/prt/pkgmk", "-io")
	cmd.Dir = p.loc
	if v {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}

	if err := cmd.Run(); err != nil {
		i, _ := strconv.Atoi(strings.Split(err.Error(), " ")[2])
		return trErr(i, "install", portBaseLoc(p.loc))
	}

	return nil
}

// post runs a pre-install scripts.
func (p pkg) post(v bool) error {
	cmd := exec.Command("bash", "./post-install")
	cmd.Dir = p.loc
	if v {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("pkg post-install %s: Something went wrong", portBaseLoc(p.loc))
	}

	return nil
}

// pre runs a pre-install scripts.
func (p pkg) pre(v bool) error {
	cmd := exec.Command("bash", "./pre-install")
	cmd.Dir = p.loc
	if v {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("pkg pre-install %s: Something went wrong", portBaseLoc(p.loc))
	}

	return nil
}

// uninstall uninstalls a package.
// TODO
func pkgUninstall(todo string) error {
	cmd := exec.Command("pkgrm", todo)

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("pkg uninstall %s: Something went wrong", todo)
	}

	return nil
}

// unpack unpacks a port sources.
func (p pkg) unpack(v bool) error {
	cmd := exec.Command("/usr/share/prt/pkgmk", "-eo")
	cmd.Dir = p.loc
	if v {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}

	if err := cmd.Run(); err != nil {
		i, _ := strconv.Atoi(strings.Split(err.Error(), " ")[2])
		return trErr(i, "unpack", portBaseLoc(p.loc))
	}

	return nil
}

// update updates a package.
func (p pkg) update(v bool) error {
	cmd := exec.Command("/usr/share/prt/pkgmk", "-uo")
	cmd.Dir = p.loc
	if v {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}

	if err := cmd.Run(); err != nil {
		i, _ := strconv.Atoi(strings.Split(err.Error(), " ")[2])
		return trErr(i, "update", portBaseLoc(p.loc))
	}

	return nil
}
