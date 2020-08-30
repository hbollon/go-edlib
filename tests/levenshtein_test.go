package edlib

import (
	"testing"

	"github.com/hbollon/go-edlib"
)

func Test_levenshteinDistance(t *testing.T) {
	type args struct {
		str1 string
		str2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"First arg empty", args{"", "abcde"}, 5},
		{"Second arg empty", args{"abcde", ""}, 5},
		{"Same args", args{"abcde", "abcde"}, 0},
		{"ab/aa", args{"ab", "aa"}, 1},
		{"ab/ba", args{"ab", "ba"}, 2},
		{"ab/aaa", args{"ab", "aaa"}, 2},
		{"bbb/a", args{"bbb", "a"}, 3},
		{"kitten/sitting", args{"kitten", "sitting"}, 3},
		{"distance/difference", args{"distance", "difference"}, 5},
		{"a cat/an abct", args{"a cat", "an abct"}, 4},
		{"こにんち/こんにちは", args{"こにんち", "こんにちは"}, 3}, // "Hello" in Japanese
		{"🙂😄🙂😄/😄🙂😄🙂", args{"🙂😄🙂😄", "😄🙂😄🙂"}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := edlib.LevenshteinDistance(tt.args.str1, tt.args.str2); got != tt.want {
				t.Errorf("levenshteinDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOSADamerauLevenshteinDistance(t *testing.T) {
	type args struct {
		str1 string
		str2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"First arg empty", args{"", "abcde"}, 5},
		{"Second arg empty", args{"abcde", ""}, 5},
		{"Same args", args{"abcde", "abcde"}, 0},
		{"ab/aa", args{"ab", "aa"}, 1},
		{"ab/ba", args{"ab", "ba"}, 1},
		{"ab/aaa", args{"ab", "aaa"}, 2},
		{"bbb/a", args{"bbb", "a"}, 3},
		{"ca/abc", args{"ca", "abc"}, 3},
		{"a cat/an abct", args{"a cat", "an abct"}, 4},
		{"dixon/dicksonx", args{"dixon", "dicksonx"}, 4},
		{"jellyfish/smellyfish", args{"jellyfish", "smellyfish"}, 2},
		{"こにんち/こんにちは", args{"こにんち", "こんにちは"}, 2}, // "Hello" in Japanese
		{"🙂😄🙂😄/😄🙂😄🙂", args{"🙂😄🙂😄", "😄🙂😄🙂"}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := edlib.OSADamerauLevenshteinDistance(tt.args.str1, tt.args.str2); got != tt.want {
				t.Errorf("OSADamerauLevenshteinDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDamerauLevenshteinDistance(t *testing.T) {
	type args struct {
		str1 string
		str2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"First arg empty", args{"", "abcde"}, 5},
		{"Second arg empty", args{"abcde", ""}, 5},
		{"Same args", args{"abcde", "abcde"}, 0},
		{"ab/aa", args{"ab", "aa"}, 1},
		{"ab/ba", args{"ab", "ba"}, 1},
		{"ab/aaa", args{"ab", "aaa"}, 2},
		{"bbb/a", args{"bbb", "a"}, 3},
		{"ca/abc", args{"ca", "abc"}, 2},
		{"a cat/an abct", args{"a cat", "an abct"}, 3},
		{"dixon/dicksonx", args{"dixon", "dicksonx"}, 4},
		{"jellyfish/smellyfish", args{"jellyfish", "smellyfish"}, 2},
		{"こにんち/こんにちは", args{"こにんち", "こんにちは"}, 2}, // "Hello" in Japanese
		{"🙂😄🙂😄/😄🙂😄🙂", args{"🙂😄🙂😄", "😄🙂😄🙂"}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := edlib.DamerauLevenshteinDistance(tt.args.str1, tt.args.str2); got != tt.want {
				t.Errorf("DamerauLevenshteinDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
