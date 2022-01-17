package main

import "testing"
import "fmt"

func ExampleCalcAge() {
	ageInDays := calcAge(20)
	fmt.Println(ageInDays)
	// Output: 7300
}

func TestCalcAge(t *testing.T) {
	assert := func(t testing.TB, ageInDays int, expectedAgeInDays int) {
		t.Helper()
		if ageInDays != expectedAgeInDays {
			t.Errorf("\nExpected: %v,\nGot %v", expectedAgeInDays, ageInDays)
		}
	}

	t.Run("65 Years", func(t *testing.T) {
		ageInDays := calcAge(65)
		expectedAgeInDays := 23725

		assert(t, ageInDays, expectedAgeInDays)
	})

	t.Run("0 Years", func(t *testing.T) {
		ageInDays := calcAge(0)
		expectedAgeInDays := 0

		assert(t, ageInDays, expectedAgeInDays)
	})

}
