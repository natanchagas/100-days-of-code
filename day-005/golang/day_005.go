package main

import (
	"strings"
)

func SockPairs(socks string) int {
	var pairs int
	var analyzed string

	for _, unit := range socks {
		if !strings.Contains(analyzed, string(unit)) {
			occurence := strings.Count(socks, string(unit))
			pairs += occurence / 2
			analyzed += string(unit)
		}
	}

	return pairs
}
