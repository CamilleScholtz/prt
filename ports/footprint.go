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

	Permission []string
	Owner      []string
	File       []string
}

// Parse parses the `.footprint` file of a port and populates the various fields
// in the given `*Footprint`.
func (f *Footprint) Parse() error {
	fr, err := os.Open(path.Join(f.Location.Full(), ".footprint"))
	defer fr.Close()
	if err != nil {
		return fmt.Errorf("could not open `%s/.footprint`", f.Location.Full())
	}
	s := bufio.NewScanner(fr)

	for s.Scan() {
		l := strings.Split(s.Text(), "\t")

		f.Permission = append(f.Permission, l[0])
		f.Owner = append(f.Owner, l[1])
		f.File = append(f.File, l[2])
	}

	return nil
}
