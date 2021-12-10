package edlib

import (
	"testing"
)

func TestSorensenDiceSimilarity(t *testing.T) {
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
		{"SorensenDiceSimilarity 1", args{"night", "nacht", 2}, 0.25},
		{"SorensenDiceSimilarity 2", args{"Radiohead", "Radiohead", 2}, 1.0},
		{"SorensenDiceSimilarity 3", args{"", "", 2}, 0.0},
		{"SorensenDiceSimilarity 4", args{"Radiohead", "Carly Rae Jepsen", 2}, 0.09090909},
		{"SorensenDiceSimilarity 5", args{"I love horror movies", "Lights out is a horror movie", 2}, 0.52380955},
		{"SorensenDiceSimilarity 6", args{"love horror movies", "Lights out horror movie", 2}, 0.6111111},
		{"SorensenDiceSimilarity 7", args{"私の名前はジョンです", "私の名前はジョン・ドゥです", 2}, 0.7619048},
		{"SorensenDiceSimilarity 8", args{"🙂😄🙂😄 😄🙂😄", "🙂😄🙂😄 😄🙂😄 🙂😄🙂", 2}, 0.8888889},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SorensenDiceSimilarity(tt.args.str1, tt.args.str2, tt.args.splitLength); got != tt.want {
				t.Errorf("SorensenDiceSimilarity() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSorensenDiceDistance(t *testing.T) {
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
		{"SorensenDiceSimilarity 1", args{"night", "nacht", 2}, 0.75},
		{"SorensenDiceSimilarity 2", args{"Radiohead", "Radiohead", 2}, 0},
		{"SorensenDiceSimilarity 3", args{"", "", 2}, 1},
		{"SorensenDiceSimilarity 4", args{"Radiohead", "Carly Rae Jepsen", 2}, 0.9090909},
		{"SorensenDiceSimilarity 5", args{"I love horror movies", "Lights out is a horror movie", 2}, 0.47619045},
		{"SorensenDiceSimilarity 6", args{"love horror movies", "Lights out horror movie", 2}, 0.3888889},
		{"SorensenDiceSimilarity 7", args{"私の名前はジョンです", "私の名前はジョン・ドゥです", 2}, 0.23809522},
		{"SorensenDiceSimilarity 8", args{"🙂😄🙂😄 😄🙂😄", "🙂😄🙂😄 😄🙂😄 🙂😄🙂", 2}, 0.111111104},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SorensenDiceDistance(tt.args.str1, tt.args.str2, tt.args.splitLength); got != tt.want {
				t.Errorf("SorensenDiceDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
