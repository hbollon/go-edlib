package edlib

import (
	"testing"
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LevenshteinDistance(tt.args.str1, tt.args.str2); got != tt.want {
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := OSADamerauLevenshteinDistance(tt.args.str1, tt.args.str2); got != tt.want {
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DamerauLevenshteinDistance(tt.args.str1, tt.args.str2); got != tt.want {
				t.Errorf("DamerauLevenshteinDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
