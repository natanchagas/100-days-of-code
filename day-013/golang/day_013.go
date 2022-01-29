package main

import (
	"sort"
)

func Overweight(weights []int, add int) bool {
	sum := 0

	for i := 0; i < len(weights); i++ {
		sum += weights[i]
	}

	return sum+add > 10

}

func CanFit(weights []int, bags int) bool {

	organizing := make([][]int, 4)
	cantfit := []int{}

	sort.Sort(sort.Reverse(sort.IntSlice(weights)))

	for _, weight := range weights {
		notAdded := true
		for i := 0; i < bags; i++ {
			if !Overweight(organizing[i], weight) && notAdded {
				organizing[i] = append(organizing[i], weight)
				notAdded = false
			}
		}
		if notAdded {
			cantfit = append(cantfit, weight)
		}

	}

	if len(cantfit) == 0 {
		return true
	}

	return false
}
