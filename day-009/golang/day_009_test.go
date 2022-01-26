package main

import (
	"fmt"
	"testing"
)

func ExampleTrapWater() {
	amount := TrapWater([]int{4, 2, 0, 3, 2, 5})
	fmt.Println(amount)
	// Output: 9
}

func TestTrapWater(t *testing.T) {
	assert := func(t testing.TB, got int, want int) {
		t.Helper()

		if got != want {
			t.Errorf("Got: %v,\nWant: %v", got, want)
		}
	}

	t.Run("[0,1,0,2,1,0,1,3,2,1,2,1]", func(t *testing.T) {
		got := TrapWater([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1})
		want := 6

		assert(t, got, want)
	})

	t.Run("[4,2,0,3,2,5]", func(t *testing.T) {
		got := TrapWater([]int{4, 2, 0, 3, 2, 5})
		want := 9

		assert(t, got, want)
	})
}
