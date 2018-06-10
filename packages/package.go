package packages

// Package is a type describing an package.
type Package struct {
	Name    string
	Version string
	Release string

	Files []string
}
