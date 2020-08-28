package edlib

import "testing"

func TestJaroSimilarity(t *testing.T) {
	type args struct {
		str1 string
		str2 string
	}
	tests := []struct {
		name string
		args args
		want float32
	}{
		{"First arg empty", args{"", "abcde"}, 0.0},
		{"Second arg empty", args{"abcde", ""}, 0.0},
		{"Same args", args{"abcde", "abcde"}, 1.0},
		{"CRATE/TRACE", args{"CRATE", "TRACE"}, 0.73333335},
		{"MARTHA/MARHTA", args{"MARTHA", "MARHTA"}, 0.9444444},
		{"DIXON/DICKSONX", args{"DIXON", "DICKSONX"}, 0.76666665},
		{"jellyfish/smellyfish", args{"jellyfish", "smellyfish"}, 0.8962963},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := JaroSimilarity(tt.args.str1, tt.args.str2); got != tt.want {
				t.Errorf("JaroSimilarity() = %v, want %v", got, tt.want)
			}
		})
	}
}
