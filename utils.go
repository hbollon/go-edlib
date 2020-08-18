package edlib

// Min return the smallest integer among the two in parameters
func Min(a int, b int) int {
	if b < a {
		return b
	}
	return a
}

// Max return the largest integer among the two in parameters
func Max(a int, b int) int {
	if b > a {
		return b
	}
	return a
}
