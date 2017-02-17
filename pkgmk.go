package main

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

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

// build builds a port.
func (p port) build(f, v bool) error {
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

// checkMd5sum creates the .md5sum file.
func (p port) checkMd5sum() error {
	p.createMd5sum("/tmp")
	defer os.Remove("/tmp/.md5sum")

	o, err := os.Open(path.Join(p.Loc, ".md5sum"))
	if err != nil {
		return err
	}
	defer o.Close()
	n, err := os.Open("/tmp/.md5sum")
	if err != nil {
		return err
	}
	defer n.Close()

	var e bool
	so := bufio.NewScanner(o)
	var io int
	for so.Scan() {
		sn := bufio.NewScanner(n)
		var in int
		for i := 0; i <= io; i++ {
			sn.Scan()
			if io == in {
				lo := strings.Split(so.Text(), " ")
				ln := strings.Split(sn.Text(), " ")

				if len(sn.Text()) == 0 {
					e = true
					printe("Missing md5sum " + lo[0] + " for " + lo[2])
				} else if lo[0] != ln[0] {
					e = true
					printe("New md5sum " + ln[0] + " for " + ln[2])
				}
			}

			in++
		}

		if _, err := n.Seek(0, os.SEEK_SET); err != nil {
			return err
		}
		io++
	}

	if e {
		return fmt.Errorf("")
	}
	return nil
}

// createMd5sum creates the .md5sum file.
func (p port) createMd5sum(l string) error {
	// Get sources.
	s, err := p.variableSource("source")
	if err != nil {
		return err
	}
	sl := strings.Fields(s)
	sort.Strings(sl)

	// Create .md5sum file.
	for _, s := range sl {
		var f string
		r := regexp.MustCompile("^(http|https|ftp|file)://")
		if r.MatchString(s) {
			f = path.Join(config.SrcDir, path.Base(s))
		} else {
			f = path.Join(p.Loc, path.Base(s))
		}

		hf, err := os.Open(f)
		if err != nil {
			return err
		}
		defer hf.Close()

		h := md5.New()
		if _, err := io.Copy(h, hf); err != nil {
			return err
		}

		m, err := os.OpenFile(path.Join(l, ".md5sum"), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
		if err != nil {
			return err
		}
		defer m.Close()

		if _, err := m.WriteString(hex.EncodeToString(h.Sum(nil)) + "  " + path.Base(s) + "\n"); err != nil {
			return err
		}
	}

	return nil
}

// createWrkDir creates the necessary WrkDir directories.
func (p port) createWrkDir() (string, string, string, error) {
	// Get port name and create the workdir.
	n, err := p.variable("name")
	if err != nil {
		return "", "", "", err
	}

	wd := path.Join(config.WrkDir, n)
	wpd := path.Join(config.WrkDir, n, "pkg")
	wsd := path.Join(config.WrkDir, n, "src")
	os.Mkdir(wd, 0777)
	os.Mkdir(wpd, 0777)
	os.Mkdir(wsd, 0777)

	return wd, wpd, wsd, nil
}

// download downloads a port sources.
func (p port) download(v bool) error {
	// Get sources.
	s, err := p.variableSource("source")
	if err != nil {
		return err
	}
	sl := strings.Fields(s)
	sort.Strings(sl)

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
			return fmt.Errorf("pkgmk download %s: Something went wrong", path.Base(s))
		}

		// Remove .partial on completion.
		os.Rename(f+".partial", f)
	}

	return nil
}

// install installs a package.
func (p port) install(v bool) error {
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

// md5sum checks and optionally creates the .md5sum file.
func (p port) md5sum(v bool) error {
	if _, err := os.Stat(path.Join(p.Loc, ".md5sum")); err == nil {
		printi("Checking md5sum")
		if err := p.checkMd5sum(); err != nil {
			return err
		}
	} else {
		printi("Creating md5sum")
		if err := p.createMd5sum(p.Loc); err != nil {
			return err
		}
	}

	return nil
}

// post runs a pre-install scripts.
func (p port) post(v bool) error {
	cmd := exec.Command("bash", "./post-install")
	cmd.Dir = p.Loc
	if v {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}

	printi("Running post-install")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("pkgmk post %s: Something went wrong", portBaseLoc(p.Loc))
	}

	return nil
}

// pre runs a pre-install scripts.
func (p port) pre(v bool) error {
	cmd := exec.Command("bash", "./pre-install")
	cmd.Dir = p.Loc
	if v {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}

	printi("Running pre-install")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("pkgmk pre %s: Something went wrong", portBaseLoc(p.Loc))
	}

	return nil
}

// uninstall uninstalls a package.
// TODO
func pkgUninstall(todo string) error {
	cmd := exec.Command("pkgrm", todo)

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("pkgmk uninstall %s: Something went wrong", todo)
	}

	return nil
}

// unpack unpacks a port sources.
func (p port) unpack(v bool) error {
	// Get sources.
	s, err := p.variableSource("source")
	if err != nil {
		return err
	}
	sl := strings.Fields(s)
	sort.Strings(sl)

	_, _, wsd, err := p.createWrkDir()
	if err != nil {
		return err
	}

	// Unpack sources.
	for _, s := range sl {
		r := regexp.MustCompile(".(tar|tar.gz|tar.Z|tgz|tar.bz2|tbz2|tar.xz|txz|tar.lzma|tar.lz|zip|rpm)$")
		if r.MatchString(s) {
			printi("Unpacking " + path.Base(s))

			cmd := exec.Command("bsdtar", "-p", "-o", "-C", wsd, "-xf", path.Join(config.SrcDir, path.Base(s)))
			if v {
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
			}

			if err := cmd.Run(); err != nil {
				// TODO: Use clean here
				os.Remove(wsd)
				return fmt.Errorf("pkgmk unpack %s: Something went wrong", path.Base(s))
			}
		} else {
			printi("Moving " + path.Base(s))

			// TODO: MAke this missing
			f, _ := os.Open(path.Join(p.Loc, path.Base(s)))
			defer f.Close()

			d, err := os.Create(path.Join(wsd, path.Base(s)))
			if err != nil {
				return err
			}

			io.Copy(d, f)
			d.Close()
		}
	}

	return nil
}

// update updates a package.
func (p port) update(v bool) error {
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
