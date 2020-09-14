package orderedmap

import (
	"reflect"
	"testing"
)

var testOrderedMap OrderedMap

func init() {
	testOrderedMap = OrderedMap{
		{"test1", 5},
		{"test2", 2},
		{"test3", 4},
		{"test4", 3},
		{"test5", 6},
	}
}

func TestOrderedMap_Len(t *testing.T) {
	tests := []struct {
		name string
		p    OrderedMap
		want int
	}{
		{"OrderedMap Len", testOrderedMap, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.Len(); got != tt.want {
				t.Errorf("OrderedMap.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrderedMap_Less(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		p    OrderedMap
		args args
		want bool
	}{
		{"OrderedMap Less false", testOrderedMap, args{0, 1}, false},
		{"OrderedMap Less true", testOrderedMap, args{1, 2}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.Less(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("OrderedMap.Less() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrderedMap_Swap(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		p    OrderedMap
		args args
	}{
		{"OrderedMap Swap", testOrderedMap, args{0, testOrderedMap.Len() - 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			firstElem := testOrderedMap[0]
			lastElem := testOrderedMap[testOrderedMap.Len()-1]
			tt.p.Swap(tt.args.i, tt.args.j)
			if testOrderedMap[0].Key != lastElem.Key && testOrderedMap[0].Value != lastElem.Value &&
				testOrderedMap[testOrderedMap.Len()-1].Key != firstElem.Key && testOrderedMap[testOrderedMap.Len()-1].Value != firstElem.Value {
				t.Errorf("Swap between first and last OrderedMap element failed")
			}
		})
	}
}

func TestOrderedMap_ToArray(t *testing.T) {
	tests := []struct {
		name string
		p    OrderedMap
		want []string
	}{
		{"OrderedMap ToArray", testOrderedMap, []string{"test5", "test2", "test3", "test4", "test1"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.ToArray(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderedMap.ToArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrderedMap_SortByValues(t *testing.T) {
	tests := []struct {
		name string
		p    OrderedMap
	}{
		{"OrderedMap SortByValues", testOrderedMap},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testOrderedMap.SortByValues()
			for i := 1; i < testOrderedMap.Len(); i++ {
				if !testOrderedMap.Less(i, i-1) {
					t.Errorf("SortByValue failed, OrderedMap isn't sorted")
				}
			}
		})
	}
}
