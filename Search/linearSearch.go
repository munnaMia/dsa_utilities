package search

func linearSearch(target int, array []int) bool {
	for _, value := range array {
		if target == value {
			return true
		}
	}
	return false
}
