package sliceutils

import (
	"math"
)

// Intersects checks if there is at least one common element between two slices.
// It returns true if there is a common element, otherwise false.
//
// The function takes two slices of any comparable type as input parameters.
//
// Example usage:
//
//	slice1 := []int{1, 2, 3}
//	slice2 := []int{3, 4, 5}
//	result := Intersects(slice1, slice2) // result will be true
//
// Type Parameters:
//
//	T: The type of elements in the slices, which must be comparable.
//
// Parameters:
//
//	slice1: The first slice to be compared.
//	slice2: The second slice to be compared.
//
// Returns:
//
//	bool: true if there is at least one common element between the two slices, otherwise false.
func Intersects[T comparable](slice1 []T, slice2 []T) bool {
	for _, x := range slice1 {
		for _, y := range slice2 {
			if x == y {
				return true
			}
		}
	}
	return false
}

// Unique returns a new slice containing only the unique elements from the input slice.
// The uniqueness of each element is determined by the value returned by the key function.
//
// Type Parameters:
//   - T: The type of elements in the input slice.
//   - V: The type of the value returned by the key function, which must be comparable.
//
// Parameters:
//   - slice: The input slice from which to extract unique elements.
//   - key: A function that takes an element of type T and returns a value of type V,
//     which is used to determine the uniqueness of the element.
//
// Returns:
//
//	A new slice containing only the unique elements from the input slice.
func Unique[T any, V comparable](slice []T, key func(T) V) []T {
	memo := make(map[V]bool)
	var result []T
	for _, v := range slice {
		if _, ok := memo[key(v)]; !ok {
			memo[key(v)] = true
			result = append(result, v)
		}
	}
	return result
}

// Pair is a generic struct that holds a pair of values of the same type.
// The type of the values is specified by the type parameter T.
type Pair[T any] struct {
	First  T
	Second T
}

// Map applies a transformation function to each element of a slice and returns a new slice
// containing the transformed elements.
//
// Parameters:
//   - slice: The input slice of type T.
//   - transform: A function that takes an element of type T and returns a value of type R.
//
// Returns:
//
//	A new slice of type R containing the transformed elements.
func Map[T any, R any](slice []T, transform func(T) R) []R {
	result := make([]R, len(slice))
	for i, v := range slice {
		result[i] = transform(v)
	}
	return result
}

// Pairs returns a slice of pairs from the input slice that satisfy the given predicate.
// The function iterates over all possible pairs in the input slice and applies the predicate
// to each pair. If the predicate returns true, the pair is added to the result slice.
//
// T is a generic type parameter that can be any type.
//
// Parameters:
//   - slice: The input slice from which pairs are generated.
//   - predicate: A function that takes two elements of type T and returns a boolean indicating
//     whether the pair satisfies the condition.
//
// Returns:
//
//	A slice of pairs (of type Pair[T]) that satisfy the predicate.
func Pairs[T any](slice []T, predicate func(T, T) bool) []Pair[T] {
	result := make([]Pair[T], 0)
	var b T
	for i, a := range slice {
		for j := i + 1; j < len(slice); j++ {
			b = slice[j]
			if predicate(a, slice[j]) {
				pair := &Pair[T]{a, b}
				result = append(result, *pair)
			}
		}
	}
	return result
}

// MinDistance calculates the minimum distance between any two elements from two slices.
// The distance between elements is determined by the provided distance function.
//
// Parameters:
//   - slice1: The first slice of elements.
//   - slice2: The second slice of elements.
//   - distance: A function that takes pointers to two elements (one from each slice) and returns an integer representing the distance between them.
//
// Returns:
//
//	The minimum distance between any two elements from the two slices.
func MinDistance[T any](slice1 []T, slice2 []T, distance func(*T, *T) int) int {
	result := math.MaxInt
	var tmp int
	for _, a := range slice1 {
		for _, b := range slice2 {
			tmp = distance(&a, &b)
			if tmp < result {
				result = tmp
			}
		}
	}
	return result
}

// SameElements checks if two slices contain exactly the same elements, regardless of order.
// It returns true if both slices contain the same elements, otherwise false.
//
// T is a type parameter that must be comparable.
//
// Parameters:
//   - slice1: The first slice to compare.
//   - slice2: The second slice to compare.
//
// Returns:
//   - bool: true if both slices contain the same elements, false otherwise.
func SameElements[T comparable](slice1 []T, slice2 []T) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	elementMap := make(map[T]int, len(slice1))
	for _, val := range slice1 {
		elementMap[val]++
	}
	for _, val := range slice2 {
		if count, found := elementMap[val]; !found || count == 0 {
			return false
		}
		elementMap[val]--
	}
	return true
}
