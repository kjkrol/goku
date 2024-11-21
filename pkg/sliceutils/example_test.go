package sliceutils

import (
	"fmt"
)

// ExampleIntersects demonstrates the usage of the Intersects function.
func ExampleIntersects() {
	slice1 := []int{1, 2, 3}
	slice2 := []int{3, 4, 5}
	result := Intersects(slice1, slice2)
	fmt.Println(result)
	// Output: true
}

// ExampleUnique demonstrates the usage of the Unique function.
func ExampleUnique() {
	slice := []int{1, 3, 3, 4}
	result := Unique(slice, func(x int) int { return x })
	fmt.Println(result)
	// Output: [1 3 4]
}

// ExampleMap demonstrates the usage of the Map function.
func ExampleMap() {
	slice := []int{1, 2, 3}
	result := Map(slice, func(x int) int { return x * 2 })
	fmt.Println(result)
	// Output: [2 4 6]
}

// ExamplePairs demonstrates the usage of the Pairs function.
func ExamplePairs() {
	slice := []int{1, 2, 3, 4, 5}
	sameParity := func(i1, i2 int) bool { return i1%2 == i2%2 }
	result := Pairs(slice, sameParity)
	fmt.Println(result)
	// Output: [{1 3} {1 5} {2 4} {3 5}]
}

// ExampleMinDistance demonstrates the usage of the MinDistance function.
func ExampleMinDistance() {
	slice1 := []int{1, 3, 3, 4}
	slice2 := []int{1, 3, 4}
	distance := func(a *int, b *int) int {
		res := *a - *b
		if res < 0 {
			return -res
		}
		return res
	}
	result := MinDistance(slice1, slice2, distance)
	fmt.Println(result)
	// Output: 0
}

// ExampleSameElements demonstrates the usage of the SameElements function.
func ExampleSameElements() {
	slice1 := []int{1, 2, 3}
	slice2 := []int{3, 2, 1}
	result := SameElements(slice1, slice2)
	fmt.Println(result)
	// Output: true
}
