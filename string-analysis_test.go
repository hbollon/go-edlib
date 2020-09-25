package edlib

import (
	"reflect"
	"testing"
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
		algo Algorithm
	}
	tests := []struct {
		name    string
		args    args
		want    float32
		wantErr bool
	}{
		// Levenshtein method
		{"Levenshtein : First arg empty", args{"", "abcde", Levenshtein}, 0.0, false},
		{"Levenshtein : Second arg empty", args{"abcde", "", Levenshtein}, 0.0, false},
		{"Levenshtein : Same args", args{"abcde", "abcde", Levenshtein}, 1.0, false},
		{"Levenshtein : No characters match", args{"abcd", "effgghh", Levenshtein}, 0.0, false},
		{"Levenshtein : CRATE/TRACE", args{"CRATE", "TRACE", Levenshtein}, 0.6, false},
		{"Levenshtein : MARTHA/MARHTA", args{"MARTHA", "MARHTA", Levenshtein}, 0.6666667, false},
		{"Levenshtein : DIXON/DICKSONX", args{"DIXON", "DICKSONX", Levenshtein}, 0.50, false},
		{"Levenshtein : jellyfish/smellyfish", args{"jellyfish", "smellyfish", Levenshtein}, 0.80, false},

		// DamerauLevenshtein method
		{"DamerauLevenshtein : First arg empty", args{"", "abcde", DamerauLevenshtein}, 0.0, false},
		{"DamerauLevenshtein : Second arg empty", args{"abcde", "", DamerauLevenshtein}, 0.0, false},
		{"DamerauLevenshtein : Same args", args{"abcde", "abcde", DamerauLevenshtein}, 1.0, false},
		{"DamerauLevenshtein : No characters match", args{"abcd", "effgghh", DamerauLevenshtein}, 0.0, false},
		{"DamerauLevenshtein : ab/aaa", args{"ab", "aaa", DamerauLevenshtein}, 0.33333334, false},
		{"DamerauLevenshtein : bbb/a", args{"bbb", "a", DamerauLevenshtein}, 0.0, false},
		{"DamerauLevenshtein : ca/abc", args{"ca", "abc", DamerauLevenshtein}, 0.33333334, false},
		{"DamerauLevenshtein : a cat/an abct", args{"a cat", "an abct", DamerauLevenshtein}, 0.5714286, false},
		{"DamerauLevenshtein : dixon/dicksonx", args{"dixon", "dicksonx", DamerauLevenshtein}, 0.5, false},
		{"DamerauLevenshtein : jellyfish/smellyfish", args{"jellyfish", "smellyfish", DamerauLevenshtein}, 0.8, false},
		{"DamerauLevenshtein : こにんち/こんにちは", args{"こにんち", "こんにちは", DamerauLevenshtein}, 0.8666667, false}, // "Hello" in Japanese
		{"DamerauLevenshtein : 🙂😄🙂😄/😄🙂😄🙂", args{"🙂😄🙂😄", "😄🙂😄🙂", DamerauLevenshtein}, 0.875, false},

		// OSADamerauLevenshtein method
		{"OSADamerauLevenshtein : First arg empty", args{"", "abcde", OSADamerauLevenshtein}, 0.0, false},
		{"OSADamerauLevenshtein : Second arg empty", args{"abcde", "", OSADamerauLevenshtein}, 0.0, false},
		{"OSADamerauLevenshtein : Same args", args{"abcde", "abcde", OSADamerauLevenshtein}, 1.0, false},
		{"OSADamerauLevenshtein : No characters match", args{"abcd", "effgghh", OSADamerauLevenshtein}, 0.0, false},
		{"OSADamerauLevenshtein : ab/aaa", args{"ab", "aaa", OSADamerauLevenshtein}, 0.33333334, false},
		{"OSADamerauLevenshtein : bbb/a", args{"bbb", "a", OSADamerauLevenshtein}, 0.0, false},
		{"OSADamerauLevenshtein : ca/abc", args{"ca", "abc", OSADamerauLevenshtein}, 0.0, false},
		{"OSADamerauLevenshtein : a cat/an abct", args{"a cat", "an abct", OSADamerauLevenshtein}, 0.428571429, false},
		{"OSADamerauLevenshtein : dixon/dicksonx", args{"dixon", "dicksonx", OSADamerauLevenshtein}, 0.5, false},
		{"OSADamerauLevenshtein : jellyfish/smellyfish", args{"jellyfish", "smellyfish", OSADamerauLevenshtein}, 0.8, false},
		{"OSADamerauLevenshtein : こにんち/こんにちは", args{"こにんち", "こんにちは", OSADamerauLevenshtein}, 0.8666667, false}, // "Hello" in Japanese
		{"OSADamerauLevenshtein : 🙂😄🙂😄/😄🙂😄🙂", args{"🙂😄🙂😄", "😄🙂😄🙂", OSADamerauLevenshtein}, 0.875, false},

		// Lcs method
		{"LCS : First arg empty", args{"", "abcde", Lcs}, 0.0, false},
		{"LCS : Second arg empty", args{"abcde", "", Lcs}, 0.0, false},
		{"LCS : Same args", args{"abcde", "abcde", Lcs}, 1.0, false},
		{"LCS : ABCDGH/AEDFHR", args{"ABCDGH", "AEDFHR", Lcs}, 0.0, false},
		{"LCS : AGGTAB/GXTXAYB", args{"AGGTAB", "GXTXAYB", Lcs}, 0.2857143, false},
		{"LCS : XMJYAUZ/MZJAWXU", args{"XMJYAUZ", "MZJAWXU", Lcs}, 0.14285715, false},
		{"LCS : CRATE/TRACE", args{"CRATE", "TRACE", Lcs}, 0.2, false},
		{"LCS : MARTHA/MARHTA", args{"MARTHA", "MARHTA", Lcs}, 0.6666667, false},
		{"LCS : DIXON/DICKSONX", args{"DIXON", "DICKSONX", Lcs}, 0.375, false},
		{"LCS : jellyfish/smellyfish", args{"jellyfish", "smellyfish", Lcs}, 0.7, false},

		// Hamming method
		{"Hamming : First arg empty", args{"", "abcde", Hamming}, 0.0, true},
		{"Hamming : Second arg empty", args{"abcde", "", Hamming}, 0.0, true},
		{"Hamming : Same args", args{"abcde", "abcde", Hamming}, 1.0, false},
		{"Hamming : No characters match", args{"abcd", "effgghh", Hamming}, 0.0, true},
		{"Hamming : aa/aa", args{"aa", "aa", Hamming}, 1.0, false},
		{"Hamming : ab/aa", args{"ab", "aa", Hamming}, 0.5, false},
		{"Hamming : ab/ba", args{"ab", "ba", Hamming}, 0.0, false},
		{"Hamming : a cat/an abct", args{"a cat", "an abct", Hamming}, 0.0, true},
		{"Hamming : dixon/dicksonx", args{"dixon", "dicksonx", Hamming}, 0.0, true},
		{"Hamming : jellyfish/smellyfish", args{"jellyfish", "smellyfish", Hamming}, 0.0, true},
		{"Hamming : こにんち/こんにちは", args{"こにんち", "こんにちは", Hamming}, 0.0, true}, // "Hello" in Japanese
		{"Hamming : 🙂😄🙂😄/😄🙂😄🙂", args{"🙂😄🙂😄", "😄🙂😄🙂", Hamming}, 0.75, false},

		// Jaro method
		{"Jaro : First arg empty", args{"", "abcde", Jaro}, 0.0, false},
		{"Jaro : Second arg empty", args{"abcde", "", Jaro}, 0.0, false},
		{"Jaro : Same args", args{"abcde", "abcde", Jaro}, 1.0, false},
		{"Jaro : No characters match", args{"abcd", "effgghh", Jaro}, 0.0, false},
		{"Jaro : CRATE/TRACE", args{"CRATE", "TRACE", Jaro}, 0.73333335, false},
		{"Jaro : MARTHA/MARHTA", args{"MARTHA", "MARHTA", Jaro}, 0.9444444, false},
		{"Jaro : DIXON/DICKSONX", args{"DIXON", "DICKSONX", Jaro}, 0.76666665, false},
		{"Jaro : jellyfish/smellyfish", args{"jellyfish", "smellyfish", Jaro}, 0.8962963, false},

		// JaroWinkler method
		{"JaroWinkler : First arg empty", args{"", "abcde", JaroWinkler}, 0.0, false},
		{"JaroWinkler : Second arg empty", args{"abcde", "", JaroWinkler}, 0.0, false},
		{"JaroWinkler : Same args", args{"abcde", "abcde", JaroWinkler}, 1.0, false},
		{"JaroWinkler : No characters match", args{"abcd", "effgghh", JaroWinkler}, 0.0, false},
		{"JaroWinkler : CRATE/TRACE", args{"CRATE", "TRACE", JaroWinkler}, 0.73333335, false},
		{"JaroWinkler : MARTHA/MARHTA", args{"MARTHA", "MARHTA", JaroWinkler}, 0.96111107, false},
		{"JaroWinkler : DIXON/DICKSONX", args{"DIXON", "DICKSONX", JaroWinkler}, 0.81333333, false},
		{"JaroWinkler : jellyfish/smellyfish", args{"jellyfish", "smellyfish", JaroWinkler}, 0.8962963, false},

		// Illegal argument error
		{"Undefined integer value for method", args{"abc", "abcde", 42}, 0.0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StringsSimilarity(tt.args.str1, tt.args.str2, tt.args.algo)
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
		algo    Algorithm
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"FuzzySearch 'testing'", args{"testnig", strList, Levenshtein}, "testing", false},
		{"FuzzySearch 'testing'", args{"test", strList, Levenshtein}, "test", false},
		{"FuzzySearch 'testing' err", args{"testnig", strList, Hamming}, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FuzzySearch(tt.args.str, tt.args.strList, tt.args.algo)
			if (err != nil) != tt.wantErr {
				t.Errorf("FuzzySearch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
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
		algo    Algorithm
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"FuzzySearch 'testing'", args{"testnig", strList, 0.7, Levenshtein}, "testing", false},
		{"FuzzySearch 'testing'", args{"test", strList, 0.7, Levenshtein}, "test", false},
		{"FuzzySearch 'testing'", args{"hello", strList, 0.7, Levenshtein}, "", false},
		{"FuzzySearch 'testing' err", args{"testing", strList, 0.7, Hamming}, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FuzzySearchThreshold(tt.args.str, tt.args.strList, tt.args.minSim, tt.args.algo)
			if (err != nil) != tt.wantErr {
				t.Errorf("FuzzySearchThreshold() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
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
		algo     Algorithm
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{"FuzzySearch 'testing'", args{"testnig", strList, 3, Levenshtein}, []string{"testing", "test", "tester"}, false},
		{"FuzzySearch 'testing' err", args{"testnig", strList, 3, Hamming}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FuzzySearchSet(tt.args.str, tt.args.strList, tt.args.quantity, tt.args.algo)
			if (err != nil) != tt.wantErr {
				t.Errorf("FuzzySearchSet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
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
		algo     Algorithm
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{"FuzzySearch 'testing'", args{"testnig", strList, 3, 0.7, Levenshtein}, []string{"testing", "", ""}, false},
		{"FuzzySearch 'testing' err", args{"testnig", strList, 3, 0.7, Hamming}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FuzzySearchSetThreshold(tt.args.str, tt.args.strList, tt.args.quantity, tt.args.minSim, tt.args.algo)
			if (err != nil) != tt.wantErr {
				t.Errorf("FuzzySearchSetThreshold() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FuzzySearchSetThreshold() = %v, want %v", got, tt.want)
			}
		})
	}
}
