package edlib

import (
	"errors"
	"log"
)

// AlgorithMethod is an Integer type used to identify edit distance algorithms
type AlgorithMethod uint8

const (
	Levenshtein           AlgorithMethod = iota
	DamerauLevenshtein    AlgorithMethod = iota
	OSADamerauLevenshtein AlgorithMethod = iota
	Lcs                   AlgorithMethod = iota
	Hamming               AlgorithMethod = iota
	Jaro                  AlgorithMethod = iota
	JaroWinkler           AlgorithMethod = iota
)

// StringsSimilarity return a similarity index [0..1] between two strings based on given edit distance algorithm in parameter.
// Use defined AlgorithmMethod type.
func StringsSimilarity(str1 string, str2 string, algo AlgorithMethod) (float32, error) {
	switch algo {
	case Levenshtein:
		return matchingIndex(str1, str2, LevenshteinDistance(str1, str2)), nil
	case DamerauLevenshtein:
		return matchingIndex(str1, str2, DamerauLevenshteinDistance(str1, str2)), nil
	case OSADamerauLevenshtein:
		return matchingIndex(str1, str2, OSADamerauLevenshteinDistance(str1, str2)), nil
	case Lcs:
		return matchingIndex(str1, str2, LCSEditDistance(str1, str2)), nil
	case Hamming:
		distance, err := HammingDistance(str1, str2)
		if err == nil {
			return matchingIndex(str1, str2, distance), nil
		}
		return 0.0, err
	case Jaro:
		return JaroSimilarity(str1, str2), nil
	case JaroWinkler:
		return JaroWinklerSimilarity(str1, str2), nil
	default:
		return 0.0, errors.New("Illegal argument for algorithm method")
	}
}

// Return matching index E [0..1] from two strings and an edit distance
func matchingIndex(str1 string, str2 string, distance int) float32 {
	// Compare strings length and make a matching percentage between them
	if len(str1) >= len(str2) {
		return float32(len(str1)-distance) / float32(len(str1))
	}
	return float32(len(str2)-distance) / float32(len(str2))
}

// FuzzySearch realize an approximate search on a string list and return the closest one compared
// to the string input
func FuzzySearch(str string, strList []string, algo AlgorithMethod) string {
	var higherMatchPercent float32
	var tmpStr string
	for _, strToCmp := range strList {
		sim, err := StringsSimilarity(str, strToCmp, algo)
		if err != nil {
			log.Fatal(err)
		}

		if sim == 1.0 {
			return strToCmp
		} else if sim > higherMatchPercent {
			higherMatchPercent = sim
			tmpStr = strToCmp
		}
	}

	return tmpStr
}

// FuzzySearchThreshold realize an approximate search on a string list and return the closest one compared
// to the string input. Take an similarity threshold in parameter.
func FuzzySearchThreshold(str string, strList []string, minSim float32, algo AlgorithMethod) string {
	var higherMatchPercent float32
	var tmpStr string
	for _, strToCmp := range strList {
		sim, err := StringsSimilarity(str, strToCmp, algo)
		if err != nil {
			log.Fatal(err)
		}

		if sim == 1.0 {
			return strToCmp
		} else if sim > higherMatchPercent && sim >= minSim {
			higherMatchPercent = sim
			tmpStr = strToCmp
		}
	}
	return tmpStr
}
