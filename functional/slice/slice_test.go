package slice

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMapSameLength(t *testing.T) {
	addOne := func(i int) int { return i + 1 }

	s1 := []int{1, 2, 3, 4, 5}
	if len(s1) != len(Map(s1, addOne)) {
		t.Fatalf("Map should result in equal-length slice")
	}

	var s2 []int
	if len(s2) != len(Map(s2, addOne)) {
		t.Fatalf("Map should result in equal-length slice")
	}
}

func TestMap(t *testing.T) {
	// test simple update
	update := func(s string) string { return s + "-updated" }
	s1 := []string{"one", "two"}
	s1P := Map(s1, update)
	if s1P[0] != "one-updated" && s1P[1] != "two-updated" {
		t.Fatalf("unexpected values in returned map array: %v", s1P)
	}

	// test type-change with multi-dimensional arrays
	multiDim := func(i int) []string { return []string{fmt.Sprintf("%d", i), fmt.Sprintf("%d", i)} }
	s2 := []int{1, 2, 3, 4}
	s2P := Map(s2, multiDim)
	expected := [][]string{{"1", "1"}, {"2", "2"}, {"3", "3"}, {"4", "4"}}
	if !reflect.DeepEqual(s2P, expected) {
		t.Fatalf("expected %v, got %v", expected, s2P)
	}
}

func TestFlatMap(t *testing.T) {
	duplicate := func(i int) []int { return []int{i, i} }
	noop := func(i int) []int { return nil }

	for _, test := range []struct {
		input  []int
		output []int
		f      func(int) []int
	}{
		{
			input:  []int{1, 2, 3, 4},
			output: []int{1, 1, 2, 2, 3, 3, 4, 4},
			f:      duplicate,
		},
		{
			input:  []int{1},
			output: []int{1, 1},
			f:      duplicate,
		},
		{
			input:  []int{},
			output: []int{},
			f:      duplicate,
		},
		{
			input:  nil,
			output: []int{},
			f:      duplicate,
		},
		{
			input:  nil,
			output: []int{},
			f:      noop,
		},
		{
			input:  []int{1, 2, 3, 4},
			output: []int{},
			f:      noop,
		},
	} {
		t.Run("", func(t *testing.T) {
			out := FlatMap(test.input, test.f)
			if !reflect.DeepEqual(out, test.output) {
				t.Fatalf("expected %v, got %v", test.output, out)
			}
		})
	}
}
