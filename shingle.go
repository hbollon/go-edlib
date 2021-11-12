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

// ShingleSlidingWindow Find the k-gram of a string for a given k using a sliding window of size k
// Takes a string and an integer as parameters and return a map.
// Returns an empty map if the string is empty or if k is 0
func ShingleSlidingWindow(s string, k int) map[string]int {
	m := make(map[string]int)
	if s != "" && k != 0 {
		runeS := []rune(s)

		window := runeS[:k]
		m[string(window)]++

		for i := 1; i < len(runeS)-k+1; i++ {
			window = window[1:]
			window = append(window, runeS[i+k-1])
			m[string(window)]++
		}
	}
	return m
}
