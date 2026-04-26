package sorting

func RecursiveBubbleSort(len int, array []int) {
	if 1 == len {
		return
	}

	swap := false

	for j := 0; j < len-1; j++ {
		if array[j] > array[j+1] {
			array[j+1], array[j] = array[j], array[j+1]
			swap = true
		}
	}

	if !swap {
		return
	}

	RecursiveBubbleSort(len-1, array)
}
