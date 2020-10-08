package benchmarks

import (
	"fmt"
	"testing"

	edlib "github.com/hbollon/go-edlib"
)

type pair struct {
	first  string
	second string
}

var testingInputPairs []pair

func init() {
	testingInputPairs = []pair{
		{"dixon", "dicksonx"},
		{"jellyfish", "smellyfish"},
		{"sameLengthStringInput", "asmePenhtgTsrnigIpnut"},
		{"pneumonoultramicroscopicsilicovolcanoconiosis", "nneumonkultramicrrscoipcsilicocolvanocpnisis"},
		{"こにんちこにんちこにんちこにんち", "こんにちはこんにちはこんにちはこんにちはこんにちは"},
		{"🙂😄🙂😄🙂😄🙂😄🙂😄🙂😄", "😄🙂😄🙂😄🙂😄🙂😄🙂😄🙂"},
	}
}

func BenchmarkEdlibAlgorithms(tb *testing.B) {
	for _, pair := range testingInputPairs {
		fmt.Printf("\nBegin benchmark between '%s'/'%s' : \n", pair.first, pair.second)
		benchLevensthein(pair, tb)
		benchLCS(pair, tb)
		benchHamming(pair, tb)
		benchDamereauLevenshtein(pair, tb)
		benchOSADamereauLevenshtein(pair, tb)
		benchJaro(pair, tb)
		benchJaroWinkler(pair, tb)
	}
}

func benchLevensthein(testPair pair, tb *testing.B) {
	tb.ResetTimer()
	tb.Run(fmt.Sprintf("Levenshtein_'%s'/'%s'", testPair.first, testPair.second), func(tb *testing.B) {
		for i := 0; i < tb.N; i++ {
			edlib.StringsSimilarity(testPair.first, testPair.second, edlib.Levenshtein)
		}
	})
}

func benchLCS(testPair pair, tb *testing.B) {
	tb.ResetTimer()
	tb.Run(fmt.Sprintf("LCS_'%s'/'%s'", testPair.first, testPair.second), func(tb *testing.B) {
		for i := 0; i < tb.N; i++ {
			edlib.StringsSimilarity(testPair.first, testPair.second, edlib.Lcs)
		}
	})
}

func benchHamming(testPair pair, tb *testing.B) {
	tb.ResetTimer()
	tb.Run(fmt.Sprintf("Hamming_'%s'/'%s'", testPair.first, testPair.second), func(tb *testing.B) {
		for i := 0; i < tb.N; i++ {
			edlib.StringsSimilarity(testPair.first, testPair.second, edlib.Hamming)
		}
	})
}

func benchDamereauLevenshtein(testPair pair, tb *testing.B) {
	tb.ResetTimer()
	tb.Run(fmt.Sprintf("DamerauLevenshtein_'%s'/'%s'", testPair.first, testPair.second), func(tb *testing.B) {
		for i := 0; i < tb.N; i++ {
			edlib.StringsSimilarity(testPair.first, testPair.second, edlib.DamerauLevenshtein)
		}
	})
}

func benchOSADamereauLevenshtein(testPair pair, tb *testing.B) {
	tb.ResetTimer()
	tb.Run(fmt.Sprintf("OSADamerauLevenshtein_'%s'/'%s'", testPair.first, testPair.second), func(tb *testing.B) {
		for i := 0; i < tb.N; i++ {
			edlib.StringsSimilarity(testPair.first, testPair.second, edlib.OSADamerauLevenshtein)
		}
	})
}

func benchJaro(testPair pair, tb *testing.B) {
	tb.ResetTimer()
	tb.Run(fmt.Sprintf("Jaro_'%s'/'%s'", testPair.first, testPair.second), func(tb *testing.B) {
		for i := 0; i < tb.N; i++ {
			edlib.StringsSimilarity(testPair.first, testPair.second, edlib.Jaro)
		}
	})
}

func benchJaroWinkler(testPair pair, tb *testing.B) {
	tb.ResetTimer()
	tb.Run(fmt.Sprintf("JaroWinkler_'%s'/'%s'", testPair.first, testPair.second), func(tb *testing.B) {
		for i := 0; i < tb.N; i++ {
			edlib.StringsSimilarity(testPair.first, testPair.second, edlib.JaroWinkler)
		}
	})
}

// func benchCosine(testPair pair, tb *testing.B) {
// 	b.ResetTimer()
// 	b.Run(fmt.Sprintf("Levenshtein_%s/%s", testPair.first, testPair.second), func(b *testing.B) {
// 		edlib.LevenshteinDistance(testPair.first, testPair.second)
// 	})
// }
