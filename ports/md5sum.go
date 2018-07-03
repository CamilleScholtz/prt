package ports

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"path"
	"regexp"
	"strings"
)

// An Md5sum is a type describing the `.md5sum` file of a port. This file is
// used to validate the sources of a port.
type Md5sum struct {
	*Port

	Files []struct {
		File string
		Hash string
	}
}

// Create creates an `.md5sum` file for a Port.
func (f *Md5sum) Create() error {
	if err := f.Pkgfile.Parse(true); err != nil {
		return err
	}

	nf, err := os.OpenFile(path.Join(f.Location.Full(), ".md5sum"), os.O_CREATE|
		os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer nf.Close()

	r := regexp.MustCompile("^(http|https|ftp|file)://")
	for _, s := range f.Pkgfile.Source {
		if r.MatchString(s) {
			s = path.Join(SrcDir, path.Base(s))
		} else {
			s = path.Join(f.Location.Full(), path.Base(s))
		}

		h, err := hashFromFile(s)
		if err != nil {
			return err
		}

		if _, err := nf.WriteString(h + "  " + path.Base(s) +
			"\n"); err != nil {
			return err
		}
	}

	return nil
}

// Parse parses the `.md5sum` file of a port and populates the various fields in
// the given `*Md5sum`.
func (f *Md5sum) Parse() error {
	r, err := os.Open(path.Join(f.Location.Full(), ".md5sum"))
	if err != nil {
		return fmt.Errorf("could not open `%s/.md5sum`", f.Location.Full())
	}
	defer r.Close()
	s := bufio.NewScanner(r)

	for s.Scan() {
		l := strings.Split(s.Text(), "  ")
		f.Files = append(f.Files, struct {
			File string
			Hash string
		}{
			File: l[1],
			Hash: l[0],
		})
	}

	return nil
}

func (f *Md5sum) Validate() error {
	if err := f.Pkgfile.Parse(true); err != nil {
		return err
	}

	if err := f.Parse(); err != nil {
		return err
	}

	r := regexp.MustCompile("^(http|https|ftp|file)://")
	for _, l := range f.Md5sum.Files {
		for _, s := range f.Pkgfile.Source {
			if r.MatchString(s) {
				s = path.Join(SrcDir, path.Base(s))
			} else {
				s = path.Join(f.Location.Full(), path.Base(s))
			}

			h, err := hashFromFile(s)
			if err != nil {
				return err
			}

			if l.Hash != h {
				return fmt.Errorf("pkgmk md5sum %s:%s: Hash didn't match", f.
					Location.Base(), path.Base(s))
			}
		}
	}

	return nil
}

func hashFromFile(file string) (string, error) {
	hf, err := os.Open(file)
	if err != nil {
		return "", err
	}
	defer hf.Close()

	h := md5.New()
	if _, err := io.Copy(h, hf); err != nil {
		return "", err
	}

	return hex.EncodeToString(h.Sum(nil)), nil
}
