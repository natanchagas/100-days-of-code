package main

import (
	"strconv"
)

var keypad = map[int][]string{
	2: {"a", "b", "c"},
	3: {"d", "e", "f"},
	4: {"g", "h", "i"},
	5: {"j", "k", "l"},
	6: {"m", "n", "o"},
	7: {"p", "q", "r", "s"},
	8: {"t", "u", "v"},
	9: {"w", "x", "y", "z"},
}

func Merge(pool1, pool2 []string) []string {
	var merged []string

	for _, p1 := range pool1 {
		for _, p2 := range pool2 {
			merged = append(merged, p1+p2)
		}
	}

	return merged
}

func Type(digits string) []string {

	length := len(digits)

	switch length {
	case 0:
		return []string{}
	case 1:
		integer, _ := strconv.Atoi(string(digits))
		return keypad[integer]
	}

	var possibilities []string
	for i := 0; i < length; i++ {
		integer, _ := strconv.Atoi(string(digits[i]))
		if i == 0 {
			possibilities = keypad[integer]
		} else {
			possibilities = Merge(possibilities, keypad[integer])
		}
	}

	return possibilities

}
