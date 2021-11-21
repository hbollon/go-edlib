package edlib

import (
	"reflect"
	"testing"
)

var testArr1, testArr2 []string

func init() {
	testArr1 = []string{
		"a",
		"b",
		"d",
	}

	testArr2 = []string{
		"a",
		"e",
	}
}

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
			if got := CosineSimilarity(tt.args.str1, tt.args.str2, 0); got != tt.want {
				t.Errorf("CosineSimilarity() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCosineShingleSimilarity(t *testing.T) {
	type args struct {
		str1 string
		str2 string
	}
	tests := []struct {
		name string
		args args
		want float32
	}{
		{"Cosine shingle sim 1", args{"Radiohead", "Carly Rae Jepsen"}, 0.09759001},
		{"Cosine shingle sim 2", args{"I love horror movies", "Lights out is a horror movie"}, 0.5335784},
		{"Cosine shingle sim 3", args{"love horror movies", "Lights out horror movie"}, 0.61977977},
		{"Cosine shingle sim 4", args{"私の名前はジョンです", "私の名前はジョン・ドゥです"}, 0.76980036},
		{"Cosine shingle sim 5", args{"🙂😄🙂😄 😄🙂😄", "🙂😄🙂😄 😄🙂😄 🙂😄🙂"}, 0.8944272},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CosineSimilarity(tt.args.str1, tt.args.str2, 2); got != tt.want {
				t.Errorf("CosineSimilarity() with shingle 2 = %v, want %v", got, tt.want)
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
		{"Union function test", args{testArr1, testArr2}, [][]rune{{97}, {98}, {100}, {101}}},
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
		name string
		args args
		want int
	}{
		{"Find function test true", args{[][]rune{{97}, {98}, {100}, {101}}, []rune{101}}, 3},
		{"Find function test false", args{[][]rune{{97}, {98}, {100}, {101}}, []rune{102}}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := find(tt.args.slice, tt.args.val)
			if got != tt.want {
				t.Errorf("find() got = %v, want %v", got, tt.want)
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
		{"Sum function test", args{[]int{10, 40, 5, 2, 20}}, 77},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sum(tt.args.arr); got != tt.want {
				t.Errorf("sum() = %v, want %v", got, tt.want)
			}
		})
	}
}
