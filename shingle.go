package edlib

func Shingle(s string, k int) map[string]int {
	m := make(map[string]int)
	if s != "" && k != 0 {
		rune_of_s := []rune(s)

		for i := 0; i < len(rune_of_s)-k+1; i++ {
			m[string(rune_of_s[i:i+k])]++
		}
	}
	return m
}

func ShingleSlidingWindow(s string, k int) map[string]int {
	m := make(map[string]int)
	if s != "" && k != 0 {
		rune_of_s := []rune(s)

		window := rune_of_s[:k]
		m[string(window)]++

		for i := 1; i < len(rune_of_s)-k+1; i++ {
			window = window[1:]
			window = append(window, rune_of_s[i+k-1])
			m[string(window)]++
		}
	}

	return m

}
