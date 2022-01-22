package main

func IsPrime(num int) bool {
	var divisor []int

	for i := num; i > 0; i-- {
		if num%i == 0 {
			divisor = append(divisor, i)
		}
	}

	if len(divisor) == 2 {
		return true
	} else {
		return false
	}

}

func NextPrime(prime int) int {
	for i := prime; ; i++ {
		if IsPrime(i) {
			return i
		}
	}
}
