package ports

import "path"

// A Location describes the location of a port. A port location consist of a
// ports-tree directory, a repo and a port. An example of a valid location would
// be `usr/ports/opt/firefox`.
type Location struct {
	Root string
	Repo string
	Port string
}

// Base returns the repo and port name. An example of this would be
// `opt/firefox`.
func (l Location) Base() string {
	return path.Join(l.Repo, l.Port)
}

// Full returns the ports-tree directory, the repo and port name. An example of
// this would be `/usr/ports/opt/firefox`.
func (l Location) Full() string {
	return path.Join(l.Root, l.Repo, l.Port)
}
