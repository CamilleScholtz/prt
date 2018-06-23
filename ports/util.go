package ports

// Contains tests if a location is in a list of Ports.
func Contains(l []Port, s Location) bool {
	for _, p := range l {
		if p.Location == s {
			return true
		}
	}

	return false
}
