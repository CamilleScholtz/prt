package ports

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
)

// All lists all ports found in the PrtDir.
func All() ([]Port, error) {
	var pl []Port

	err := filepath.Walk(PrtDir, func(p string, i os.FileInfo,
		err error) error {
		if err != nil {
			return err
		}

		if !i.IsDir() && i.Name() == "Pkgfile" {
			pl = append(pl, New(path.Dir(p)))
		}

		return nil
	})
	if err != nil {
		return pl, err
	}

	return pl, nil
}

// Locate tries to locate a port using a given list of Ports. It returns a list
// with possible ports, sorted according to the order parameter.
func Locate(ports []Port, port string) ([]Port, error) {
	// Find matching port names in the `all` list.
	var pl []Port
	for _, p := range ports {
		if p.Location.Port == port {
			pl = append(pl, p)
		}
	}

	// If there have been zero matches return with an error.
	if len(pl) == 0 {
		return []Port{}, fmt.Errorf("could not find `%s` in the ports tree",
			port)
	}

	// If there are multiple matches, sort according to the `Order` variable.
	if len(pl) > 1 {
		var i int
		for _, r := range Order {
			for j, p := range pl {
				if p.Location.Repo == r {
					pl[i], pl[j] = pl[j], pl[i]
					i++
				}
			}
		}

		if i != len(pl) {
			return []Port{}, fmt.Errorf("could not order `%s`", port)
		}
	}

	return pl, nil
}
