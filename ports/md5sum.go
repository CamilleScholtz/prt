package ports

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strings"
)

// An Md5sum is a type describing the `.md5sum` file of a port. This file is
// used to validate the sources of a port.
type Md5sum struct {
	*Port

	Hashes []string
	Files  []string
}

// Parse parses the `.md5sum` file of a port and populates the various fields in
// the given `*Md5sum`.
func (f *Md5sum) Parse() error {
	r, err := os.Open(path.Join(f.Location.Full(), ".md5sum"))
	if err != nil {
		return fmt.Errorf("could not open `%s/.md5sum`", f.Location.Full())
	}
	defer r.Close()
	s := bufio.NewScanner(r)

	for s.Scan() {
		l := strings.Split(s.Text(), "  ")

		f.Hashes = append(f.Hashes, l[0])
		f.Files = append(f.Files, l[1])
	}

	return nil
}
