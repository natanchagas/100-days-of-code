package main

import (
	"fmt"
	"testing"
)

func ExampleOverweight() {
	is_overweight := Overweight([]int{2, 1, 2}, 5)
	fmt.Println(is_overweight)
	// Outuput: false
}

func TestOverweight(t *testing.T) {
	assert := func(t testing.TB, got bool, want bool) {
		t.Helper()

		if got != want {
			t.Errorf("Got: %v,\nWant: %v", got, want)
		}
	}

	t.Run("[2, 1, 2], 5", func(t *testing.T) {
		got := Overweight([]int{2, 1, 2}, 5)
		want := false

		assert(t, got, want)
	})

	t.Run("[2, 1, 2, 5], 5", func(t *testing.T) {
		got := Overweight([]int{2, 1, 2, 5}, 5)
		want := true

		assert(t, got, want)
	})
}

func ExampleCanFit() {
	can_fit := CanFit([]int{2, 1, 2, 5, 4, 3, 6, 1, 1, 9, 3, 2}, 4)
	fmt.Println(can_fit)
	// Outuput: true
}

func TestCanFit(t *testing.T) {
	assert := func(t testing.TB, got bool, want bool) {
		t.Helper()

		if got != want {
			t.Errorf("Got: %v,\nWant: %v", got, want)
		}
	}

	t.Run("[2, 1, 2, 5, 4, 3, 6, 1, 1, 9, 3, 2], 4", func(t *testing.T) {
		got := CanFit([]int{2, 1, 2, 5, 4, 3, 6, 1, 1, 9, 3, 2}, 4)
		want := true

		assert(t, got, want)
	})

	t.Run("[2, 7, 1, 3, 3, 4, 7, 4, 1, 8, 2], 4", func(t *testing.T) {
		got := CanFit([]int{2, 7, 1, 3, 3, 4, 7, 4, 1, 8, 2}, 4)
		want := false

		assert(t, got, want)
	})
}
