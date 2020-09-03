package edlib

import (
	"reflect"
	"sort"
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

func TestLCSBacktrack(t *testing.T) {
	type args struct {
		str1 string
		str2 string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"AB/empty", args{"AB", ""}, "", true},
		{"empty/AB", args{"", "AB"}, "", true},
		{"AB/AB", args{"AB", "AB"}, "AB", false},
		{"ABCD/ACBAD", args{"ABCD", "ACBAD"}, "ABD", false},
		{"ABCDGH/AEDFHR", args{"ABCDGH", "AEDFHR"}, "ADH", false},
		{"AGGTAB/GXTXAYB", args{"AGGTAB", "GXTXAYB"}, "GTAB", false},
		{"XMJYAUZ/MZJAWXU", args{"XMJYAUZ", "MZJAWXU"}, "MJAU", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := edlib.LCSBacktrack(tt.args.str1, tt.args.str2)
			if (err != nil) != tt.wantErr {
				t.Errorf("LCSBacktrack() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("LCSBacktrack() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLCSBacktrackAll(t *testing.T) {
	type args struct {
		str1 string
		str2 string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{"AB/empty", args{"AB", ""}, nil, true},
		{"empty/AB", args{"", "AB"}, nil, true},
		{"AB/AB", args{"AB", "AB"}, []string{"AB"}, false},
		{"ABCD/ACBAD", args{"ABCD", "ACBAD"}, []string{"ABD", "ACD"}, false},
		{"ABCDGH/AEDFHR", args{"ABCDGH", "AEDFHR"}, []string{"ADH"}, false},
		{"AGGTAB/GXTXAYB", args{"AGGTAB", "GXTXAYB"}, []string{"GTAB"}, false},
		{"XMJYAUZ/MZJAWXU", args{"XMJYAUZ", "MZJAWXU"}, []string{"MJAU"}, false},
		{"AZBYCWDX/ZAYBWCXD", args{"AZBYCWDX", "ZAYBWCXD"}, []string{"ABCD", "ABCX", "ABWD", "ABWX", "AYCD", "AYCX", "AYWD", "AYWX", "ZBCD", "ZBCX", "ZBWD", "ZBWX", "ZYCD", "ZYCX", "ZYWD", "ZYWX"}, false},
		{"AATCC/ACACG", args{"AATCC", "ACACG"}, []string{"AAC", "ACC"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := edlib.LCSBacktrackAll(tt.args.str1, tt.args.str2)
			if (err != nil) != tt.wantErr {
				t.Errorf("LCSBacktrackAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			sort.Strings(got)
			sort.Strings(tt.want)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LCSBacktrackAll() = %v, want %v", got, tt.want)
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
		{"No characters match", args{"abcd", "effgghh"}, 11},
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
