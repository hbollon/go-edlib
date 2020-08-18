package edlib

// levenshteinDistance calculate the distance between two string
// This algorithm allow insertions, deletions and substitutions to change one string to the second
// Compatible with non-ASCII characters
func levenshteinDistance(str1, str2 string) int {
	// Convert string parameters to rune arrays to be compatible with non-ASCII
	runeStr1 := []rune(str1)
	runeStr2 := []rune(str2)

	// Get and store lenght of these strings
	runeStr1len := len(runeStr1)
	runeStr2len := len(runeStr2)
	if runeStr1len == 0 {
		return runeStr2len
	} else if runeStr2len == 0 {
		return runeStr1len
	}

	column := make([]int, runeStr1len+1)

	for y := 1; y <= runeStr1len; y++ {
		column[y] = y
	}
	for x := 1; x <= runeStr2len; x++ {
		column[0] = x
		lastkey := x - 1
		for y := 1; y <= runeStr1len; y++ {
			oldkey := column[y]
			var i int
			if runeStr1[y-1] != runeStr2[x-1] {
				i = 1
			}
			column[y] = Min(column[y]+1, column[y-1]+1, lastkey+i)
			lastkey = oldkey
		}
	}

	return column[runeStr1len]
}
