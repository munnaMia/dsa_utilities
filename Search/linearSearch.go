package search

func LinearSearch(target int, array []int) (bool, int) {
	for idx, value := range array {
		if target == value {
			return true, idx
		}
	}
	return false, -1
}
