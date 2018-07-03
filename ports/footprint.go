package ports

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strings"
)

// A Footprint describes the `.footprint` file of a port. This file is used for
// regression testing and contains a list of files a package is expected to
// contain once it is built.
// TODO: Handle symlinks (`->`).
type Footprint struct {
	*Port

	Files []struct {
		Path       string
		Permission Permission
	}
}

// Parse parses the `.footprint` file of a port and populates the various fields
// in the given `*Footprint`.
func (f *Footprint) Parse() error {
	r, err := os.Open(path.Join(f.Location.Full(), ".footprint"))
	if err != nil {
		return fmt.Errorf("could not open `%s/.footprint`", f.Location.Full())
	}
	defer r.Close()
	s := bufio.NewScanner(r)

	for s.Scan() {
		l := strings.Split(s.Text(), "\t")
		f.Files = append(f.Files, struct {
			Path       string
			Permission Permission
		}{
			Path: l[2],
			Permission: Permission{
				FileMode: toFileMode(l[0]),
				Owner:    l[1],
			},
		})
	}

	return nil
}
