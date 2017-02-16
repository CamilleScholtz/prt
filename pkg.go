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
func (p pkgfile) build(f, v bool) error {
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
func (p pkgfile) download(v bool) error {
	// Get sources.
	s, err := p.variableSource("source")
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

		printi("Downloading " + path.Base(s))
		if err := cmd.Run(); err != nil {
			return fmt.Errorf("pkg download %s: Something went wrong", path.Base(s))
		}

		// Remove .partial on completion.
		os.Rename(f+".partial", f)
	}

	return nil
}

// install installs a package.
func (p pkgfile) install(v bool) error {
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
func (p pkgfile) post(v bool) error {
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
func (p pkgfile) pre(v bool) error {
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
func (p pkgfile) unpack(v bool) error {
	// Get sources.
	s, err := p.variableSource("source")
	if err != nil {
		return err
	}
	sl := strings.Fields(s)

	// Unpack sources.
	for _, s := range sl {
		n, err := p.variable("name")
		if err != nil {
			return err
		}
		wd := path.Join(config.WrkDir, n)
		os.Mkdir(wd, 0777)

		// Continue if file is not an URL.
		var cmd *exec.Cmd
		r := regexp.MustCompile(".(tar|tar.gz|tar.Z|tgz|tar.bz2|tbz2|tar.xz|txz|tar.lzma|tar.lz|zip|rpm)$")
		if r.MatchString(s) {
			cmd = exec.Command("bsdtar", "-p", "-o", "-C", wd, "-xf", path.Join(config.SrcDir, path.Base(s)))
		} else {
			cmd = exec.Command("cp", path.Join(p.Loc, path.Base(s)), wd)
		}
		if v {
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
		}

		printi("Unpacking " + path.Base(s))
		if err := cmd.Run(); err != nil {
			os.Remove(wd)
			return fmt.Errorf("pkg unpack %s: Something went wrong", path.Base(s))
		}
	}

	return nil
}

// update updates a package.
func (p pkgfile) update(v bool) error {
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
