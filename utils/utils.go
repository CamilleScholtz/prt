package utils

// StringInList checks if a string is in a list
func StringInList(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}

	return false
}

// TrimString trims a string to a certain lenght
func TrimString(s string, n int) string {
	if len(s) <= n {
		return s
	}

	return s[0:n-1] + "…"
}
