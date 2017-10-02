// pkgmk.go contains functions related to building and installing
// ports/packages. These include function such downloading port
// sourcers, creating the wrkdir, removing the wrkdir and installing
// package contents to right location.

package main

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path"
	"regexp"
	"sort"
	"strings"

	"github.com/mholt/archiver"
)

// build builds a port. It does this by running a custom fork
// of pkgmk.
func (p port) build(v bool) error {
	var cmd *exec.Cmd
	cmd = exec.Command("/usr/share/prt/pkgmk")
	cmd.Dir = p.Location
	if v {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}

	// TODO: Make this behave like check/createmd5sum in regards to
	// updating or not.
	printi("Building package")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf(
			"build %s: Something went wrong", p.getBaseDir())
	}

	return nil
}

// check checks if all directories need to build a port are present.
func checkDir() error {
	if _, err := os.Stat(config.SrcDir); os.IsNotExist(err) {
		return err
	}
	if _, err := os.Stat(config.PkgDir); os.IsNotExist(err) {
		return err
	}
	if _, err := os.Stat(config.WrkDir); os.IsNotExist(err) {
		return err
	}

	return nil
}

// checkMd5sum validates the .md5sum file.
func (p port) checkMd5sum() error {
	p.createMd5sum("/tmp/prt")

	var t port
	t.Location = path.Join("/tmp/prt", p.Pkgfile.Name)
	if err := t.parseMd5sum(); err != nil {
		return err
	}

	var e bool
	for pi, pl := range p.Md5sum.Hash {
		for ti, tl := range t.Md5sum.Hash {
			if pl == tl {
				if len(tl) == 0 {
					e = true
					printe("1 Mismatch " + pl)
				} else if pl != tl {
					e = true
					printe("2 Mismatch " + tl)
				}
			}

			if ti <= pi {
				break
			}
		}
	}

	if e {
		return fmt.Errorf("pkgmk md5sum %s: verification failed",
			p.getBaseDir())
	}
	return nil
}

// check check if all needed variables are present.
func (p port) checkPkgfile() error {
	if p.Pkgfile.Name == "" {
		return fmt.Errorf("pkgfile checkPkgfile %s: Name variable is empty",
			p.getBaseDir())
	}
	if p.Pkgfile.Version == "" {
		return fmt.Errorf("pkgfile checkPkgfile %s: Version variable is empty",
			p.getBaseDir())
	}
	if p.Pkgfile.Release == "" {
		return fmt.Errorf("pkgfile checkPkgfile %s: Release variable is empty",
			p.getBaseDir())
	}
	// TODO: Add a function function in port.go.
	//if err := p.function("build"); err != nil {
	//	return fmt.Errorf(
	//	"pkgfile checkPkgfile %s: Build function is empty",
	//	p.getBaseDir())
	//}

	return nil
}

// checkSignature checks the .signature file.
// TODO: Rewrite this.
func (p port) checkSignature() error {
	sl := p.Pkgfile.Source
	sort.Sort(byBase(sl))

	// Prepend Pkgfile and .footprint to sources.
	sl = append([]string{"Pkgfile", ".footprint"}, sl...)

	for _, s := range sl {
		r := regexp.MustCompile("^(http|https|ftp|file)://")
		if r.MatchString(s) {
			s = path.Join(config.SrcDir, path.Base(s))
		} else {
			s = path.Join(p.Location, path.Base(s))
		}

		if err := os.Symlink(s, path.Join("/tmp/prt/"+path.Base(s))); err !=
			nil {
			return err
		}
	}

	// TODO: Do this in Go.
	cmd := exec.Command("signify", "-q", "-C", "-x", path.Join(p.Location,
		".signature"))
	cmd.Dir = "/tmp/prt"
	var b bytes.Buffer
	cmd.Stderr = &b

	if err := cmd.Run(); err != nil {
		for _, l := range strings.Split(b.String(), "\n") {
			if len(l) == 0 {
				continue
			}

			printe("Mismatch " + strings.Trim(l, ": FAIL"))
		}
		return fmt.Errorf("pkgmk signature %s: verification failed",
			p.getBaseDir())
	}

	return nil
}

// cleanWrk removes the necessary WrkDir directories.
func (p port) cleanWrk() error {
	if err := os.RemoveAll(path.Join(config.WrkDir,
		p.Pkgfile.Name)); err != nil {
		return err
	}

	// Temp directory used by some functions.
	if err := os.RemoveAll("/tmp/prt"); err != nil {
		return err
	}

	return nil
}

// createMd5sum creates a .md5sum file.
func (p port) createMd5sum(l string) error {
	sl := p.Pkgfile.Source
	sort.Sort(byBase(sl))

	f, err := os.OpenFile(path.Join(l, ".md5sum"),
		os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	defer f.Close()
	if err != nil {
		return err
	}

	// Just leave the file empty if there are no sources.
	if len(sl) == 0 {
		return nil
	}

	for _, s := range sl {
		r := regexp.MustCompile("^(http|https|ftp|file)://")
		if r.MatchString(s) {
			s = path.Join(config.SrcDir, path.Base(s))
		} else {
			s = path.Join(p.Location, path.Base(s))
		}

		hf, err := os.Open(s)
		defer hf.Close()
		if err != nil {
			return err
		}

		h := md5.New()
		if _, err := io.Copy(h, hf); err != nil {
			return err
		}

		if _, err := f.WriteString(hex.EncodeToString(h.Sum(nil)) + "  " +
			path.Base(s) + "\n"); err != nil {
			return err
		}
	}

	return nil
}

// createWrk creates the necessary WrkDir directories.
func (p port) createWrk() error {
	if err := os.Mkdir(path.Join(config.WrkDir, p.Pkgfile.Name), 0777); err !=
		nil {
		return err
	}
	if err := os.Mkdir(path.Join(config.WrkDir, p.Pkgfile.Name, "pkg"), 0777); cerr != nil {
		return err
	}
	if err := os.Mkdir(path.Join(config.WrkDir, p.Pkgfile.Name, "src"), 0777); err != nil {
		return err
	}

	// Temp directory used by some functions.
	// TODO: Is this needed?
	if err := os.Mkdir("/tmp/prt", 0777); err != nil {
		return err
	}

	return nil
}

// download downloads a port sources.
func (p port) download(v bool) error {
	sl := p.Pkgfile.Source
	sort.Sort(byBase(sl))

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

		// TODO: Can I use some Go package for this?
		cmd := exec.Command("curl", "-L", "-#", "--fail",
			"--ftp-pasv", "-C", "-", "-o", f+".partial", s)
		if v {
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
		}

		printi("Downloading " + path.Base(s))
		if err := cmd.Run(); err != nil {
			return fmt.Errorf(
				"pkgmk download %s: Could not download source",
				path.Base(s))
		}

		// Remove .partial on completion.
		os.Rename(f+".partial", f)
	}

	return nil
}

// install installs a package.
func (p port) install(v bool) error {
	cmd := exec.Command("/usr/share/prt/pkgmk", "-io")
	cmd.Dir = p.Location
	if v {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}

	if err := cmd.Run(); err != nil {
		return fmt.Errorf(
			"install %s: Something went wrong", p.getBaseDir())
	}

	return nil
}

// md5sum checks the .md5sum file.
func (p port) md5sum() error {
	// Only check md5sum if there is no signature.
	// TODO
	//if p.Signature != nil {
	//	return nil
	//}

	// Check .md5sum if it exists, else create it.
	if _, err := os.Stat(path.Join(p.Location,
		".md5sum")); err == nil {
		printi("Checking md5sum")
		if err := p.checkMd5sum(); err != nil {
			return err
		}
	} else {
		printi("Creating md5sum")
		if err := p.createMd5sum(p.Location); err != nil {
			return err
		}
	}

	return nil
}

// post runs a pre-install scripts.
func (p port) post(v bool) error {
	if _, err := os.Stat(path.Join(p.Location,
		"post-install")); err != nil {
		return nil
	}

	cmd := exec.Command("bash", "./post-install")
	cmd.Dir = p.Location
	if v {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}

	printi("Running post-install")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf(
			"pkgmk post %s: Something went wrong", p.getBaseDir())
	}

	return nil
}

// pre runs a pre-install scripts.
func (p port) pre(v bool) error {
	if _, err := os.Stat(path.Join(p.Location, "pre-install")); err != nil {
		return nil
	}

	cmd := exec.Command("bash", "./pre-install")
	cmd.Dir = p.Location
	if v {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}

	printi("Running pre-install")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf(
			"pkgmk pre %s: Something went wrong", p.getBaseDir())
	}

	return nil
}

// uninstall uninstalls a package.
// TODO: Rewrite this.
func pkgUninstall(todo string) error {
	cmd := exec.Command("pkgrm", todo)

	if err := cmd.Run(); err != nil {
		return fmt.Errorf(
			"pkgmk uninstall %s: Something went wrong", todo)
	}

	return nil
}

// unpack unpacks a port sources.
// TODO: Don't run this if the Pkgfile has its own unpack function.
// TODO: Or should I, it's a lot of added complexity for basically just
// the Go port, I could try to rewrite that port.
func (p port) unpack() error {
	sl := p.Pkgfile.Source
	sort.Sort(byBase(sl))

	// Unpack sources.
	for _, s := range sl {
		printi("Unpacking " + path.Base(s))

		for _, ff := range archiver.SupportedFormats {
			if !ff.Match(path.Base(s)) {
				continue
			}

			if err := ff.Open(path.Join(config.SrcDir, path.Base(s)), path.Join(config.WrkDir, path.Base(p.Location), "src")); err != nil {
				return err
			}
			continue
		}

		// TODO: Make this missing.
		f, _ := os.Open(path.Join(p.Location, path.Base(s)))
		defer f.Close()

		d, err := os.Create(path.Join(path.Join(config.WrkDir, path.Base(p.Location), "src"), path.Base(s)))
		if err != nil {
			return err
		}

		io.Copy(d, f)
		d.Close()
	}

	return nil
}

// update updates a package.
func (p port) update(v bool) error {
	cmd := exec.Command("/usr/share/prt/pkgmk", "-uo")
	cmd.Dir = p.Location
	if v {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}

	if err := cmd.Run(); err != nil {
		return fmt.Errorf(
			"update %s: Something went wrong", p.getBaseDir())
	}

	return nil
}

// pkgmk is a wrapper for all the functions in pkgmk.go.
func (p port) pkgmk(inst []string, v bool) error {
	if err := checkDir(); err != nil {
		return err
	}
	if err := p.checkPkgfile(); err != nil {
		return err
	}
	if err := p.createWrk(); err != nil {
		return err
	}
	defer p.cleanWrk()
	if err := p.pre(v); err != nil {
		return err
	}
	if err := p.download(v); err != nil {
		return err
	}
	if err := p.md5sum(); err != nil {
		return err
	}
	/*
		if err := p.unpack(); err != nil {
			return err
		}
		if !stringInList(path.Base(p.Location), inst) {
			if err := p.build(v); err != nil {
				return err
			}
		}
		if stringInList(path.Base(p.Location), inst) {
			printi("Updating package")
			if err := p.update(v); err != nil {
				return err
			}
		} else {
			printi("Installing package")
			if err := p.install(v); err != nil {
				return err
			}
		}
		if err := p.post(v); err != nil {
			return err
		}
	*/

	return nil
}
