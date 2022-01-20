package main

func progressDays(arr []int) int {
	var progressDays int
	var latestDistance int

	for day, distance := range arr {
		if day == 0 {
			latestDistance = distance
		} else {
			if distance > latestDistance {
				progressDays += 1
			}
			latestDistance = distance
		}

	}

	return progressDays
}
