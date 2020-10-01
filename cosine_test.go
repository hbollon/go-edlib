package edlib

import (
	"reflect"
	"testing"
)

func TestCosineSimilarity(t *testing.T) {
	type args struct {
		str1 string
		str2 string
	}
	tests := []struct {
		name string
		args args
		want float32
	}{
		{"Cosine sim 1", args{"Radiohead", "Carly Rae Jepsen"}, 0.0},
		{"Cosine sim 2", args{"I love horror movies", "Lights out is a horror movie"}, 0.20412414},
		{"Cosine sim 3", args{"love horror movies", "Lights out horror movie"}, 0.28867513},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CosineSimilarity(tt.args.str1, tt.args.str2); got != tt.want {
				t.Errorf("CosineSimilarity() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_union(t *testing.T) {
	type args struct {
		a []string
		b []string
	}
	tests := []struct {
		name string
		args args
		want [][]rune
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := union(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("union() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_find(t *testing.T) {
	type args struct {
		slice [][]rune
		val   []rune
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := find(tt.args.slice, tt.args.val)
			if got != tt.want {
				t.Errorf("find() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("find() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_sum(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sum(tt.args.arr); got != tt.want {
				t.Errorf("sum() = %v, want %v", got, tt.want)
			}
		})
	}
}
