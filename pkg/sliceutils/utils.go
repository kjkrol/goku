package sliceutils

import "math"

func Intersects[T comparable](a []T, b []T) bool {
	for _, x := range a {
		for _, y := range b {
			if x == y {
				return true
			}
		}
	}
	return false
}

func Unique[T any, V comparable](input []T, key func(T) V) []T {
	memo := make(map[V]bool)
	var output []T
	for _, v := range input {
		if _, ok := memo[key(v)]; !ok {
			memo[key(v)] = true
			output = append(output, v)
		}
	}
	return output
}

type Pair[T any] struct {
	First  T
	Second T
}

func Map[T any, R any](arr []T, transform func(T) R) []R {
	res := make([]R, len(arr))
	for i, v := range arr {
		res[i] = transform(v)
	}
	return res
}

func Pairs[T any](arr []T, predicate func(T, T) bool) []Pair[T] {
	res := make([]Pair[T], 0)
	var b T
	for i, a := range arr {
		for j := i + 1; j < len(arr); j++ {
			b = arr[j]
			if predicate(a, arr[j]) {
				pair := &Pair[T]{a, b}
				res = append(res, *pair)
			}
		}
	}
	return res
}

func MinDistance[T any](arr1 []T, arr2 []T, distance func(*T, *T) int) int {
	res := math.MaxInt
	var tmp int
	for _, a := range arr1 {
		for _, b := range arr2 {
			tmp = distance(&a, &b)
			if tmp < res {
				res = tmp
			}
		}
	}
	return res
}
