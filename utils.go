package edlib

// Return the smallest integer among the two in parameters
func min(a int, b int) int {
	if b < a {
		return b
	}
	return a
}

// Return the largest integer among the two in parameters
func max(a int, b int) int {
	if b > a {
		return b
	}
	return a
}

// Compare two rune arrays and return if they are equals or not
func equal(a, b []rune) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
