package main

import (
	"fmt"
	"strings"
)

func Reverse(word string) string {
	reversed := ""
	for _, v := range word {
		letter := fmt.Sprintf("%c", v)
		reversed = strings.Join([]string{letter, reversed}, "")
	}

	return reversed
}

func Encrypt(word string) string {

	hashes := map[string]string{
		"a": "0",
		"e": "1",
		"i": "2",
		"o": "2",
		"u": "3",
	}

	word = Reverse(word)

	for hash, value := range hashes {
		if strings.Contains(word, hash) {
			word = strings.ReplaceAll(word, hash, value)
		}
	}

	return word + "aca"
}
