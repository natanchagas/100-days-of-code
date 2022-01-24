package main

import (
	"fmt"
	"reflect"
	"testing"
)

func ExampleMergeArray() {
	arr1 := []int{9, 7, 11, 5}
	arr2 := []int{3, 2, 4, 8}

	sorted := MergeArray(arr1, arr2)
	fmt.Println(sorted)
	// Output: [2 3 4 5 7 8 9 11]
}

func TestMergeArray(t *testing.T) {
	arr1 := []int{1, 2, 3}
	arr2 := []int{2, 5, 6}

	got := MergeArray(arr1, arr2)
	want := []int{1, 2, 2, 3, 5, 6}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Got: %v,\nWant: %v", got, want)
	}

}
