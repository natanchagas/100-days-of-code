package main

import (
	"fmt"
	"testing"
)

func ExampleSocksPairs() {
	socks := "AA"
	fmt.Println(SockPairs(socks))
	// Output: 1
}

func TestSockPairs(t *testing.T) {
	assert := func(t testing.TB, got int, want int) {
		t.Helper()

		if got != want {
			t.Errorf("Got: %v,\nWant: %v", got, want)
		}
	}

	t.Run("CABBACCC", func(t *testing.T) {
		socks := "CABBACCC"

		got := SockPairs(socks)
		want := 4

		assert(t, got, want)
	})

	t.Run("ABABC", func(t *testing.T) {
		socks := "ABABC"

		got := SockPairs(socks)
		want := 2

		assert(t, got, want)
	})
}
