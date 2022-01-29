package main

func isIncreasing(array []int) bool {
	number := 0

	for i := 0; i < len(array); i++ {
		if i == 0 {
			number = array[i]
		} else {
			if array[i] < number {
				return false
			}
			number = array[i]
		}
	}

	return true
}

func isDecreasing(array []int) bool {
	number := 0

	for i := 0; i < len(array); i++ {
		if i == 0 {
			number = array[i]
		} else {
			if array[i] > number {
				return false
			}
			number = array[i]
		}
	}

	return true
}

func MountainOrValley(heights []int) string {

	data := make(map[string]int)

	for position, height := range heights {

		if position == 0 {
			data["peak"] = height
			data["peak_position"] = position
			data["trough"] = height
			data["trough_position"] = position
		} else {

			if data["peak"] < height {
				data["peak"] = height
				data["peak_position"] = position
			} else if data["trough"] > height {
				data["trough"] = height
				data["trough_position"] = position
			}

		}

	}

	// Check if is peak:
	if isIncreasing(heights[:data["peak_position"]]) && isDecreasing(heights[data["peak_position"]+1:]) && data["peak_position"] != 0 && data["peak_position"] != len(heights)-1 {
		return "mountain"
	} else if isDecreasing(heights[:data["trough_position"]]) && isIncreasing(heights[data["trough_position"]+1:]) && data["trough_position"] != 0 && data["trough_position"] != len(heights)-1 {
		return "valley"
	} else {
		return "neither"
	}
}
