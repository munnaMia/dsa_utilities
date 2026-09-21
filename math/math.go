package math

// Give a number factorial.
func Factorial(n int) int {
	if n == 1 {
		return n
	}
	return n * Factorial(n-1)
}

func FindAllDivisor1(n int) []int {
	result := []int{}
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			result = append(result, i)
		}
	}

	return result
}

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
