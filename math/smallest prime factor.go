package math

// prine 1 to n all primes
func SmallestPrimeFactor(n int) []int {
	PrimeArray := make([]int, n+1)

	for i := 2; i < len(PrimeArray); i++ {
		PrimeArray[i] = i
	}

	// outer loop run upto lenght sqr root
	for i := 0; i*i <= len(PrimeArray); i++ {
		if PrimeArray[i] != 0 {
			for k := i * i; k < len(PrimeArray); k += i {
				if i < PrimeArray[k] {
					PrimeArray[k] = i
				}
			}
		}

	}

	return PrimeArray
}
