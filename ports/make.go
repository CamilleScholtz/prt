package ports

import (
	"os"
	"path"
)

// TODO: Move Download function from prt.
// TODO: Move Unpack function from prt.

// CreateWrk creates the necessary WrkDir directories.
func (p Port) CreateWrk() error {
	if err := os.MkdirAll(path.Join(WrkDir, p.Pkgfile.Name, "pkg"),
		0777); err != nil {
		return err
	}
	if err := os.MkdirAll(path.Join(WrkDir, p.Pkgfile.Name, "src"),
		0777); err != nil {
		return err
	}

	return nil
}

// CleanWrk removes the necessary WrkDir directories.
/*func (p Port) CleanWrk() error {
	if err := os.RemoveAll(path.Join(config.WrkDir,
		p.Pkgfile.Name)); err != nil {
		return err
	}

	// Temp directory used by some functions.
	if err := os.RemoveAll("/tmp/prt"); err != nil {
		return err
	}

	return nil
}*/

/*
// checkSignature checks the .signature file.
// TODO: Rewrite this.
func (p Port) checkSignature() error {
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

// install installs a package.
func (p port) install(v bool) error {
	cmd := exec.Command("/usr/share/prt/pkgmk", "-io")
	cmd.Dir = p.Location
	if v {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("install %s: Something went wrong", p.getBaseDir())
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
/*
	return nil
}
*/
