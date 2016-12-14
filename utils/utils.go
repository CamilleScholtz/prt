package utils

import "bytes"

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

	var b bytes.Buffer
	for i := 0; i <= len(s); i++ {
		if i < n {
			b.WriteString(string(s[i]))
		} else {
			b.WriteString("…")
			break
		}
	}

	return b.String()
}
