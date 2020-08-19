package edlib

// LCS takes two strings and compute their LCS(Longuest Subsequence Problem)
func LCS(str1, str2 string) int {
	// Convert strings to rune array to handle no-ASCII characters
	runeStr1 := []rune(str1)
	runeStr2 := []rune(str2)

	// 2D Array that will contain str1 and str2 LCS
	lcsMatrix := make([][]int, len(runeStr1)+1) // Make the outer slice and give it size 10
	for i := 0; i < len(runeStr1)+1; i++ {
		lcsMatrix[i] = make([]int, len(runeStr2)+1)
		for j := 0; j < len(runeStr2)+1; j++ {
			lcsMatrix[i][j] = 0
		}
	}

	for i := 1; i <= len(runeStr1); i++ {
		for j := 1; j <= len(runeStr2); j++ {
			if runeStr1[i-1] == runeStr2[j-1] {
				lcsMatrix[i][j] = lcsMatrix[i-1][j-1] + 1
			} else {
				lcsMatrix[i][j] = max(lcsMatrix[i][j-1], lcsMatrix[i-1][j])
			}
		}
	}

	return lcsMatrix[len(runeStr1)][len(runeStr2)]
}

// LCSEditDistance determines the edit distance between two strings using LCS function
// (allow only insert and delete operations)
func LCSEditDistance(str1, str2 string) int {
	lcs := LCS(str1, str2)
	return (len(str1) - lcs) + (len(str2) - lcs)
}
