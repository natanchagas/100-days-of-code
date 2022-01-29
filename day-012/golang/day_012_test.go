package main

import (
	"fmt"
	"testing"
)

func ExampleIsIncreasing() {
	is_increasing := isIncreasing([]int{1, 2, 7, 8})
	fmt.Println(is_increasing)
	// Output: true
}

func TestIsIncreasing(t *testing.T) {
	got := isIncreasing([]int{1, 2, 7, 8})
	want := true

	if got != want {
		t.Errorf("Got: %v,\nWant: %v", got, want)
	}
}

func ExampleIsDecreasing() {
	is_decreasing := isDecreasing([]int{9, 6, 3, 0})
	fmt.Println(is_decreasing)
	// Output: true
}

func TestIsDecreasing(t *testing.T) {
	got := isDecreasing([]int{9, 6, 3, 0})
	want := true

	if got != want {
		t.Errorf("Got: %v,\nWant: %v", got, want)
	}
}

func ExampleMountainOrValley() {
	mountain_valley := MountainOrValley([]int{1, 3, 5, 4, 3, 2})
	fmt.Println(mountain_valley)
	// Output: mountain
}

func TestMountainOrValley(t *testing.T) {
	assert := func(t testing.TB, got string, want string) {
		t.Helper()

		if got != want {
			t.Errorf("Got: %v,\nWant: %v", got, want)
		}
	}

	t.Run("3, 4, 5, 4, 3", func(t *testing.T) {
		got := MountainOrValley([]int{3, 4, 5, 4, 3})
		want := "mountain"

		assert(t, got, want)
	})

	t.Run("9, 7, 3, 1, 2, 4", func(t *testing.T) {
		got := MountainOrValley([]int{9, 7, 3, 1, 2, 4})
		want := "valley"

		assert(t, got, want)
	})

	t.Run("-1, -1, 0, -1, -1", func(t *testing.T) {
		got := MountainOrValley([]int{-1, -1, 0, -1, -1})
		want := "mountain"

		assert(t, got, want)
	})

	t.Run("1, 2, 3, 2, 4, 1", func(t *testing.T) {
		got := MountainOrValley([]int{1, 2, 3, 2, 4, 1})
		want := "neither"

		assert(t, got, want)
	})

	t.Run("5, 4, 3, 2, 1", func(t *testing.T) {
		got := MountainOrValley([]int{5, 4, 3, 2, 1})
		want := "neither"

		assert(t, got, want)
	})

	t.Run("0, -1, -1, 0, -1, -1", func(t *testing.T) {
		got := MountainOrValley([]int{0, -1, -1, 0, -1, -1})
		want := "neither"

		assert(t, got, want)
	})
}
