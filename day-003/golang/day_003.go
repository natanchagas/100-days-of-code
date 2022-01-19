package main

import "strings"

func checkServings(servings []string) []int {
	var veggieSkewer int
	var nonVeggieSkewer int

	for _, skewer := range servings {
		if strings.Contains(skewer, "x") {
			nonVeggieSkewer += 1
		} else {
			veggieSkewer += 1
		}
	}
	return []int{veggieSkewer, nonVeggieSkewer}
}
