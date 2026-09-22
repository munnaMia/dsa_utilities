package math

// return bool if a number is a prime or not
func IsPrime(n int) bool {
	if n == 2 {
		return true
	}
	if n%2 == 0 || n <= 1 {
		return false
	}

	for i := 3; i*i <= n; i += 2 {
		if n%i == 0 {
			return false
		}
	}

	return true

}
