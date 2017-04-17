package main

var check []string
var level int

// recursive is a function that calculates dependencies recursively.
func recursive(p port, dl map[string][]string, alias bool, all, inst []string) map[string][]string {
	// Continue if already checked.
	if stringInList(p.Location, check) {
		return dl
	}

	for _, n := range p.Pkgfile.Depends {
		ll, err := location(n, all)
		if err != nil {
			printe(err.Error())
			continue
		}

		d, err := parsePort(fullLocation(ll[0]), "Pkgfile")
		if err != nil {
			printe(err.Error())
			continue
		}

		// Alias ports if needed.
		if alias {
			d.alias()
		}

		// Finally print the port.
		dl[p.Location] = append(dl[baseLocation(p.Location)], baseLocation(d.Location))

		// Loop.
		level++
		recursive(d, dl, alias, all, inst)
		level--
	}

	// Append port to checked ports.
	check = append(check, p.Location)

	return dl
}
