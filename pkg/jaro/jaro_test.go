package jaro

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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Similarity(tt.args.str1, tt.args.str2); got != tt.want {
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WinklerSimilarity(tt.args.str1, tt.args.str2); got != tt.want {
				t.Errorf("JaroWinklerSimilarity() = %v, want %v", got, tt.want)
			}
		})
	}
}
