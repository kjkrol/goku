package sliceutils

import (
	"slices"
	"testing"
)

func TestUnique(t *testing.T) {
	for _, test := range []struct {
		input    []int
		expected []int
	}{
		{[]int{1, 3, 3, 4}, []int{1, 3, 4}},
		{[]int{1, 2, 3, 4}, []int{1, 2, 3, 4}},
		{[]int{1, 1, 1, 4}, []int{1, 4}},
	} {
		if output := Unique(test.input, func(x int) int { return x }); len(output) != len(test.expected) {
			t.Errorf("output %v not equal to expected %v", output, test.expected)
		}
	}
}

func TestIntersects(t *testing.T) {
	for _, test := range []struct {
		arg1     []int
		arg2     []int
		expected bool
	}{
		{[]int{1, 3, 3, 4}, []int{1, 3, 4}, true},
		{[]int{1, 2, 3, 4}, []int{1, 2, 3, 4}, true},
		{[]int{1, 1, 1, 4}, []int{2}, false},
	} {
		if output := Intersects(test.arg1, test.arg2); output != test.expected {
			t.Errorf("output %v not equal to expected %v", output, test.expected)
		}
	}
}

func TestPairs(t *testing.T) {
	sampeParity := func(i1, i2 int) bool { return i1%2 == i2%2 }
	for _, test := range []struct {
		arg1     []int
		expected []Pair[int]
	}{
		{[]int{}, []Pair[int]{}},
		{[]int{1}, []Pair[int]{}},
		{[]int{1, 2, 3, 4, 5}, []Pair[int]{{1, 3}, {1, 5}, {2, 4}, {3, 5}}},
	} {
		if output := Pairs(test.arg1, sampeParity); !slices.Equal(output, test.expected) {
			t.Errorf("output %v not equal to expected %v", output, test.expected)
		}
	}
}

func TestMinDistance(t *testing.T) {
	distance := func(a *int, b *int) int {
		res := *a - *b
		if res < 0 {
			return -res
		}
		return res
	}
	for _, test := range []struct {
		arg1     []int
		arg2     []int
		expected int
	}{
		{[]int{1, 3, 3, 4}, []int{1, 3, 4}, 0},
		{[]int{1, 2, 3, 4}, []int{10, 12, 13, 14}, 6},
		{[]int{1, 1, 1, 4}, []int{8}, 4},
	} {
		if output := MinDistance(test.arg1, test.arg2, distance); output != test.expected {
			t.Errorf("output %v not equal to expected %v", output, test.expected)
		}
	}
}
