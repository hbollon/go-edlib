package edlib

import (
	"testing"
)

func TestSorensenDiceCoefficient(t *testing.T) {
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
		{"SorensenDiceCoefficient 1", args{"night", "nacht", 2}, 0.25},
		{"SorensenDiceCoefficient 2", args{"Radiohead", "Radiohead", 2}, 1.0},
		{"SorensenDiceCoefficient 3", args{"", "", 2}, 0.0},
		{"SorensenDiceCoefficient 4", args{"Radiohead", "Carly Rae Jepsen", 2}, 0.09090909},
		{"SorensenDiceCoefficient 5", args{"I love horror movies", "Lights out is a horror movie", 2}, 0.52380955},
		{"SorensenDiceCoefficient 6", args{"love horror movies", "Lights out horror movie", 2}, 0.6111111},
		{"SorensenDiceCoefficient 7", args{"私の名前はジョンです", "私の名前はジョン・ドゥです", 2}, 0.7619048},
		{"SorensenDiceCoefficient 8", args{"🙂😄🙂😄 😄🙂😄", "🙂😄🙂😄 😄🙂😄 🙂😄🙂", 2}, 0.8888889},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SorensenDiceCoefficient(tt.args.str1, tt.args.str2, tt.args.splitLength); got != tt.want {
				t.Errorf("SorensenDiceCoefficient() = %v, want %v", got, tt.want)
			}
		})
	}
}
