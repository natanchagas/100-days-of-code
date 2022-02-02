package main

func FreedPrisoners(prisioners []int) int {

	ReverseValues := func(values *[]int) {

		for i := 0; i < len(*values); i++ {
			if (*values)[i] == 0 {
				(*values)[i] = 1
			} else {
				(*values)[i] = 0
			}
		}
	}

	if prisioners[0] == 0 {
		return 0
	}

	freed := 0

	for i := 1; i < len(prisioners); i++ {
		if prisioners[i] == 1 {
			freed += 1
			ReverseValues(&prisioners)
		}
	}

	return freed

}

func main() {
	FreedPrisoners([]int{1, 1, 0, 0, 0, 1, 0})
}
