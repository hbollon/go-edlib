package edlib

import (
	"testing"
)

func TestJaccardSimilarity(t *testing.T) {
	type args struct {
		str1 string
		str2 string
	}
	tests := []struct {
		name string
		args args
		want float32
	}{
		{"Jaccard sim 1", args{"Radiohead", "Carly Rae Jepsen"}, 0.0},
		{"Jaccard sim 2", args{"I love horror movies", "Lights out is a horror movie"}, 1.0 / 9.0},
		{"Jaccard sim 3", args{"love horror movies", "Lights out horror movie"}, 1.0 / 6.0},
		{"Jaccard sim 4", args{"私の名前はジョンです", "私の名前はジョン・ドゥです"}, 0.0},
		{"Jaccard sim 5", args{"🙂😄🙂😄 😄🙂😄", "🙂😄🙂😄 😄🙂😄 🙂😄🙂"}, 2.0 / 3.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := JaccardSimilarity(tt.args.str1, tt.args.str2, 0); got != tt.want {
				t.Errorf("JaccardSimilarity() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJaccardShingleSimilarity(t *testing.T) {
	type args struct {
		str1 string
		str2 string
	}
	tests := []struct {
		name string
		args args
		want float32
	}{
		{"Jaccard shingle sim 1", args{"Radiohead", "Carly Rae Jepsen"}, 0.04761905},
		{"Jaccard shingle sim 2", args{"I love horror movies", "Lights out is a horror movie"}, 0.3548387},
		{"Jaccard shingle sim 3", args{"love horror movies", "Lights out horror movie"}, 0.44},
		{"Jaccard shingle sim 4", args{"私の名前はジョンです", "私の名前はジョン・ドゥです"}, 0.61538464},
		{"Jaccard shingle sim 5", args{"🙂😄🙂😄 😄🙂😄", "🙂😄🙂😄 😄🙂😄 🙂😄🙂"}, 0.8},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := JaccardSimilarity(tt.args.str1, tt.args.str2, 2); got != tt.want {
				t.Errorf("JaccardSimilarity() with shingle 2 = %v, want %v", got, tt.want)
			}
		})
	}
}
