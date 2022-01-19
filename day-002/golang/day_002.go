package main

import (
	"fmt"
	"strings"
)

const (
	foundPrefix = `I found Nemo at `
	notFound    = `I can't find Nemo :(`
)

func findNemo(phrase string) string {
	words := strings.Split(phrase, " ")

	for index, word := range words {
		if word == "Nemo" {
			result := fmt.Sprintf("%v%v!", foundPrefix, index+1)
			return result
		}
	}

	return notFound

}
