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
type Footprint struct {
	*Port

	Permissions []string
	Owners      []string
	Files       []string
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

		f.Permissions = append(f.Permissions, l[0])
		f.Owners = append(f.Owners, l[1])
		f.Files = append(f.Files, l[2])
	}

	return nil
}
