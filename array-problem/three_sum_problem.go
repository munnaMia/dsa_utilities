package arrayproblem

// if the array contain 3 unique value and there sum return 0 the the func return true.
// this solution i am using maps
func ThreeSumBetter(arr ...int) bool {
	hash := make(map[int]int)

	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			k := -(arr[i] + arr[j])
			if _, ok := hash[k]; ok {
				if k == arr[i] || k == arr[j] {
					continue
				}
				return true
			}
			hash[arr[j]] = j
		}
	}

	return false
}
