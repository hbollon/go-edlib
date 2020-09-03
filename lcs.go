package edlib

import (
	"errors"
)

// LCS takes two strings and compute their LCS(Longuest Subsequence Problem)
func LCS(str1, str2 string) int {
	// Convert strings to rune array to handle no-ASCII characters
	runeStr1 := []rune(str1)
	runeStr2 := []rune(str2)

	if len(runeStr1) == 0 || len(runeStr2) == 0 {
		return 0
	} else if equal(runeStr1, runeStr2) {
		return len(runeStr1)
	}

	lcsMatrix := lcsProcess(runeStr1, runeStr2)
	return lcsMatrix[len(runeStr1)][len(runeStr2)]
}

func lcsProcess(runeStr1, runeStr2 []rune) [][]int {
	// 2D Array that will contain str1 and str2 LCS
	lcsMatrix := make([][]int, len(runeStr1)+1)
	for i := 0; i <= len(runeStr1); i++ {
		lcsMatrix[i] = make([]int, len(runeStr2)+1)
		for j := 0; j <= len(runeStr2); j++ {
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

	return lcsMatrix
}

// LCSBacktrack returns all choices taken during LCS process
func LCSBacktrack(str1, str2 string) (string, error) {
	runeStr1 := []rune(str1)
	runeStr2 := []rune(str2)

	if len(runeStr1) == 0 || len(runeStr2) == 0 {
		return "", errors.New("Can't process and backtrack any LCS with empty string")
	} else if equal(runeStr1, runeStr2) {
		return str1, nil
	}

	lcsMatrix := make([][]int, len(runeStr1)+1)
	for i := 0; i <= len(runeStr1); i++ {
		lcsMatrix[i] = make([]int, len(runeStr2)+1)
		for j := 0; j <= len(runeStr2); j++ {
			lcsMatrix[i][j] = 0
		}
	}

	return processLCSBacktrack(str1, str2, lcsProcess(runeStr1, runeStr2), len(str1), len(str2)), nil
}

func processLCSBacktrack(str1 string, str2 string, lcsMatrix [][]int, m, n int) string {
	// Convert strings to rune array to handle no-ASCII characters
	runeStr1 := []rune(str1)
	runeStr2 := []rune(str2)

	if m == 0 || n == 0 {
		return ""
	} else {
		if runeStr1[m-1] == runeStr2[n-1] {
			return processLCSBacktrack(str1, str2, lcsMatrix, m-1, n-1) + string(str1[m-1])
		}
		if lcsMatrix[m][n-1] > lcsMatrix[m-1][n] {
			return processLCSBacktrack(str1, str2, lcsMatrix, m, n-1)
		}
	}

	return processLCSBacktrack(str1, str2, lcsMatrix, m-1, n)
}

// LCSBacktrackAll returns an array containing all common substrings between str1 and str2
func LCSBacktrackAll(str1, str2 string) ([]string, error) {
	runeStr1 := []rune(str1)
	runeStr2 := []rune(str2)

	if len(runeStr1) == 0 || len(runeStr2) == 0 {
		return nil, errors.New("Can't process and backtrack any LCS with empty string")
	} else if equal(runeStr1, runeStr2) {
		return []string{str1}, nil
	}

	lcsMatrix := make([][]int, len(runeStr1)+1)
	for i := 0; i <= len(runeStr1); i++ {
		lcsMatrix[i] = make([]int, len(runeStr2)+1)
		for j := 0; j <= len(runeStr2); j++ {
			lcsMatrix[i][j] = 0
		}
	}

	all := processLCSBacktrackAll(str1, str2, lcsProcess(runeStr1, runeStr2), len(str1), len(str2))
	var lcss []string
	var index int
	for key := range all {
		lcss = append(lcss, key)
		index++
	}
	return lcss, nil
}

func processLCSBacktrackAll(str1 string, str2 string, lcsMatrix [][]int, m, n int) map[string]struct{} {
	// Convert strings to rune array to handle no-ASCII characters
	runeStr1 := []rune(str1)
	runeStr2 := []rune(str2)

	// Map containing all commons substrings (Hash set builded from map)
	substrings := make(map[string]struct{})

	if m == 0 || n == 0 {
		substrings[""] = struct{}{}
	} else if runeStr1[m-1] == runeStr2[n-1] {
		for key := range processLCSBacktrackAll(str1, str2, lcsMatrix, m-1, n-1) {
			substrings[key+string(str1[m-1])] = struct{}{}
		}
	} else {
		if lcsMatrix[m-1][n] >= lcsMatrix[m][n-1] {
			for key := range processLCSBacktrackAll(str1, str2, lcsMatrix, m-1, n) {
				substrings[key] = struct{}{}
			}
		}
		if lcsMatrix[m][n-1] >= lcsMatrix[m-1][n] {
			for key := range processLCSBacktrackAll(str1, str2, lcsMatrix, m, n-1) {
				substrings[key] = struct{}{}
			}
		}
	}

	return substrings
}

// LCSEditDistance determines the edit distance between two strings using LCS function
// (allow only insert and delete operations)
func LCSEditDistance(str1, str2 string) int {
	if len(str1) == 0 {
		return len(str2)
	} else if len(str2) == 0 {
		return len(str1)
	} else if str1 == str2 {
		return 0
	}

	lcs := LCS(str1, str2)
	return (len(str1) - lcs) + (len(str2) - lcs)
}
