package edlib

import (
	"reflect"
	"testing"

	"github.com/hbollon/go-edlib"
)

var strList []string

func init() {
	strList = []string{
		"test",
		"tester",
		"tests",
		"testers",
		"testing",
		"tsting",
		"sting",
	}
}

func TestStringsSimilarity(t *testing.T) {
	type args struct {
		str1 string
		str2 string
		algo edlib.AlgorithMethod
	}
	tests := []struct {
		name    string
		args    args
		want    float32
		wantErr bool
	}{
		// Levenshtein method
		{"Levenshtein : First arg empty", args{"", "abcde", edlib.Levenshtein}, 0.0, false},
		{"Levenshtein : Second arg empty", args{"abcde", "", edlib.Levenshtein}, 0.0, false},
		{"Levenshtein : Same args", args{"abcde", "abcde", edlib.Levenshtein}, 1.0, false},
		{"Levenshtein : No characters match", args{"abcd", "effgghh", edlib.Levenshtein}, 0.0, false},
		{"Levenshtein : CRATE/TRACE", args{"CRATE", "TRACE", edlib.Levenshtein}, 0.6, false},
		{"Levenshtein : MARTHA/MARHTA", args{"MARTHA", "MARHTA", edlib.Levenshtein}, 0.6666667, false},
		{"Levenshtein : DIXON/DICKSONX", args{"DIXON", "DICKSONX", edlib.Levenshtein}, 0.50, false},
		{"Levenshtein : jellyfish/smellyfish", args{"jellyfish", "smellyfish", edlib.Levenshtein}, 0.80, false},

		// DamerauLevenshtein method
		{"DamerauLevenshtein : First arg empty", args{"", "abcde", edlib.DamerauLevenshtein}, 0.0, false},
		{"DamerauLevenshtein : Second arg empty", args{"abcde", "", edlib.DamerauLevenshtein}, 0.0, false},
		{"DamerauLevenshtein : Same args", args{"abcde", "abcde", edlib.DamerauLevenshtein}, 1.0, false},
		{"DamerauLevenshtein : No characters match", args{"abcd", "effgghh", edlib.DamerauLevenshtein}, 0.0, false},
		{"DamerauLevenshtein : ab/aaa", args{"ab", "aaa", edlib.DamerauLevenshtein}, 0.33333334, false},
		{"DamerauLevenshtein : bbb/a", args{"bbb", "a", edlib.DamerauLevenshtein}, 0.0, false},
		{"DamerauLevenshtein : ca/abc", args{"ca", "abc", edlib.DamerauLevenshtein}, 0.33333334, false},
		{"DamerauLevenshtein : a cat/an abct", args{"a cat", "an abct", edlib.DamerauLevenshtein}, 0.5714286, false},
		{"DamerauLevenshtein : dixon/dicksonx", args{"dixon", "dicksonx", edlib.DamerauLevenshtein}, 0.5, false},
		{"DamerauLevenshtein : jellyfish/smellyfish", args{"jellyfish", "smellyfish", edlib.DamerauLevenshtein}, 0.8, false},
		{"DamerauLevenshtein : こにんち/こんにちは", args{"こにんち", "こんにちは", edlib.DamerauLevenshtein}, 0.8666667, false}, // "Hello" in Japanese
		{"DamerauLevenshtein : 🙂😄🙂😄/😄🙂😄🙂", args{"🙂😄🙂😄", "😄🙂😄🙂", edlib.DamerauLevenshtein}, 0.875, false},

		// OSADamerauLevenshtein method
		{"OSADamerauLevenshtein : First arg empty", args{"", "abcde", edlib.OSADamerauLevenshtein}, 0.0, false},
		{"OSADamerauLevenshtein : Second arg empty", args{"abcde", "", edlib.OSADamerauLevenshtein}, 0.0, false},
		{"OSADamerauLevenshtein : Same args", args{"abcde", "abcde", edlib.OSADamerauLevenshtein}, 1.0, false},
		{"OSADamerauLevenshtein : No characters match", args{"abcd", "effgghh", edlib.OSADamerauLevenshtein}, 0.0, false},
		{"OSADamerauLevenshtein : ab/aaa", args{"ab", "aaa", edlib.OSADamerauLevenshtein}, 0.33333334, false},
		{"OSADamerauLevenshtein : bbb/a", args{"bbb", "a", edlib.OSADamerauLevenshtein}, 0.0, false},
		{"OSADamerauLevenshtein : ca/abc", args{"ca", "abc", edlib.OSADamerauLevenshtein}, 0.0, false},
		{"OSADamerauLevenshtein : a cat/an abct", args{"a cat", "an abct", edlib.OSADamerauLevenshtein}, 0.428571429, false},
		{"OSADamerauLevenshtein : dixon/dicksonx", args{"dixon", "dicksonx", edlib.OSADamerauLevenshtein}, 0.5, false},
		{"OSADamerauLevenshtein : jellyfish/smellyfish", args{"jellyfish", "smellyfish", edlib.OSADamerauLevenshtein}, 0.8, false},
		{"OSADamerauLevenshtein : こにんち/こんにちは", args{"こにんち", "こんにちは", edlib.OSADamerauLevenshtein}, 0.8666667, false}, // "Hello" in Japanese
		{"OSADamerauLevenshtein : 🙂😄🙂😄/😄🙂😄🙂", args{"🙂😄🙂😄", "😄🙂😄🙂", edlib.OSADamerauLevenshtein}, 0.875, false},

		// Lcs method
		{"LCS : First arg empty", args{"", "abcde", edlib.Lcs}, 0.0, false},
		{"LCS : Second arg empty", args{"abcde", "", edlib.Lcs}, 0.0, false},
		{"LCS : Same args", args{"abcde", "abcde", edlib.Lcs}, 1.0, false},
		{"LCS : ABCDGH/AEDFHR", args{"ABCDGH", "AEDFHR", edlib.Lcs}, 0.0, false},
		{"LCS : AGGTAB/GXTXAYB", args{"AGGTAB", "GXTXAYB", edlib.Lcs}, 0.2857143, false},
		{"LCS : XMJYAUZ/MZJAWXU", args{"XMJYAUZ", "MZJAWXU", edlib.Lcs}, 0.14285715, false},
		{"LCS : CRATE/TRACE", args{"CRATE", "TRACE", edlib.Lcs}, 0.2, false},
		{"LCS : MARTHA/MARHTA", args{"MARTHA", "MARHTA", edlib.Lcs}, 0.6666667, false},
		{"LCS : DIXON/DICKSONX", args{"DIXON", "DICKSONX", edlib.Lcs}, 0.375, false},
		{"LCS : jellyfish/smellyfish", args{"jellyfish", "smellyfish", edlib.Lcs}, 0.7, false},

		// Hamming method
		{"Hamming : First arg empty", args{"", "abcde", edlib.Hamming}, 0.0, true},
		{"Hamming : Second arg empty", args{"abcde", "", edlib.Hamming}, 0.0, true},
		{"Hamming : Same args", args{"abcde", "abcde", edlib.Hamming}, 1.0, false},
		{"Hamming : No characters match", args{"abcd", "effgghh", edlib.Hamming}, 0.0, true},
		{"Hamming : aa/aa", args{"aa", "aa", edlib.Hamming}, 1.0, false},
		{"Hamming : ab/aa", args{"ab", "aa", edlib.Hamming}, 0.5, false},
		{"Hamming : ab/ba", args{"ab", "ba", edlib.Hamming}, 0.0, false},
		{"Hamming : a cat/an abct", args{"a cat", "an abct", edlib.Hamming}, 0.0, true},
		{"Hamming : dixon/dicksonx", args{"dixon", "dicksonx", edlib.Hamming}, 0.0, true},
		{"Hamming : jellyfish/smellyfish", args{"jellyfish", "smellyfish", edlib.Hamming}, 0.0, true},
		{"Hamming : こにんち/こんにちは", args{"こにんち", "こんにちは", edlib.Hamming}, 0.0, true}, // "Hello" in Japanese
		{"Hamming : 🙂😄🙂😄/😄🙂😄🙂", args{"🙂😄🙂😄", "😄🙂😄🙂", edlib.Hamming}, 0.75, false},

		// Jaro method
		{"Jaro : First arg empty", args{"", "abcde", edlib.Jaro}, 0.0, false},
		{"Jaro : Second arg empty", args{"abcde", "", edlib.Jaro}, 0.0, false},
		{"Jaro : Same args", args{"abcde", "abcde", edlib.Jaro}, 1.0, false},
		{"Jaro : No characters match", args{"abcd", "effgghh", edlib.Jaro}, 0.0, false},
		{"Jaro : CRATE/TRACE", args{"CRATE", "TRACE", edlib.Jaro}, 0.73333335, false},
		{"Jaro : MARTHA/MARHTA", args{"MARTHA", "MARHTA", edlib.Jaro}, 0.9444444, false},
		{"Jaro : DIXON/DICKSONX", args{"DIXON", "DICKSONX", edlib.Jaro}, 0.76666665, false},
		{"Jaro : jellyfish/smellyfish", args{"jellyfish", "smellyfish", edlib.Jaro}, 0.8962963, false},

		// JaroWinkler method
		{"JaroWinkler : First arg empty", args{"", "abcde", edlib.JaroWinkler}, 0.0, false},
		{"JaroWinkler : Second arg empty", args{"abcde", "", edlib.JaroWinkler}, 0.0, false},
		{"JaroWinkler : Same args", args{"abcde", "abcde", edlib.JaroWinkler}, 1.0, false},
		{"JaroWinkler : No characters match", args{"abcd", "effgghh", edlib.JaroWinkler}, 0.0, false},
		{"JaroWinkler : CRATE/TRACE", args{"CRATE", "TRACE", edlib.JaroWinkler}, 0.73333335, false},
		{"JaroWinkler : MARTHA/MARHTA", args{"MARTHA", "MARHTA", edlib.JaroWinkler}, 0.96111107, false},
		{"JaroWinkler : DIXON/DICKSONX", args{"DIXON", "DICKSONX", edlib.JaroWinkler}, 0.81333333, false},
		{"JaroWinkler : jellyfish/smellyfish", args{"jellyfish", "smellyfish", edlib.JaroWinkler}, 0.8962963, false},

		// Illegal argument error
		{"Undefined integer value for method", args{"abc", "abcde", 42}, 0.0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := edlib.StringsSimilarity(tt.args.str1, tt.args.str2, tt.args.algo)
			if (err != nil) != tt.wantErr {
				t.Errorf("StringsSimilarity() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("StringsSimilarity() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFuzzySearch(t *testing.T) {
	type args struct {
		str     string
		strList []string
		algo    edlib.AlgorithMethod
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"FuzzySearch 'testing'", args{"testnig", strList, edlib.Levenshtein}, "testing"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := edlib.FuzzySearch(tt.args.str, tt.args.strList, tt.args.algo); got != tt.want {
				t.Errorf("FuzzySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFuzzySearchThreshold(t *testing.T) {
	type args struct {
		str     string
		strList []string
		minSim  float32
		algo    edlib.AlgorithMethod
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"FuzzySearch 'testing'", args{"testnig", strList, 0.7, edlib.Levenshtein}, "testing"},
		{"FuzzySearch 'testing'", args{"hello", strList, 0.7, edlib.Levenshtein}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := edlib.FuzzySearchThreshold(tt.args.str, tt.args.strList, tt.args.minSim, tt.args.algo); got != tt.want {
				t.Errorf("FuzzySearchThreshold() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFuzzySearchSet(t *testing.T) {
	type args struct {
		str      string
		strList  []string
		quantity int
		algo     edlib.AlgorithMethod
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"FuzzySearch 'testing'", args{"testnig", strList, 3, edlib.Levenshtein}, []string{"testing", "test", "tester"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := edlib.FuzzySearchSet(tt.args.str, tt.args.strList, tt.args.quantity, tt.args.algo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FuzzySearchSet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFuzzySearchSetThreshold(t *testing.T) {
	type args struct {
		str      string
		strList  []string
		quantity int
		minSim   float32
		algo     edlib.AlgorithMethod
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"FuzzySearch 'testing'", args{"testnig", strList, 3, 0.7, edlib.Levenshtein}, []string{"testing", "", ""}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := edlib.FuzzySearchSetThreshold(tt.args.str, tt.args.strList, tt.args.quantity, tt.args.minSim, tt.args.algo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FuzzySearchSetThreshold() = %v, want %v", got, tt.want)
			}
		})
	}
}
