package main

import (
	"fmt"
	"path"
	"regexp"

	"github.com/go2c/optparse"
	"github.com/hacdias/fileutils"
	"github.com/onodera-punpun/prt/ports"
	"github.com/onodera-punpun/prt/tar"
)

// unpackCommand unpacks port sources
// TODO: Use https://github.com/mholt/archiver.
func unpackCommand(input []string) error {
	// Define valid arguments.
	o := optparse.New()
	argh := o.Bool("help", 'h', false)

	// Parse arguments.
	_, err := o.Parse(input)
	if err != nil {
		return fmt.Errorf("invaild argument, use `-h` for a list of arguments")
	}

	// Print help.
	if *argh {
		fmt.Println("Usage: prt unpack [arguments]")
		fmt.Println("")
		fmt.Println("arguments:")
		fmt.Println("  -h,   --help            print help and exit")

		return nil
	}

	if err := downloadCommand([]string{}); err != nil {
		return err
	}

	p := ports.New(".")
	if err := p.Pkgfile.Parse(true); err != nil {
		return err
	}

	// Print a new line if files have been downloaded.
	// TODO: Can I simplify this?
	var sl = []string{}
	r := regexp.MustCompile("^(http|https|ftp|file)://")
	for _, s := range p.Pkgfile.Source {
		if r.MatchString(s) {
			sl = append(sl, s)
		}
	}
	if len(sl) > 0 {
		fmt.Println()
	}

	if err := p.CreateWrk(); err != nil {
		return err
	}

	to := path.Join(config.WrkDir, p.Pkgfile.Name, "src")
	for i, s := range p.Pkgfile.Source {
		f := path.Base(s)

		fmt.Printf("Unpacking source %d/%d, %s.\n", i+1, len(p.Pkgfile.Source),
			light(f))

		if tar.IsArchive(f) {
			if err := tar.Unpack(path.Join(config.SrcDir, f), to); err != nil {
				return err
			}
		} else {
			if err := fileutils.CopyFile(path.Join(p.Location.Full(), f),
				to); err != nil {
				return err
			}
		}
	}

	return nil
}
