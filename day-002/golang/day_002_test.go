package main

import (
	"fmt"
	"testing"
)

func ExampleFindNemo() {
	nemoSearch := findNemo("I Nemo am")
	fmt.Println(nemoSearch)
	// Output: I found Nemo at 2!
}

func TestFindNemo(t *testing.T) {
	assert := func(t testing.TB, nemoSearch string, want string) {
		t.Helper()
		if nemoSearch != want {
			t.Errorf("Want: %v,\nGot: %v", want, nemoSearch)
		}
	}

	t.Run("I am finding Nemo !", func(t *testing.T) {
		nemoSearch := findNemo("I am finding Nemo !")
		want := "I found Nemo at 4!"

		assert(t, nemoSearch, want)
	})

	t.Run("Nemo is me", func(t *testing.T) {
		nemoSearch := findNemo("Nemo is me")
		want := "I found Nemo at 1!"

		assert(t, nemoSearch, want)
	})

	t.Run("He is not here", func(t *testing.T) {
		nemoSearch := findNemo("He is not here")
		want := "I can't find Nemo :("

		assert(t, nemoSearch, want)
	})
}
