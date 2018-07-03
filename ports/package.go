package ports

// Package is a type describing a package.
type Package struct {
	Name    string
	Version string
	Release string

	Files []string
}
