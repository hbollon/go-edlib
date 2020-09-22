package hamming

import (
	"testing"
)

func TestHammingDistance(t *testing.T) {
	type args struct {
		str1 string
		str2 string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{"aa/aa", args{"aa", "aa"}, 0, false},
		{"ab/aa", args{"ab", "aa"}, 1, false},
		{"ab/ba", args{"ab", "ba"}, 2, false},
		{"ab/aaa", args{"ab", "aaa"}, 0, true},
		{"bbb/a", args{"bbb", "a"}, 0, true},
		{"🙂😄🙂😄/😄🙂😄🙂", args{"🙂😄🙂😄", "😄🙂😄🙂"}, 4, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Distance(tt.args.str1, tt.args.str2)
			if (err != nil) != tt.wantErr {
				t.Errorf("HammingDistance() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("HammingDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
