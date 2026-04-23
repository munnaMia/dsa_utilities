package sorting

func SelectionSort(array []int) {
	for i := range array {
		min := i

		for j := i + 1; j < len(array); j++ {
			if array[j] < array[min] {
				min = j
			}
		}

		array[i], array[min] = array[min], array[i]
	}
}
