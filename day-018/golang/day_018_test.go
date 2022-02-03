package main

import (
	"fmt"
	"testing"
)

func ExampleUniquePaths() {
	unique_paths := UniquePaths(3, 7)
	fmt.Println(unique_paths)
	// Output: 28
}

func TestUniquePaths(t *testing.T) {
	assert := func(t testing.TB, got int, want int) {
		t.Helper()

		if got != want {
			t.Errorf("\nGot: %v,\nWant: %v.", got, want)
		}
	}

	t.Run("3, 7", func(t *testing.T) {
		got := UniquePaths(3, 7)
		want := 28

		assert(t, got, want)
	})

	t.Run("3, 2", func(t *testing.T) {
		got := UniquePaths(3, 2)
		want := 3

		assert(t, got, want)
	})

	t.Run("7, 3", func(t *testing.T) {
		got := UniquePaths(7, 3)
		want := 28

		assert(t, got, want)
	})

	t.Run("3, 3", func(t *testing.T) {
		got := UniquePaths(3, 3)
		want := 6

		assert(t, got, want)
	})
}
