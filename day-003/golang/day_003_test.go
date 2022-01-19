package main

import (
	"fmt"
	"reflect"
	"testing"
)

func ExampleCheckServings() {
	servings := []string{
		"--xo--x--ox--",
		"--xx--x--xx--",
		"--oo--o--oo--",
		"--xx--x--ox--",
		"--xx--x--ox--",
	}

	servingsCount := checkServings(servings)
	fmt.Println(servingsCount)
	// Output: [1 4]
}

func TestCheckServings(t *testing.T) {
	assert := func(t testing.TB, got []int, want []int) {
		t.Helper()

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Got: %v,\nWant: %v", got, want)
		}
	}

	t.Run("[\"--oooo-ooo--\",\"--xx--x--xx--\",\"--o---o--oo--\",\"--xx--x--ox--\",\"--xx--x--ox--\"]", func(t *testing.T) {

		servings := []string{
			"--oooo-ooo--",
			"--xx--x--xx--",
			"--o---o--oo--",
			"--xx--x--ox--",
			"--xx--x--ox--",
		}

		got := checkServings(servings)
		want := []int{2, 3}

		assert(t, got, want)
	})

	t.Run("[\"--oooo-ooo--\",\"--xxxxxxxx--\",\"--o---\",\"-o-----o---x--\",\"--o---o-----\"]", func(t *testing.T) {

		servings := []string{
			"--oooo-ooo--",
			"--xxxxxxxx--",
			"--o---",
			"-o-----o---x--",
			"--o---o-----",
		}

		got := checkServings(servings)
		want := []int{3, 2}

		assert(t, got, want)
	})
}
