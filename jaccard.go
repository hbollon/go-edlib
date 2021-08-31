package edlib

import (
	"strings"
)

// JaccardSimilarity compute the jaccard similarity coeffecient between two strings
// Takes two strings as parameters and return an index.
// This algorithm is only effective between sentences and not unique words.
func JaccardSimilarity(str1, str2 string) float32 {
	// Split string before rune conversion for cosine calculation
	splittedStr1 := strings.Split(str1, " ")
	splittedStr2 := strings.Split(str2, " ")
	// Conversion of splitted string into rune array
	runeStr1 := make([][]rune, len(splittedStr1))
	for i, str := range splittedStr1 {
		runeStr1[i] = []rune(str)
	}
	runeStr2 := make([][]rune, len(splittedStr2))
	for i, str := range splittedStr2 {
		runeStr2[i] = []rune(str)
	}

	// Create union keywords slice between input strings
	unionStr := union(splittedStr1, splittedStr2)

	jacc := float32(len(runeStr1) + len(runeStr2) - len(unionStr))

	return jacc / float32(len(unionStr))
}
