package main

import (
	"fmt"
	"reflect"
	"testing"
)

func ExampleMerge() {
	possibilities := Merge([]string{"a", "b", "c"}, []string{"d", "e", "f"})
	fmt.Println(possibilities)
	// Output: [ad ae af bd be bf cd ce cf]
}

func TestMerge(t *testing.T) {
	assert := func(t testing.TB, got []string, want []string) {
		t.Helper()

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Got: %v,\nWant: %v", got, want)
		}
	}

	t.Run("{\"a\",\"b\",\"c\"} x {\"d\",\"e\",\"f\"}", func(t *testing.T) {
		got := Merge([]string{"a", "b", "c"}, []string{"d", "e", "f"})
		want := []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}

		assert(t, got, want)
	})
}

func ExampleType() {
	possibilities := Type("23")
	fmt.Println(possibilities)
	// Output: [ad ae af bd be bf cd ce cf]
}

func TestType(t *testing.T) {
	assert := func(t testing.TB, got []string, want []string) {
		t.Helper()

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Got: %v,\nWant: %v", got, want)
		}
	}

	t.Run("", func(t *testing.T) {
		got := Type("")
		want := []string{}

		assert(t, got, want)
	})

	t.Run("2", func(t *testing.T) {
		got := Type("2")
		want := []string{"a", "b", "c"}

		assert(t, got, want)
	})

	t.Run("23", func(t *testing.T) {
		got := Type("23")
		want := []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}

		assert(t, got, want)
	})

	t.Run("234", func(t *testing.T) {
		got := Type("234")
		want := []string{"adg", "adh", "adi", "aeg", "aeh", "aei", "afg", "afh", "afi", "bdg", "bdh", "bdi", "beg", "beh", "bei", "bfg", "bfh", "bfi", "cdg", "cdh", "cdi", "ceg", "ceh", "cei", "cfg", "cfh", "cfi"}

		assert(t, got, want)
	})
}
