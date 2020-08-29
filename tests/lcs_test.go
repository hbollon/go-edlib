package edlib

import (
	"testing"

	"github.com/hbollon/go-edlib"
)

func TestLCS(t *testing.T) {
	type args struct {
		str1 string
		str2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"AB/empty", args{"AB", ""}, 0},
		{"empty/AB", args{"", "AB"}, 0},
		{"AB/AB", args{"AB", "AB"}, 2},
		{"ABCD/ACBAD", args{"ABCD", "ACBAD"}, 3},
		{"ABCDGH/AEDFHR", args{"ABCDGH", "AEDFHR"}, 3},
		{"AGGTAB/GXTXAYB", args{"AGGTAB", "GXTXAYB"}, 4},
		{"XMJYAUZ/MZJAWXU", args{"XMJYAUZ", "MZJAWXU"}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := edlib.LCS(tt.args.str1, tt.args.str2); got != tt.want {
				t.Errorf("LCS() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLCSEditDistance(t *testing.T) {
	type args struct {
		str1 string
		str2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"AB/empty", args{"AB", ""}, 2},
		{"empty/AB", args{"", "AB"}, 2},
		{"AB/AB", args{"AB", "AB"}, 0},
		{"CAT/CUT", args{"CAT", "CUT"}, 2},
		{"ACB/AB", args{"ACB", "AB"}, 1},
		{"ABC/ACD", args{"ABC", "ACD"}, 2},
		{"ABCD/ACBAD", args{"ABCD", "ACBAD"}, 3},
		{"ABCDGH/AEDFHR", args{"ABCDGH", "AEDFHR"}, 6},
		{"AGGTAB/GXTXAYB", args{"AGGTAB", "GXTXAYB"}, 5},
		{"XMJYAUZ/MZJAWXU", args{"XMJYAUZ", "MZJAWXU"}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := edlib.LCSEditDistance(tt.args.str1, tt.args.str2); got != tt.want {
				t.Errorf("LCSEditDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
