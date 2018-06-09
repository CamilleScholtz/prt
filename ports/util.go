package ports

// Contains tests if a string is in a list of Ports.
// TODO: Possibly use Location instead of string here.
func Contains(l []Port, s string) bool {
	for _, p := range l {
		if p.Location.Full() == s {
			return true
		}
	}

	return false
}
