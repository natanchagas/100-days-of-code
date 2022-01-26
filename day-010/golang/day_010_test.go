package main

import (
	"fmt"
	"testing"
)

func ExampleFactorial() {
	factorial := Factorial(3)
	fmt.Println(factorial)
	// Output: 6
}

func TestFactorial(t *testing.T) {
	assert := func(t testing.TB, got int, want int) {
		t.Helper()

		if got != want {
			t.Errorf("Got: %v,\nWant: %v", got, want)
		}
	}

	t.Run("3", func(t *testing.T) {
		got := Factorial(3)
		want := 6

		assert(t, got, want)
	})

	t.Run("5", func(t *testing.T) {
		got := Factorial(5)
		want := 120

		assert(t, got, want)
	})
}

func ExampleUniqueBSTs() {
	bsts := UniqueBSTs(3)
	fmt.Println(bsts)
	// Output: 5
}

func TestUniqueBSTs(t *testing.T) {
	assert := func(t testing.TB, got int, want int) {
		t.Helper()

		if got != want {
			t.Errorf("Got: %v,\nWant: %v", got, want)
		}
	}

	t.Run("3", func(t *testing.T) {
		got := UniqueBSTs(3)
		want := 5

		assert(t, got, want)
	})

	t.Run("4", func(t *testing.T) {
		got := UniqueBSTs(4)
		want := 14

		assert(t, got, want)
	})
}
