package main

import (
	"fmt"
	"testing"
)

func ExampleCanWin() {
	can_win := CanWin(4)
	fmt.Println(can_win)
	// Output: false
}

func TestCanWin(t *testing.T) {
	assert := func(t testing.TB, got bool, want bool) {
		t.Helper()

		if got != want {
			t.Errorf("Got: %v,\nWant: %v", got, want)
		}
	}

	t.Run("4", func(t *testing.T) {
		got := CanWin(4)
		want := false

		assert(t, got, want)
	})

	t.Run("2", func(t *testing.T) {
		got := CanWin(2)
		want := true

		assert(t, got, want)
	})

	t.Run("5", func(t *testing.T) {
		got := CanWin(5)
		want := true

		assert(t, got, want)
	})
}
