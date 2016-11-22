package main

// This function checks if a string is in a list
func StringInList(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}

	return false
}
