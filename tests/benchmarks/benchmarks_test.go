package benchmarks

import (
	"fmt"
	"testing"

	edlib "github.com/tilotech/go-edlib"
)

type pair struct {
	first  string
	second string
}

var testingInputPairs []pair

func init() {
	testingInputPairs = []pair{
		{"sameLengthStringInput", "asmePenhtgTsrnigIpnut"},
		{"pneumonoultramicroscopicsilicovolcanoconiosis", "nneumonkultramicrrscoipcsilicocolvanocpnisis"},
		{"こにんちこにんちこにんちこにんち", "こんにちはこんにちはこんにちはこんにちはこんにちは"},
		{"I love horror movies", "Lights out is a horror movie"},
	}
}

func BenchmarkEdlibAlgorithms(tb *testing.B) {
	for _, pair := range testingInputPairs {
		fmt.Printf("\nBegin benchmark between %s/%s : \n", pair.first, pair.second)
		benchLevensthein(pair, tb)
		benchLCS(pair, tb)
		benchHamming(pair, tb)
		benchDamereauLevenshtein(pair, tb)
		benchOSADamereauLevenshtein(pair, tb)
		benchJaro(pair, tb)
		benchJaroWinkler(pair, tb)
		benchCosine(pair, tb)
	}
}

func benchLevensthein(testPair pair, tb *testing.B) {
	tb.ResetTimer()
	tb.Run(fmt.Sprintf("Levenshtein_%s/%s", testPair.first, testPair.second), func(tb *testing.B) {
		for i := 0; i < tb.N; i++ {
			edlib.StringsSimilarity(testPair.first, testPair.second, edlib.Levenshtein)
		}
	})
	sim, err := edlib.StringsSimilarity(testPair.first, testPair.second, edlib.Levenshtein)
	fmt.Printf("Sim: %f, Err: %s \n\n", sim, err)
}

func benchLCS(testPair pair, tb *testing.B) {
	tb.ResetTimer()
	tb.Run(fmt.Sprintf("LCS_%s/%s", testPair.first, testPair.second), func(tb *testing.B) {
		for i := 0; i < tb.N; i++ {
			edlib.StringsSimilarity(testPair.first, testPair.second, edlib.Lcs)
		}
	})
	sim, err := edlib.StringsSimilarity(testPair.first, testPair.second, edlib.Lcs)
	fmt.Printf("Sim: %f, Err: %s \n\n", sim, err)
}

func benchHamming(testPair pair, tb *testing.B) {
	tb.ResetTimer()
	tb.Run(fmt.Sprintf("Hamming_%s/%s", testPair.first, testPair.second), func(tb *testing.B) {
		for i := 0; i < tb.N; i++ {
			edlib.StringsSimilarity(testPair.first, testPair.second, edlib.Hamming)
		}
	})
	sim, err := edlib.StringsSimilarity(testPair.first, testPair.second, edlib.Hamming)
	fmt.Printf("Sim: %f, Err: %s \n\n", sim, err)
}

func benchDamereauLevenshtein(testPair pair, tb *testing.B) {
	tb.ResetTimer()
	tb.Run(fmt.Sprintf("DamerauLevenshtein_%s/%s", testPair.first, testPair.second), func(tb *testing.B) {
		for i := 0; i < tb.N; i++ {
			edlib.StringsSimilarity(testPair.first, testPair.second, edlib.DamerauLevenshtein)
		}
	})
	sim, err := edlib.StringsSimilarity(testPair.first, testPair.second, edlib.DamerauLevenshtein)
	fmt.Printf("Sim: %f, Err: %s \n\n", sim, err)
}

func benchOSADamereauLevenshtein(testPair pair, tb *testing.B) {
	tb.ResetTimer()
	tb.Run(fmt.Sprintf("OSADamerauLevenshtein_%s/%s", testPair.first, testPair.second), func(tb *testing.B) {
		for i := 0; i < tb.N; i++ {
			edlib.StringsSimilarity(testPair.first, testPair.second, edlib.OSADamerauLevenshtein)
		}
	})
	sim, err := edlib.StringsSimilarity(testPair.first, testPair.second, edlib.OSADamerauLevenshtein)
	fmt.Printf("Sim: %f, Err: %s \n\n", sim, err)
}

func benchJaro(testPair pair, tb *testing.B) {
	tb.ResetTimer()
	tb.Run(fmt.Sprintf("Jaro_%s/%s", testPair.first, testPair.second), func(tb *testing.B) {
		for i := 0; i < tb.N; i++ {
			edlib.StringsSimilarity(testPair.first, testPair.second, edlib.Jaro)
		}
	})
	sim, err := edlib.StringsSimilarity(testPair.first, testPair.second, edlib.Jaro)
	fmt.Printf("Sim: %f, Err: %s \n\n", sim, err)
}

func benchJaroWinkler(testPair pair, tb *testing.B) {
	tb.ResetTimer()
	tb.Run(fmt.Sprintf("JaroWinkler_%s/%s", testPair.first, testPair.second), func(tb *testing.B) {
		for i := 0; i < tb.N; i++ {
			edlib.StringsSimilarity(testPair.first, testPair.second, edlib.JaroWinkler)
		}
	})
	sim, err := edlib.StringsSimilarity(testPair.first, testPair.second, edlib.JaroWinkler)
	fmt.Printf("Sim: %f, Err: %s \n\n", sim, err)
}

func benchCosine(testPair pair, tb *testing.B) {
	tb.ResetTimer()
	tb.Run(fmt.Sprintf("Cosine_%s/%s", testPair.first, testPair.second), func(tb *testing.B) {
		for i := 0; i < tb.N; i++ {
			edlib.StringsSimilarity(testPair.first, testPair.second, edlib.Cosine)
		}
	})
	sim, err := edlib.StringsSimilarity(testPair.first, testPair.second, edlib.Cosine)
	fmt.Printf("Sim: %f, Err: %s \n\n", sim, err)
}
