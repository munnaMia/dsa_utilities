package math

// Give a number factorial.
func Factorial(n int) int {
	if n == 1 {
		return n
	}
	return n * Factorial(n-1)
}

// find all the divisor of a number
func FindAllDivisor1(n int) []int {
	result := []int{}
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			result = append(result, i)
		}
	}

	return result
}

// find all the divisor of a number
func FindAllDivisor2(n int) []int {
	result := []int{}
	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			if i != n/i {
				result = append(result, i, n/i)
			} else {
				result = append(result, i)
			}
		}
	}

	return result
}

// return x^n as float64 value
func Pow(x, n int) float64 {
	var ans float64 = 1
	pIsPositive := true

	if n < 0 {
		pIsPositive = false
		n *= -1
	}

	if n == 0 {
		return 1
	}

	for n > 0 {
		if n%2 == 0 {
			x *= x
			n /= 2
		} else {
			ans *= float64(x)
			n--
		}
	}

	if pIsPositive {
		return ans
	}
	return 1 / ans
}

// print prime number between L to R number  and tell how many prime are there
// func FindPrime(L, R int) int {

// }

// PrimeFactor finds all unique prime factors of a given number n using trial division up to sqrt(n).
//
// Logic:
// It checks all potential divisors i from 1 up to sqrt(n).
// If i is a factor, it tests whether both i and its pair (n / i) are prime.
//
// Examples:
//
//	PrimeFactor(12)  => [2, 3]    (Factors of 12 are 1, 2, 3, 4, 6, 12; primes are 2 and 3)
//	PrimeFactor(28)  => [2, 7]    (Factors of 28 are 1, 2, 4, 7, 14, 28; primes are 2 and 7)
//	PrimeFactor(13)  => [13]      (Prime number itself)
//
// Note: Starts from i = 1, so IsPrime(1) should return false to avoid including 1.
func PrimeFactor(n int) (result []int) {
	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			if IsPrime(i) {
				result = append(result, i)
			}
			if IsPrime(n / i) {
				if i != n/i {
					result = append(result, n/i)
				}
			}
		}
	}
	return
}

// return prime factor of a number more efficient school standard method
// PrimeFactor1 extracts unique prime factors using prime factorization by successive division.
//
// Logic:
// Starting from i = 2, whenever a prime factor i is found, it strips away ALL occurrences
// of i from n by repeatedly dividing n by i. This guarantees n reduces faster.
// Any remaining value of n after the loop is guaranteed to be prime (or 1).
//
// Examples:
//
//	PrimeFactor1(12) => [2, 3]    (Finds 2, divides 12 by 2 twice to leave 3; 3 added at the end)
//	PrimeFactor1(100) => [2, 5]   (Finds 2, divides by 2 twice to leave 25; loop continues and finds 5)
//	PrimeFactor1(13) => [13]      (Loop doesn't run for prime <= 13; appends 13 at the end)
//
// Note:
// The loop condition `i*i < n` leaves edge cases for perfect squares (e.g., n = 9, 25, 49).
// Changing the condition to `i*i <= n` prevents missing perfect squares.
func PrimeFactor1(n int) (result []int) {
	for i := 2; i*i < n; i++ {
		if n%i == 0 {
			result = append(result, i)
			for n%i == 0 {
				n /= i
			}
		}
	}

	if n != 1 {
		result = append(result, n)
	}
	return
}

func PrimeFactorization(n int) (result []int) {
	spf := SmallestPrimeFactor(100000)

	for i := spf[n]; i*i < n; i++ {
		if n%i == 0 {
			for n%i == 0 {
				result = append(result, i)
				n /= i
			}
		}
	}

	if n != 1 {
		result = append(result, n)
	}
	return
}
