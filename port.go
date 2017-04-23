// port.go contains functions that interact with single ports. These
// include functions such as getting the location of a port, reading
// out a .footprint file, or reading out a Pkgfile.

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"
)

// port is a struct type with all the files a port can have.
type port struct {
	// The location of the port, is often used as the "key" for the
	// port type.
	Location string

	// This is a recursive variable that gets filled by getDepends.
	Depends []port

	// TODO: Add signature, .nostrip, et cetera.
	Footprint footprint
	Md5sum    md5sum
	Pkgfile   pkgfile
}

// footprint is a struct type describing a .footprint file.
type footprint struct {
	Permission []string
	Owner      []string
	File       []string
}

// md5sum is a struct type describing a .md5sum file.
type md5sum struct {
	Hash []string
	File []string
}

// pkgfile is a struct type describing a parsed Pkgfile file.
type pkgfile struct {
	// Comments with various information that isn't strictly needed in
	// order to build a package.
	Description string
	URL         string
	Maintainer  string

	// Comments with information about dependencies. These need some
	// more parsing because there isn't an official styling guideline,
	// so some Pkgfiles use commas to separate dependencies, and some
	// don't.
	Depends  []string
	Optional []string

	// Variables with various information that is required in order to
	// build a package.
	Name    string
	Version string
	Release string

	// A variable array with the needed sources of a port. We probably
	// need to parse this by actually using bash because people often
	// use variables (such as $name or $version) and bashisms in the
	// source variable.
	Source []string
}

// Global variable needed by getDepends.
var check []string

// alias aliases ports using the config.Alias values.
func (p *port) alias() {
	for _, a := range config.Alias {
		if a[0] == p.getBaseDir() {
			p.Location = path.Join(config.PrtDir, a[1])
		}
	}
}

// getDepends is a function that calculates dependencies recursively.
func (p *port) getDepends(alias bool, all []port) {
	// Continue if already checked.
	if stringInList(p.Location, check) {
		return
	}

	for _, n := range p.Pkgfile.Depends {
		pl, err := location(n, all)
		if err != nil {
			continue
		}
		d := pl[0]

		// Alias ports if needed.
		if alias {
			d.alias()
		}

		// Read out Pkgfile.
		if err := d.parsePkgfile(); err != nil {
			printe(err.Error())
			continue
		}

		// Append to depends.
		p.Depends = append(p.Depends, d)

		// Append port to checked ports.
		check = append(check, p.Location)

		// Loop.
		p.Depends[len(p.Depends)-1].getDepends(alias, all)
	}
}

// getPortDir returns the port directory name. So `prtdir/repo/port`
// becomes `port`.
func (p port) getPortDir() string {
	return path.Base(p.Location)
}

// getRepoDir returns the port directory name. So `prtdir/repo/port`
// becomes `repo`.
func (p port) getRepoDir() string {
	return (path.Base(path.Dir(p.Location)))
}

// getBaseDir returns the port directory name. So `prtdir/repo/port`
// becomes `repo/port`.
// TODO: What if the location is `.` or `notprtdir/port`?
func (p port) getBaseDir() string {
	return strings.TrimPrefix(p.Location, config.PrtDir+"/")
}

// parseFootprint parses a .footprint file. It will read the
// .footprint file into a footprint type, which is a struct containing
// permissions and ownership information and their matching files.
func (p *port) parseFootprint() error {
	f, err := os.Open(path.Join(p.Location, ".footprint"))
	defer f.Close()
	if err != nil {
		return err
	}
	s := bufio.NewScanner(f)

	for s.Scan() {
		l := strings.Split(s.Text(), "\t")

		p.Footprint.Permission = append(p.Footprint.Permission, l[0])
		p.Footprint.Owner = append(p.Footprint.Owner, l[1])
		p.Footprint.File = append(p.Footprint.File, l[2])
	}

	return nil
}

// parseMd5sum parses a .md5sum file. It will read the .md5sum file
// into a md5sum type, which is a struct containing hashes and their
// matching files.
func (p *port) parseMd5sum() error {
	f, err := os.Open(path.Join(p.Location, ".md5sum"))
	defer f.Close()
	if err != nil {
		return err
	}
	s := bufio.NewScanner(f)

	for s.Scan() {
		l := strings.Split(s.Text(), "  ")

		p.Md5sum.Hash = append(p.Md5sum.Hash, l[0])
		p.Md5sum.File = append(p.Md5sum.File, l[1])
	}

	return nil
}

// parsePkgfile parses a Pkgfile file. It will read the Pkgfile file
// into a pkgfile type, which is a struct containing the various info
// a Pkgfile contains.
func (p *port) parsePkgfile() error {
	f, err := os.Open(path.Join(p.Location, "Pkgfile"))
	defer f.Close()
	if err != nil {
		return err
	}
	s := bufio.NewScanner(f)

	for s.Scan() {
		i := s.Text()

		if strings.HasPrefix(i, "#") {
			switch {
			case p.Pkgfile.Description == "" &&
				strings.HasPrefix(i, "# Description:"):
				p.Pkgfile.Description = strings.TrimSpace(
					strings.TrimPrefix(i, "# Description:"))
			case p.Pkgfile.URL == "" &&
				strings.HasPrefix(i, "# URL:"):
				p.Pkgfile.URL = strings.TrimSpace(
					strings.TrimPrefix(i, "# URL:"))
			case p.Pkgfile.Maintainer == "" &&
				strings.HasPrefix(i, "# Maintainer:"):
				p.Pkgfile.Maintainer = strings.TrimSpace(
					strings.TrimPrefix(i, "# Maintainer:"))

			case len(p.Pkgfile.Depends) == 0 &&
				strings.HasPrefix(i, "# Depends on:"):
				p.Pkgfile.Depends = strings.Fields(strings.Replace(
					strings.TrimSpace(strings.TrimPrefix(
						i, "# Depends on:")), ",", "", -1))
			case len(p.Pkgfile.Optional) == 0 &&
				strings.HasPrefix(i, "# Optional:"):
				p.Pkgfile.Optional = strings.Fields(strings.Replace(
					strings.TrimSpace(strings.TrimPrefix(
						i, "# Optional:")), ",", "", -1))
			case len(p.Pkgfile.Optional) == 0 &&
				strings.Contains(i, "# Nice to have:"):
				p.Pkgfile.Optional = strings.Fields(strings.Replace(
					strings.TrimSpace(strings.TrimPrefix(
						i, "# Nice to have:")), ",", "", -1))
			}
		} else {
			switch {
			case p.Pkgfile.Name == "" &&
				strings.HasPrefix(i, "name="):
				p.Pkgfile.Name = strings.TrimSpace(
					strings.TrimPrefix(i, "name="))
			case p.Pkgfile.Version == "" &&
				strings.HasPrefix(i, "version="):
				p.Pkgfile.Version = strings.TrimSpace(
					strings.TrimPrefix(i, "version="))
			case p.Pkgfile.Release == "" &&
				strings.HasPrefix(i, "release="):
				p.Pkgfile.Release = strings.TrimSpace(
					strings.TrimPrefix(i, "release="))

			case len(p.Pkgfile.Source) == 0 &&
				strings.HasPrefix(i, "source="):
				p.Pkgfile.Source = strings.Fields(
					strings.TrimSpace(
						strings.TrimPrefix(i, "source=")))
				break
			}
		}
	}

	return nil
}

// parsePkgfileStrict parses a Pkgfile file. It will read the Pkgfile
// file into a pkgfile type, which is a struct containing the various
// info a Pkgfile contains. parsePkgfileStrict differs from
// parsePkgfile in that source is used to parse the source field.
// Since this forks to bash this is relatively slow.
func (p *port) parsePkgfileStrict() error {
	f, err := os.Open(path.Join(p.Location, "Pkgfile"))
	defer f.Close()
	if err != nil {
		return err
	}
	s := bufio.NewScanner(f)

	for s.Scan() {
		i := s.Text()

		if strings.HasPrefix(i, "#") {
			switch {
			case p.Pkgfile.Description == "" &&
				strings.HasPrefix(i, "# Description:"):
				p.Pkgfile.Description = strings.TrimSpace(
					strings.TrimPrefix(i, "# Description:"))
			case p.Pkgfile.URL == "" &&
				strings.HasPrefix(i, "# URL:"):
				p.Pkgfile.URL = strings.TrimSpace(
					strings.TrimPrefix(i, "# URL:"))
			case p.Pkgfile.Maintainer == "" &&
				strings.HasPrefix(i, "# Maintainer:"):
				p.Pkgfile.Maintainer = strings.TrimSpace(
					strings.TrimPrefix(i, "# Maintainer:"))

			case len(p.Pkgfile.Depends) == 0 &&
				strings.HasPrefix(i, "# Depends on:"):
				p.Pkgfile.Depends = strings.Fields(strings.Replace(
					strings.TrimSpace(strings.TrimPrefix(
						i, "# Depends on:")), ",", "", -1))
			case len(p.Pkgfile.Optional) == 0 &&
				strings.HasPrefix(i, "# Optional:"):
				p.Pkgfile.Optional = strings.Fields(strings.Replace(
					strings.TrimSpace(strings.TrimPrefix(
						i, "# Optional:")), ",", "", -1))
			case len(p.Pkgfile.Optional) == 0 &&
				strings.Contains(i, "# Nice to have:"):
				p.Pkgfile.Optional = strings.Fields(strings.Replace(
					strings.TrimSpace(strings.TrimPrefix(
						i, "# Nice to have:")), ",", "", -1))
			}
		} else {
			switch {
			case p.Pkgfile.Name == "" &&
				strings.HasPrefix(i, "name="):
				p.Pkgfile.Name = strings.TrimSpace(
					strings.TrimPrefix(i, "name="))
			case p.Pkgfile.Version == "" &&
				strings.HasPrefix(i, "version="):
				p.Pkgfile.Version = strings.TrimSpace(
					strings.TrimPrefix(i, "version="))
			case p.Pkgfile.Release == "" &&
				strings.HasPrefix(i, "release="):
				p.Pkgfile.Release = strings.TrimSpace(
					strings.TrimPrefix(i, "release="))

			case len(p.Pkgfile.Source) == 0 &&
				strings.HasPrefix(i, "source="):
				t, err := p.source("source")
				if err != nil {
					return err
				}
				p.Pkgfile.Source = strings.Fields(t)
				break
			}
		}
	}

	return nil
}

// location tries to get the location of a port. It returns a list
// with possible ports, ordered using the config Order value.
func location(n string, all []port) ([]port, error) {
	var pl []port
	for _, p := range all {
		if path.Base(p.Location) == n {
			pl = append(pl, p)
		}
	}

	if len(pl) == 0 {
		return []port{}, fmt.Errorf(
			"location %s: Not in the ports tree", n)
	}

	// If there are multiple matches, sort using the config Order
	// value.
	if len(pl) > 1 {
		var i int
		for _, r := range config.Order {
			npl := path.Join(config.PrtDir, r, pl[i].getPortDir())
			if stringInPorts(npl, all) {
				pl[i].Location = npl
				i++
			}

			// Break if everything has been ordered.
			if i == len(pl) {
				break
			}
		}
	}

	return pl, nil
}

// source reads a variable from a Pkgfile, this actually uses bash
// source. This is relatively slow but also more precise because it
// completes variables. This is especially (only?) useful for the
// source variable in Pkgfiles.
func (p port) source(k string) (string, error) {
	cmd := exec.Command("bash", "-c",
		"source ./Pkgfile && echo ${"+k+"[@]}")
	cmd.Dir = p.Location
	var b bytes.Buffer
	cmd.Stdout = &b

	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("source %s: Could not source", k)
	}

	if len(b.String()) == 0 {
		return "", fmt.Errorf("source %s: No such variable", k)
	}

	return b.String(), nil
}
