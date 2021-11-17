package edlib

// Shingle Find the k-gram of a string for a given k
// Takes a string and an integer as parameters and return a map.
// Returns an empty map if the string is empty or if k is 0
func Shingle(s string, k int) map[string]int {
	m := make(map[string]int)
	if s != "" && k != 0 {
		runeS := []rune(s)

		for i := 0; i < len(runeS)-k+1; i++ {
			m[string(runeS[i:i+k])]++
		}
	}
	return m
}
