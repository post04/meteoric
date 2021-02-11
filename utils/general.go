package utils

// Find a value inside the specified list.
func Find(list []string, value string) bool {
	for _, item := range list {
		if item == value {
			return true
		}
	}
	return false
}
