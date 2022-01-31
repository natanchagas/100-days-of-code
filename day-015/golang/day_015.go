package main

import (
	"reflect"
	"sort"
)

func IsAnagram(s string, t string) bool {

	var s_chars []int
	var t_chars []int

	for _, v := range s {
		s_chars = append(s_chars, int(v))
	}

	for _, v := range t {
		t_chars = append(t_chars, int(v))
	}

	sort.Ints(s_chars)
	sort.Ints(t_chars)

	return reflect.DeepEqual(s_chars, t_chars)
}
