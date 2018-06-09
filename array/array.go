package array

// ContainsString tests if a string is in a list of string.
func ContainsString(l []string, s string) bool {
	for _, i := range l {
		if i == s {
			return true
		}
	}

	return false
}
