package server

import "regexp"

// Returns true if the input string is acceptable from a security point of view.
func isValidString(input string) bool {
	var isValidString = regexp.MustCompile(`^[1-9a-z-]+$`).MatchString

	if isValidString(input) {
		return true
	}
	return false
}

// Returns true if the namespace and name are acceptable.
func isValidNameAndNamespace(namespace, name string) bool {
	if isValidString(namespace) && isValidString(name) && len(namespace) > 0 && len(name) > 0 {
		return true
	}
	return false
}
