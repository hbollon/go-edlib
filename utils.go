package edlib

// StringHashMap is HashMap substitue for string
type StringHashMap map[string]struct{}

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

/*
	StringHashMap methods
*/

func (m StringHashMap) addAll(srcMap StringHashMap) {
	for key := range srcMap {
		m[key] = struct{}{}
	}
}

// Convert and return an StringHashMap to string array
func (m StringHashMap) toArray() []string {
	var arr []string
	var index int
	for key := range m {
		arr = append(arr, key)
		index++
	}

	return arr
}
