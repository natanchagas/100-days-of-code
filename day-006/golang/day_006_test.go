package main

import (
	"fmt"
	"testing"
)

func ExampleIsPrime() {
	isPrime := IsPrime(13)
	fmt.Println(isPrime)
	// Output: true
}

func TestIsPrime(t *testing.T) {
	assert := func(t testing.TB, got bool, want bool) {
		t.Helper()

		if got != want {
			t.Errorf("Got: %v,\nWant: %v", got, want)
		}
	}

	t.Run("24", func(t *testing.T) {
		got := IsPrime(24)
		want := false

		assert(t, got, want)
	})

	t.Run("11", func(t *testing.T) {
		got := IsPrime(11)
		want := true

		assert(t, got, want)
	})
}

func ExampleNextPrime() {
	nextPrime := NextPrime(12)
	fmt.Println(nextPrime)
	// Output: 13
}

func TestNextPrime(t *testing.T) {
	assert := func(t testing.TB, got int, want int) {
		t.Helper()

		if got != want {
			t.Errorf("Got: %v,\nWant: %v", got, want)
		}
	}

	t.Run("24", func(t *testing.T) {
		got := NextPrime(24)
		want := 29

		assert(t, got, want)
	})

	t.Run("11", func(t *testing.T) {
		got := NextPrime(11)
		want := 11

		assert(t, got, want)
	})
}
