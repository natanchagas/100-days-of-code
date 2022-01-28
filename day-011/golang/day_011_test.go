package main

import (
	"fmt"
	"reflect"
	"testing"
)

func ExampleIsValidIP() {
	isValid := IsValidIP("255.255.11.135")
	fmt.Println(isValid)
	// Output: true
}

func TestIsValidIP(t *testing.T) {
	assert := func(t testing.TB, got bool, want bool) {
		t.Helper()

		if got != want {
			t.Errorf("Got: %v,\nWant: %v", got, want)
		}
	}

	t.Run("255.255.11.135", func(t *testing.T) {
		got := IsValidIP("255.255.11.135")
		want := true

		assert(t, got, want)
	})

	t.Run("0.0.0.0", func(t *testing.T) {
		got := IsValidIP("0.0.0.0")
		want := true

		assert(t, got, want)
	})

	t.Run("0.100.1.0", func(t *testing.T) {
		got := IsValidIP("0.100.1.0")
		want := true

		assert(t, got, want)
	})

	t.Run("1.0.102.3", func(t *testing.T) {
		got := IsValidIP("1.0.102.3")
		want := true

		assert(t, got, want)
	})

	t.Run("101.0.2.3", func(t *testing.T) {
		got := IsValidIP("101.0.2.3")
		want := true

		assert(t, got, want)
	})
}

func ExampleGeIPAddresses() {
	ipAdresses := GeIPAddresses("25525511135")
	fmt.Println(ipAdresses)
	// Output: [255.255.11.135 255.255.111.35]
}

func TestGeIPAddresses(t *testing.T) {
	assert := func(t testing.TB, got []string, want []string) {
		t.Helper()

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Got: %v,\nWant: %v", got, want)
		}
	}

	t.Run("25525511135", func(t *testing.T) {
		got := GeIPAddresses("25525511135")
		want := []string{"255.255.11.135", "255.255.111.35"}

		assert(t, got, want)
	})

	t.Run("0000", func(t *testing.T) {
		got := GeIPAddresses("0000")
		want := []string{"0.0.0.0"}

		assert(t, got, want)
	})

	t.Run("1111", func(t *testing.T) {
		got := GeIPAddresses("1111")
		want := []string{"1.1.1.1"}

		assert(t, got, want)
	})

	t.Run("010010", func(t *testing.T) {
		got := GeIPAddresses("010010")
		want := []string{"0.10.0.10", "0.100.1.0"}

		assert(t, got, want)
	})

	t.Run("101023", func(t *testing.T) {
		got := GeIPAddresses("101023")
		want := []string{"1.0.10.23", "1.0.102.3", "10.1.0.23", "10.10.2.3", "101.0.2.3"}

		assert(t, got, want)
	})
}
