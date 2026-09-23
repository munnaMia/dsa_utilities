package math

// prine 1 to n all primes
func SeiveOfEratosthenes(n int) []int {
	PrimeArray := make([]int, n+1)

	for i := 2; i < len(PrimeArray); i++ {
		PrimeArray[i] = 1
	}

	// outer loop run upto lenght sqr root
	for i := 0; i*i <=len(PrimeArray); i++ {
		if PrimeArray[i] == 1 {
			for k := i * i; k < len(PrimeArray); k += i {
				PrimeArray[k] = 0
			}
		}

	}

	return PrimeArray
}
