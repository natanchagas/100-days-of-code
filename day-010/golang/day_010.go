package main

func Factorial(num int) int {
	factorial := 1
	for i := 1; i <= num; i++ {
		factorial *= i
	}
	return factorial
}

func UniqueBSTs(num int) int {
	return Factorial(2*num) / (Factorial(num+1) * Factorial(num))
}
