package main

import (
	"sort"
)

func MergeArray(arr1 []int, arr2 []int) []int {
	sorted := make([]int, 0)

	sorted = append(sorted, arr1...)
	sorted = append(sorted, arr2...)

	sort.Ints(sorted)

	return sorted
}
