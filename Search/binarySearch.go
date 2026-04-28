package search

func BinarySearch(target int, array []int) (bool, int) {
	low := 0
	high := len(array) - 1

	for low <= high {
		mid := (low + high) / 2

		if array[mid] == target {
			return true, mid
		} else if array[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return false, -1
}
