package main

import (
	"fmt"
	"testing"
)

func ExampleProgressDays() {
	progressDays := progressDays([]int{3, 4, 1, 2})
	fmt.Println(progressDays)
	// Output: 2
}

func TestProgressDays(t *testing.T) {
	assert := func(t testing.TB, got int, want int) {
		t.Helper()

		if got != want {
			t.Errorf("Got: %v,\nWant: %v", got, want)
		}
	}

	t.Run("[10,11,12,9,10]", func(t *testing.T) {
		got := progressDays([]int{10, 11, 12, 9, 10})
		want := 3

		assert(t, got, want)

	})

	t.Run("[6,5,4,3,2,9]", func(t *testing.T) {
		got := progressDays([]int{6, 5, 4, 3, 2, 9})
		want := 1

		assert(t, got, want)

	})

	t.Run("[9,9]", func(t *testing.T) {
		got := progressDays([]int{9, 9})
		want := 0

		assert(t, got, want)

	})
}
