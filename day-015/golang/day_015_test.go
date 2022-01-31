package main

import (
	"fmt"
	"testing"
)

func ExampleIsAnagram() {
	is_anagram := IsAnagram("flower", "reflow")
	fmt.Println(is_anagram)
	// Output: true
}

func TestIsAnagram(t *testing.T) {
	assert := func(t testing.TB, got bool, want bool) {
		t.Helper()

		if got != want {
			t.Errorf("Got: %v,\nWant: %v", got, want)
		}
	}

	t.Run("Flower, Reflow", func(t *testing.T) {
		got := IsAnagram("flower", "reflow")
		want := true

		assert(t, got, want)
	})

	t.Run("anagram, nagaram", func(t *testing.T) {
		got := IsAnagram("anagram", "nagaram")
		want := true

		assert(t, got, want)
	})

	t.Run("nagaram, car", func(t *testing.T) {
		got := IsAnagram("nagaram", "car")
		want := false

		assert(t, got, want)
	})
}
