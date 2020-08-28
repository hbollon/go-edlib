package edlib

func JaroSimilarity(str1, str2 string) float32 {
	// Convert string parameters to rune arrays to be compatible with non-ASCII
	runeStr1 := []rune(str1)
	runeStr2 := []rune(str2)

	// Get and store length of these strings
	runeStr1len := len(runeStr1)
	runeStr2len := len(runeStr2)
	if runeStr1len == 0 || runeStr2len == 0 {
		return 0.0
	} else if equal(runeStr1, runeStr2) {
		return 1.0
	}

	var match int
	// Maximum matching distance allowed
	maxDist := max(runeStr1len, runeStr2len)/2 - 1
	// Correspondence tables (1 for matching and 0 if it's not the case)
	str1Table := make([]int, runeStr1len)
	str2Table := make([]int, runeStr2len)

	// Check for matching characters in both strings
	for i := 0; i < runeStr1len; i++ {
		for j := max(0, i-maxDist); j < min(runeStr2len, i+maxDist+1); j++ {
			if runeStr1[i] == runeStr2[j] && str2Table[j] == 0 {
				str1Table[i] = 1
				str2Table[j] = 1
				match++
				break
			}
		}
	}
	if match == 0 {
		return 0.0
	}

	var t float32
	var p int
	// Check for possible translations
	for i := 0; i < runeStr1len; i++ {
		if str1Table[i] == 1 {
			for str2Table[p] == 0 {
				p++
			}
			if runeStr1[i] != runeStr2[p] {
				t++
			}
			p++
		}
	}
	t /= 2

	return (float32(match)/float32(runeStr1len) +
		float32(match)/float32(runeStr2len) +
		(float32(match)-t)/float32(match)) / 3.0
}
