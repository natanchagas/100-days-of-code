package main

func UniquePaths(height int, length int) int {

	paths := make([][]int, length)

	for i := 0; i < length; i++ {
		paths[i] = make([]int, height)
	}

	for i := 0; i < height; i++ {
		paths[0][i] = 1
	}

	for i := 0; i < length; i++ {
		paths[i][0] = 1
	}

	for i := 1; i < length; i++ {
		for j := 1; j < height; j++ {
			paths[i][j] = paths[i-1][j] + paths[i][j-1]
		}
	}

	return paths[length-1][height-1]
}
