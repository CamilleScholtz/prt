package main

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"regexp"
	"strconv"
	"strings"
)

type pkg struct {
	Loc string
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
	cmd.Dir = p.Loc
	if v {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}

	if err := cmd.Run(); err != nil {
		i, _ := strconv.Atoi(strings.Split(err.Error(), " ")[2])
		return trErr(i, "build", portBaseLoc(p.Loc))
	}

	return nil
}

// download downloads a port sources.
func (p pkg) download(v bool) error {
	// Read out Pkgfile.
	f, err := readPkgfile(path.Join(p.Loc, "Pkgfile"))
	if err != nil {
		return err
	}

	// Get sources.
	s, err := f.variableSource("source")
	if err != nil {
		return err
	}
	sl := strings.Fields(s)

	// Download sources.
	for _, s := range sl {
		f := path.Join(config.SrcDir, path.Base(s))

		// Continue if file has already been downloaded.
		if _, err := os.Stat(f); err == nil {
			continue
		}

		// Continue if file is not an URL.
		r := regexp.MustCompile("^(http|https|ftp|file)://")
		if !r.MatchString(s) {
			continue
		}

		cmd := exec.Command("curl", "-L", "-#", "--fail", "--ftp-pasv", "-C", "-", "-o", f+".partial", s)
		if v {
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
		}

		printi("Downloading " + s)
		if err := cmd.Run(); err != nil {
			return fmt.Errorf("pkg download %s: Something went wrong", portBaseLoc(p.Loc))
		}

		// Remove .partial on completion.
		os.Rename(f+".partial", f)
	}

	return nil
}

// install installs a package.
func (p pkg) install(v bool) error {
	cmd := exec.Command("/usr/share/prt/pkgmk", "-io")
	cmd.Dir = p.Loc
	if v {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}

	if err := cmd.Run(); err != nil {
		i, _ := strconv.Atoi(strings.Split(err.Error(), " ")[2])
		return trErr(i, "install", portBaseLoc(p.Loc))
	}

	return nil
}

// post runs a pre-install scripts.
func (p pkg) post(v bool) error {
	cmd := exec.Command("bash", "./post-install")
	cmd.Dir = p.Loc
	if v {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("pkg post %s: Something went wrong", portBaseLoc(p.Loc))
	}

	return nil
}

// pre runs a pre-install scripts.
func (p pkg) pre(v bool) error {
	cmd := exec.Command("bash", "./pre-install")
	cmd.Dir = p.Loc
	if v {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("pkg pre %s: Something went wrong", portBaseLoc(p.Loc))
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
	cmd.Dir = p.Loc
	if v {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}

	if err := cmd.Run(); err != nil {
		i, _ := strconv.Atoi(strings.Split(err.Error(), " ")[2])
		return trErr(i, "unpack", portBaseLoc(p.Loc))
	}

	return nil
}

// update updates a package.
func (p pkg) update(v bool) error {
	cmd := exec.Command("/usr/share/prt/pkgmk", "-uo")
	cmd.Dir = p.Loc
	if v {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}

	if err := cmd.Run(); err != nil {
		i, _ := strconv.Atoi(strings.Split(err.Error(), " ")[2])
		return trErr(i, "update", portBaseLoc(p.Loc))
	}

	return nil
}
