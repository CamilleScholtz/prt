package ports

// Contains tests if a location is in a list of Ports.
func Contains(ports []Port, location Location) bool {
	for _, p := range ports {
		if p.Location == location {
			return true
		}
	}

	return false
}
