package edlib

// LevenshteinDistance calculate the distance between two string
// This algorithm allow insertions, deletions and substitutions to change one string to the second
// Compatible with non-ASCII characters
func LevenshteinDistance(str1, str2 string) int {
	// Convert string parameters to rune arrays to be compatible with non-ASCII
	runeStr1 := []rune(str1)
	runeStr2 := []rune(str2)

	// Get and store length of these strings
	runeStr1len := len(runeStr1)
	runeStr2len := len(runeStr2)
	if runeStr1len == 0 {
		return runeStr2len
	} else if runeStr2len == 0 {
		return runeStr1len
	} else if equal(runeStr1, runeStr2) {
		return 0
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
			column[y] = min(min(column[y]+1, column[y-1]+1), lastkey+i)
			lastkey = oldkey
		}
	}

	return column[runeStr1len]
}

// DamerauLevenshteinDistance calculate the distance between two string
// Optimal string alignment distance variant that use extention of the Wagner-Fisher dynamic programming algorithm
// Allowing insertions, deletions, substitutions and transpositions to change one string to the second
// Compatible with non-ASCII characters
func OSADamerauLevenshteinDistance(str1, str2 string) int {
	// Convert string parameters to rune arrays to be compatible with non-ASCII
	runeStr1 := []rune(str1)
	runeStr2 := []rune(str2)

	// 2D Array
	matrix := make([][]int, len(runeStr1))
	for i := 0; i < len(runeStr1); i++ {
		matrix[i] = make([]int, len(runeStr2))
		for j := 0; j < len(runeStr2); j++ {
			matrix[i][j] = 0
		}
	}

	var count int
	for i := 1; i < len(runeStr1); i++ {
		for j := 1; j < len(runeStr2); j++ {
			if runeStr1[i] == runeStr2[j] {
				count = 0
			} else {
				count = 1
			}

			matrix[i][j] = min(min(matrix[i-1][j], matrix[i][j-1]), matrix[i-1][j-1]+count)
			if i > 1 && j > 1 && runeStr1[i] == runeStr2[j-1] && runeStr1[i-1] == runeStr2[j] {
				matrix[i][j] = min(matrix[i][j], matrix[i-2][j-2]+1)
			}
		}
	}
	return matrix[len(runeStr1)][len(runeStr2)]
}
