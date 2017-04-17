package main

import (
	"path"
	"path/filepath"
	"strings"
)

// ports lists all ports found in the PrtDir.
func ports() ([]string, error) {
	// TODO: Is there something more efficient than Glob?
	dl, err := filepath.Glob(path.Join(config.PrtDir, "/*/*/Pkgfile"))
	if err != nil {
		return []string{}, err
	}

	// Remove PrtDir from output.
	var p []string
	for _, d := range dl {
		p = append(p, baseLocation(path.Dir(d)))
	}

	return p, nil
}

// baseLocation removes PrtDir from a string.
// TODO: Make it path.Base if config.PrtDir is not found.
func baseLocation(n string) string {
	return strings.TrimPrefix(n, config.PrtDir+"/")
}

// fullLocation adds the PrtDir to a string.
// TODO: Make it path.Base if config.PrtDir is not found.
func fullLocation(n string) string {
	return path.Join(config.PrtDir, n)
}
