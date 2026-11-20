package edlib

import (
	"testing"
)

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
		{"No characters match", args{"abcd", "effgghh"}, 0.0},
		{"CRATE/TRACE", args{"CRATE", "TRACE"}, 0.73333335},
		{"MARTHA/MARHTA", args{"MARTHA", "MARHTA"}, 0.9444444},
		{"DIXON/DICKSONX", args{"DIXON", "DICKSONX"}, 0.76666665},
		{"jellyfish/smellyfish", args{"jellyfish", "smellyfish"}, 0.8962963},
		// 3 out of order matches = 1 transposition (rounded down), not 1.5
		{"parish/liar", args{"parish", "liar"}, 0.6388889},
		{"abcb/ceabcf", args{"abcb", "ceabcf"}, 0.6388889},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := JaroSimilarity(tt.args.str1, tt.args.str2); got != tt.want {
				t.Errorf("JaroSimilarity() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJaroWinklerSimilarity(t *testing.T) {
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
		{"No characters match", args{"abcd", "effgghh"}, 0.0},
		{"TRACE/TRACE", args{"TRACE", "TRACE"}, 1.0},
		{"CRATE/TRACE", args{"CRATE", "TRACE"}, 0.73333335},
		{"TRATE/TRACE", args{"TRATE", "TRACE"}, 0.90666664},
		{"DIXON/DICKSONX", args{"DIXON", "DICKSONX"}, 0.81333333},
		{"parish/liar", args{"parish", "liar"}, 0.6388889},
		// No prefix boost when the Jaro similarity is not above 0.7
		{"less/lede", args{"less", "lede"}, 0.6666667},
		{"cavy/camb", args{"cavy", "camb"}, 0.6666667},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := JaroWinklerSimilarity(tt.args.str1, tt.args.str2); got != tt.want {
				t.Errorf("JaroWinklerSimilarity() = %v, want %v", got, tt.want)
			}
		})
	}
}
