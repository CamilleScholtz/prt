package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path"
	"strings"
)

// TODO: Make footprint its own type.
// TODO: Add signature.
type port struct {
	Loc string

	Footprint io.Reader
	Md5sum    md5sum
	Pkgfile   pkgfile
}

type md5sum struct {
	Hash []string
	File []string
}

type pkgfile struct {
	Description string
	URL         string
	Maintainer  string

	Depends  []string
	Optional []string

	Name    string
	Version string
	Release string

	Source []string
}

// decodeFootprint decodes a .footprint.
func decodeFootprint(l string) (io.Reader, error) {
	s, err := os.Open(path.Join(l, ".footprint"))
	if err != nil {
		return nil, err
	}
	// TODO: I should probably close this somewhere, right?
	// defer o.Close()

	return s, nil
}

// decodeMd5sum decodes a .md5sum.
func decodeMd5sum(l string) (md5sum, error) {
	var m md5sum

	mf, err := os.Open(path.Join(l, ".md5sum"))
	if err != nil {
		return m, err
	}
	defer mf.Close()
	s := bufio.NewScanner(mf)

	for s.Scan() {
		l := strings.Fields(s.Text())

		m.Hash = append(m.Hash, l[0])
		m.File = append(m.File, l[1])
	}

	return m, nil
}

// decodePkgfile decodes a Pkgfile.
func decodePkgfile(l string, strict bool) (pkgfile, error) {
	var f pkgfile

	pf, err := os.Open(path.Join(l, "Pkgfile"))
	if err != nil {
		return f, err
	}
	defer pf.Close()
	s := bufio.NewScanner(pf)

	for s.Scan() {
		i := s.Text()

		// TODO: Use FieldsFunc maybe.
		switch {
		case strings.HasPrefix(i, "# Description:"):
			f.Description = strings.TrimSpace(strings.TrimPrefix(i, "# Description:"))
		case strings.HasPrefix(i, "# URL:"):
			f.URL = strings.TrimSpace(strings.TrimPrefix(i, "# URL:"))
		case strings.HasPrefix(i, "# Maintainer:"):
			f.Maintainer = strings.TrimSpace(strings.TrimPrefix(i, "# Maintainer:"))

		case strings.HasPrefix(i, "# Depends on:"):
			f.Depends = strings.Fields(strings.Replace(strings.TrimSpace(strings.TrimPrefix(i, "# Depends on:")), ",", "", -1))
		case strings.HasPrefix(i, "# Optional:"), strings.Contains(i, "# Nice to have:"):
			f.Optional = strings.Fields(strings.TrimSpace(strings.TrimPrefix(i, "# Optional:")))
			if len(f.Optional) == 0 {
				f.Optional = strings.Fields(strings.TrimSpace(strings.TrimPrefix(i, "# Nice to have:")))
			}

		case strings.HasPrefix(i, "name="):
			f.Name = strings.TrimSpace(strings.TrimPrefix(i, "name="))
		case strings.HasPrefix(i, "version="):
			f.Version = strings.TrimSpace(strings.TrimPrefix(i, "version="))
		case strings.HasPrefix(i, "release="):
			f.Release = strings.TrimSpace(strings.TrimPrefix(i, "release="))

		case strings.HasPrefix(i, "source="):
			if strict {
				t, err := source(l, "source")
				if err != nil {
					return f, err
				}
				f.Source = strings.Fields(t)
			} else {
				f.Source = strings.Fields(strings.TrimSpace(strings.TrimPrefix(i, "source=")))
			}
			break
		}
	}

	return f, nil
}

// decodePort decodes a port.
func decodePort(l string, tl ...string) (port, error) {
	p := port{l, nil, md5sum{}, pkgfile{}}

	var err error
	for _, t := range tl {
		switch t {
		case "Footprint":
			p.Footprint, err = decodeFootprint(p.Loc)
		case "Md5sum":
			p.Md5sum, err = decodeMd5sum(p.Loc)
		case "Pkgfile":
			p.Pkgfile, err = decodePkgfile(p.Loc, false)
		default:
			return port{}, fmt.Errorf("port decodePort %s: No type '%s'", portBaseLoc(l), t)
		}
		if err != nil {
			return port{}, err
		}
	}

	return p, nil
}

// decodePortStrict decodes a port using source, which is slower.
func decodePortStrict(l string, tl ...string) (port, error) {
	p := port{l, nil, md5sum{}, pkgfile{}}

	var err error
	for _, t := range tl {
		switch t {
		case "Footprint":
			p.Footprint, err = decodeFootprint(p.Loc)
		case "Md5sum":
			p.Md5sum, err = decodeMd5sum(p.Loc)
		case "Pkgfile":
			p.Pkgfile, err = decodePkgfile(p.Loc, true)
		default:
			return port{}, fmt.Errorf("port decodePort %s: No type '%s'", portBaseLoc(l), t)
		}
		if err != nil {
			return port{}, err
		}
	}

	return p, nil
}

// source reads a variable from a Pkgfile, this actually uses bash source
// This is relatively slow but also more precise because it completes variables.
// This is especially (only?) seful for the source variable in Pkgfiles.
func source(l, k string) (string, error) {
	cmd := exec.Command("bash", "-c", "source ./Pkgfile && echo ${"+k+"[@]}")
	cmd.Dir = l
	var b bytes.Buffer
	cmd.Stdout = &b

	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("port source %s: Could not source", k)
	}

	if len(b.String()) == 0 {
		return "", fmt.Errorf("port source %s: No such variable", k)
	}

	return b.String(), nil
}
