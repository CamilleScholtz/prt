package ports

var (
	// Prtdir defines The root directory of the ports tree.
	PrtDir = "/usr/ports"

	// PkgDir defines the directory where created packages get stored.
	PkgDir = "."

	// SrcDir defines the directory where downloaded sources get stored.
	SrcDir = "."

	// WrkDir defines the root directory where packages get build..
	WrkDir = "."

	// Order gets used to determine how to order multiple ports with the same
	// name, but residing in a different repository. The repository found first
	// in the Order variable being ordered firt.
	Order []string

	// Alias gets used used to alias one port location to another port location.
	Aliases [][]Location
)
