package edlib

import (
	"reflect"
	"testing"
)

func TestShingle(t *testing.T) {
	type args struct {
		str string
		k   int
	}
	tests := []struct {
		name string
		args args
		want map[string]int
	}{
		{"shingle 1", args{"Radiohead", 2}, map[string]int{"Ra": 1, "ad": 2, "di": 1, "ea": 1, "he": 1, "io": 1, "oh": 1}},
		{"shingle 1-1", args{"Radiohead", 3}, map[string]int{"Rad": 1, "adi": 1, "dio": 1, "ead": 1, "hea": 1, "ioh": 1, "ohe": 1}},
		{"shingle 2", args{"I love horror movies", 2}, map[string]int{" h": 1, " l": 1, " m": 1, "I ": 1, "e ": 1, "es": 1, "ho": 1, "ie": 1, "lo": 1, "mo": 1, "or": 2, "ov": 2, "r ": 1, "ro": 1, "rr": 1, "ve": 1, "vi": 1}},
		{"shingle 3", args{"私の名前はジョンです", 2}, map[string]int{"です": 1, "の名": 1, "はジ": 1, "ジョ": 1, "ョン": 1, "ンで": 1, "前は": 1, "名前": 1, "私の": 1}},
		{"shingle 4", args{"🙂😄🙂😄 😄🙂😄", 2}, map[string]int{" 😄": 1, "😄 ": 1, "😄🙂": 2, "🙂😄": 3}},
		{"shingle 5", args{"", 100}, make(map[string]int)},
		{"shingle 6", args{"hello", 0}, make(map[string]int)},
		{"shingle 7", args{"四畳半神話大系", 7}, map[string]int{"四畳半神話大系": 1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Shingle(tt.args.str, tt.args.k)
			eq := reflect.DeepEqual(got, tt.want)
			if !eq {
				t.Errorf("Shingle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestShingleSlidingWindow(t *testing.T) {
	type args struct {
		str string
		k   int
	}
	tests := []struct {
		name string
		args args
		want map[string]int
	}{
		{"shingle with window 1", args{"Radiohead", 2}, map[string]int{"Ra": 1, "ad": 2, "di": 1, "ea": 1, "he": 1, "io": 1, "oh": 1}},
		{"shingle with window 1-1", args{"Radiohead", 3}, map[string]int{"Rad": 1, "adi": 1, "dio": 1, "ead": 1, "hea": 1, "ioh": 1, "ohe": 1}},
		{"shingle with window 2", args{"I love horror movies", 2}, map[string]int{" h": 1, " l": 1, " m": 1, "I ": 1, "e ": 1, "es": 1, "ho": 1, "ie": 1, "lo": 1, "mo": 1, "or": 2, "ov": 2, "r ": 1, "ro": 1, "rr": 1, "ve": 1, "vi": 1}},
		{"shingle with window 3", args{"私の名前はジョンです", 2}, map[string]int{"です": 1, "の名": 1, "はジ": 1, "ジョ": 1, "ョン": 1, "ンで": 1, "前は": 1, "名前": 1, "私の": 1}},
		{"shingle with window 4", args{"🙂😄🙂😄 😄🙂😄", 2}, map[string]int{" 😄": 1, "😄 ": 1, "😄🙂": 2, "🙂😄": 3}},
		{"shingle with window 5", args{"", 100}, make(map[string]int)},
		{"shingle with window 6", args{"hello", 0}, make(map[string]int)},
		{"shingle with window 7", args{"四畳半神話大系", 7}, map[string]int{"四畳半神話大系": 1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ShingleSlidingWindow(tt.args.str, tt.args.k)
			eq := reflect.DeepEqual(got, tt.want)
			if !eq {
				t.Errorf("ShingleSlidingWindow() = %v, want %v", got, tt.want)
			}
		})
	}

}

func benchmarkShinglingFunctions(f func(string, int) map[string]int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		f("Radiohead", 2)
		f("asdfl;ja;sdjva;ljdsf;lakjdsf;l私の名前はジョンです私の名前はジョンですjasdvjal;sdjf;lasjdv;l私の名前はジョンですajsdfj;lv", 15)
		f("asdfl;ja;sdjva;ljdsf;lakjdsf;l私の名前はジョンです私の名前はジョンですjasdvjal;sdjf;lasjdv;l私の名前はジョンですajsdfj;lv", 5)
	}
}

func BenchmarkShingle(b *testing.B) {
	benchmarkShinglingFunctions(Shingle, b)
}

func BenchmarkShingleSlidingWindow(b *testing.B) {
	benchmarkShinglingFunctions(ShingleSlidingWindow, b)
}
