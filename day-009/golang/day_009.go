package main

func TrapWater(elevation []int) int {
	var waterTrapped int

	for i := 0; i < len(elevation); i++ {
		height := elevation[i]

		if height > 0 {
			puddle := 0

			for j := i + 1; j < len(elevation); j++ {
				nextHeight := elevation[j]

				if nextHeight < height {
					if height-nextHeight > 1 {
						puddle += (height - nextHeight)
					} else {
						puddle += 1
					}
				} else if nextHeight >= height {
					waterTrapped += puddle
					i = j - 1
					break
				}

			}
		}

	}

	return waterTrapped
}
