package ports

import "path"

// A Port describes a port. A port is a directory containing the files needed
// for building a package.
type Port struct {
	// Location specifies the location of the port, this is used as the "primary
	// key" of a port type.
	Location Location

	// TODO: Add signature, .nostrip, et cetera.
	Footprint Footprint
	Md5sum    Md5sum
	Pkgfile   Pkgfile

	// Depends is a "recursive variable" that list all dependencies recursively.
	// This
	Depends []*Port
}

// New returns a Port with the Location field populated. Use the various
// `Parse*` functions to populate the other fields.
func New(location string) Port {
	var p Port

	p.Location = Location{
		Root: path.Dir(path.Dir(location)),
		Repo: path.Base(path.Dir(location)),
		Port: path.Base(location),
	}

	p.Footprint = Footprint{Port: &p}
	p.Md5sum = Md5sum{Port: &p}
	p.Pkgfile = Pkgfile{Port: &p}

	return p
}

// Alias aliases ports by using the `Aliases` variable. An example of this would
// be aliasing `core/openssl` to `6c37/libressl`.
func (p *Port) Alias() {
	for _, a := range Aliases {
		if a[0] == p.Location {
			p.Location = a[1]
		}
	}
}

// Check used by `ParseDepends()`.
var check []*Port

// ParseDepends is a function that calculates dependencies recursively and
// populates `Depends`.
//TODO: This requires `Pkgfile.Parse` has been run on the given Port.
func (p *Port) ParseDepends(ports []Port, alias bool) error {
	// Continue if already checked.
	for _, c := range check {
		if c.Pkgfile.Name == p.Pkgfile.Name {
			p.Depends = c.Depends
			return nil
		}
	}

	var err error
	for _, n := range p.Pkgfile.Depends {
		pl, err := Locate(ports, n)
		if err != nil {
			return err
		}
		d := pl[0]

		// Alias ports if needed.
		if alias {
			d.Alias()
		}

		// Read out Pkgfile.
		if err = d.Pkgfile.Parse(); err != nil {
			return err
		}

		// Append to `p.Depends`.
		p.Depends = append(p.Depends, &d)

		// Append port to checked ports.
		check = append(check, p)

		// Loop.
		if err = p.Depends[len(p.Depends)-1].ParseDepends(ports,
			alias); err != nil {
			return err
		}
	}

	return err
}
