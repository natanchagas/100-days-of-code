package main

import (
	"fmt"
	"testing"
)

func ExampleReverse() {
	reversed := Reverse("banana")
	fmt.Println(reversed)
	// Output: ananab
}

func TestReverse(t *testing.T) {
	assert := func(t testing.TB, got string, want string) {
		t.Helper()

		if got != want {
			t.Errorf("Got: %v,\nWant: %v", got, want)
		}
	}

	t.Run("banana", func(t *testing.T) {
		got := Reverse("banana")
		want := "ananab"

		assert(t, got, want)
	})

	t.Run("karaca", func(t *testing.T) {
		got := Reverse("karaca")
		want := "acarak"

		assert(t, got, want)
	})

	t.Run("burak", func(t *testing.T) {
		got := Reverse("burak")
		want := "karub"

		assert(t, got, want)
	})

	t.Run("alpaca", func(t *testing.T) {
		got := Reverse("alpaca")
		want := "acapla"

		assert(t, got, want)
	})
}

func ExampleEncrypt() {
	encrypted := Encrypt("banana")
	fmt.Println(encrypted)
	// Output: 0n0n0baca
}

func TestEncrypt(t *testing.T) {
	assert := func(t testing.TB, got string, want string) {
		t.Helper()

		if got != want {
			t.Errorf("Got: %v,\nWant: %v", got, want)
		}
	}

	t.Run("banana", func(t *testing.T) {
		got := Encrypt("banana")
		want := "0n0n0baca"

		assert(t, got, want)
	})

	t.Run("karaca", func(t *testing.T) {
		got := Encrypt("karaca")
		want := "0c0r0kaca"

		assert(t, got, want)
	})

	t.Run("burak", func(t *testing.T) {
		got := Encrypt("burak")
		want := "k0r3baca"

		assert(t, got, want)
	})

	t.Run("alpaca", func(t *testing.T) {
		got := Encrypt("alpaca")
		want := "0c0pl0aca"

		assert(t, got, want)
	})

}
