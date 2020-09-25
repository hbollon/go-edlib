package utils

import (
	"testing"
)

var hashMapA, hashMapB StringHashMap

func init() {
	hashMapA = StringHashMap{
		"a": struct{}{},
		"b": struct{}{},
		"c": struct{}{},
	}

	hashMapB = StringHashMap{
		"d": struct{}{},
		"e": struct{}{},
		"f": struct{}{},
	}
}

func TestMin(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Min between 2/4", args{2, 4}, 2},
		{"Min between -25/-42", args{-25, -42}, -42},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Min(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Min() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMax(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Min between 2/4", args{2, 4}, 4},
		{"Min between -25/-42", args{-25, -42}, -25},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Max(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Max() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEqual(t *testing.T) {
	type args struct {
		a []rune
		b []rune
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Equal between toto/test", args{[]rune("toto"), []rune("test")}, false},
		{"Equal between toto/toto", args{[]rune("toto"), []rune("toto")}, true},
		{"Equal between Toto/toto", args{[]rune("Toto"), []rune("toto")}, false},
		{"Equal between 🙂😄/🙂😄", args{[]rune("🙂😄"), []rune("🙂😄")}, true},
		{"Equal between 🙂😄/🙂😄🙂😄", args{[]rune("🙂😄"), []rune("🙂😄🙂😄")}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Equal(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Equal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringHashMap_AddAll(t *testing.T) {
	type args struct {
		srcMap StringHashMap
	}
	tests := []struct {
		name string
		m    StringHashMap
		args args
	}{
		{"AddAll between hashMapA/hashMapB", hashMapA, args{hashMapB}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			oldLen := len(tt.m)
			tt.m.AddAll(tt.args.srcMap)
			if len(tt.m) != oldLen+len(tt.args.srcMap) {
				t.Errorf("AddAll() failed for test case : \"%s\"", tt.name)
			}
		})
	}
}

func TestStringHashMap_ToArray(t *testing.T) {
	tests := []struct {
		name string
		m    StringHashMap
	}{
		{"ToString() hashMapA", hashMapA},
		{"ToString() hashMapB", hashMapB},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.m.ToArray()
			if len(got) != len(tt.m) {
				t.Errorf("ToArray() failed for test case : \"%s\"", tt.name)
			}
		})
	}
}
