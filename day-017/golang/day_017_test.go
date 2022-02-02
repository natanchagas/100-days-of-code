package main

import (
	"fmt"
	"testing"
)

func ExampleFreedPrisoners() {
	freed_prisioners := FreedPrisoners([]int{1, 1, 0, 0, 0, 1, 0})
	fmt.Println(freed_prisioners)
	// Output: 4
}

func TestFreedPrisoners(t *testing.T) {
	assert := func(t testing.TB, got int, want int) {
		t.Helper()

		if got != want {
			t.Errorf("Got: %v,\n Want: %v", got, want)
		}
	}

	t.Run("[1, 1, 0, 0, 0, 1, 0]", func(t *testing.T) {
		got := FreedPrisoners([]int{1, 1, 0, 0, 0, 1, 0})
		want := 4

		assert(t, got, want)
	})

	t.Run("[1, 1, 1]", func(t *testing.T) {
		got := FreedPrisoners([]int{1, 1, 1})
		want := 1

		assert(t, got, want)
	})

	t.Run("[0, 0, 0]", func(t *testing.T) {
		got := FreedPrisoners([]int{0, 0, 0})
		want := 0

		assert(t, got, want)
	})

	t.Run("[0, 1, 1, 1]", func(t *testing.T) {
		got := FreedPrisoners([]int{0, 1, 1, 1})
		want := 0

		assert(t, got, want)
	})
}
