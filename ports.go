// ports.go contains functions that interact with multiple ports.
// These include functions such as getting a list of all ports in the
// prtdir, prefixing the prtdir to string, or removing the prtdir from
// a string.

package main

import (
	"path"
	"path/filepath"
)

// ports lists all ports found in the prtdir.
func ports() ([]port, error) {
	pl, err := filepath.Glob(path.Join(config.PrtDir, "/*/*/Pkgfile"))
	if err != nil {
		return []port{}, err
	}

	var ptl []port
	for _, p := range pl {
		var pt port
		pt.Location = path.Dir(p)
		ptl = append(ptl, pt)
	}

	return ptl, nil
}
