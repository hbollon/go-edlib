package edlib

import "fmt"

func JaroSimilarity(str1, str2 string) float32 {
	// Convert string parameters to rune arrays to be compatible with non-ASCII
	runeStr1 := []rune(str1)
	runeStr2 := []rune(str2)

	// Get and store length of these strings
	runeStr1len := len(runeStr1)
	runeStr2len := len(runeStr2)
	maxStrLength := max(runeStr1len, runeStr2len)
	if runeStr1len == 0 || runeStr2len == 0 {
		return 0.0
	} else if equal(runeStr1, runeStr2) {
		return 1.0
	}

	var match int
	// Maximum matching distance allowed
	maxDist := maxStrLength/2 - 1
	// Correspondence matrix (1 for matching and 0 if it's not the case)
	// [1..lenStr1+1][1..lenStr2+1]
	matchingMatrix := make([][]int, maxStrLength+1)
	for i := 0; i <= maxStrLength; i++ {
		matchingMatrix[i] = make([]int, maxStrLength+1)
		for j := 0; j <= maxStrLength; j++ {
			matchingMatrix[i][j] = 0
		}
	}

	// Check for matching characters in both strings
	for i := 0; i < runeStr1len; i++ {
		for j := max(0, i-maxDist); j < min(runeStr2len, i+maxDist+1); j++ {
			if runeStr1[i] == runeStr2[j] && matchingMatrix[i+1][j+1] == 0 {
				matchingMatrix[i+1][j+1] = 1
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
		if matchingMatrix[i+1][i+1] == 1 {
			for matchingMatrix[i+1][p+1] == 0 {
				p++
			}
			if runeStr1[i] != runeStr2[p] {
				t++
			}
		}
	}
	t /= 2

	fmt.Printf("Jaro (%s/%s): %f\n", str1, str2, (float32(match)/float32(runeStr1len)+float32(match)/float32(runeStr2len)+(float32(match)-t)/float32(match))/3.0)
	return (float32(match)/float32(runeStr1len) + float32(match)/float32(runeStr2len) + (float32(match)-t)/float32(match)) / 3.0
}
