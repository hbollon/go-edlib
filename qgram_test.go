package edlib

import (
	"testing"
)

func TestQgramDistance(t *testing.T) {
	type args struct {
		str1        string
		str2        string
		splitLength int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Qgram sim 1", args{"Radiohead", "Radiohead", 2}, 0.0},
		{"Qgram sim 2", args{"ABCD", "ABCE", 2}, 2.0},
		{"Qgram sim 3", args{"Radiohead", "Carly Rae Jepsen", 2}, 21.0},
		{"Qgram sim 4", args{"I love horror movies", "Lights out is a horror movie", 2}, 22.0},
		{"Qgram sim 5", args{"love horror movies", "Lights out horror movie", 2}, 15.0},
		{"Qgram sim 6", args{"私の名前はジョンです", "私の名前はジョン・ドゥです", 2}, 5},
		{"Qgram sim 7", args{"🙂😄🙂😄 😄🙂😄", "🙂😄🙂😄 😄🙂😄 🙂😄🙂", 2}, 4},
		{"Qgram sim 8", args{"", "", 2}, 0.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := QgramDistance(tt.args.str1, tt.args.str2, tt.args.splitLength); got != tt.want {
				t.Errorf("QgramDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQgramSimilarity(t *testing.T) {
	type args struct {
		str1        string
		str2        string
		splitLength int
	}
	tests := []struct {
		name string
		args args
		want float32
	}{
		{"identical", args{"abcde", "abcde", 2}, 1.0},
		{"no match", args{"abcd", "effgghh", 2}, 0.0},
		// Strings shorter than the shingle length produce no q-grams; identical
		// inputs must still be fully similar, not NaN.
		{"identical shorter than k", args{"a", "a", 2}, 1.0},
		{"both empty", args{"", "", 2}, 1.0},
		{"identical equal to shorter than k", args{"ab", "ab", 3}, 1.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := QgramSimilarity(tt.args.str1, tt.args.str2, tt.args.splitLength); got != tt.want {
				t.Errorf("QgramSimilarity() = %v, want %v", got, tt.want)
			}
		})
	}
}
