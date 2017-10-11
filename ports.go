// ports.go contains functions that interact with multiple ports. These include
// functions such as getting a list of all ports in the prtdir.

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
		// TODO: Repair
		pt.Location.Repo = path.Dir(p)
		pt.Location.Port = path.Dir(p)
		ptl = append(ptl, pt)
	}

	return ptl, nil
}
