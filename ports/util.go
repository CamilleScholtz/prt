package ports

// stringInStrings tests if a string is in a list.
func stringInStrings(s string, l []string) bool {
	for _, ls := range l {
		if ls == s {
			return true
		}
	}

	return false
}

// stringInPorts test if a string is in a list of Ports.
func stringInPorts(s string, l []Port) bool {
	for _, p := range l {
		if p.Location.Full() == s {
			return true
		}
	}

	return false
}
