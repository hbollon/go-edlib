package edlib

import (
	"errors"

	"slices"
)

// HammingDistance calculate the edit distance between two given strings using only substitutions
// Return edit distance integer and an error
func HammingDistance(str1, str2 string) (int, error) {
	// Convert strings to rune array to handle no-ASCII characters
	runeStr1 := []rune(str1)
	runeStr2 := []rune(str2)

	if len(runeStr1) != len(runeStr2) {
		return 0, errors.New("undefined for strings of unequal length")
	} else if slices.Equal(runeStr1, runeStr2) {
		return 0, nil
	}

	var counter int
	for i := range runeStr1 {
		if runeStr1[i] != runeStr2[i] {
			counter++
		}
	}

	return counter, nil
}
