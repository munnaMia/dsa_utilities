package sorting

func Quicksort(low, high int, array []int) {
	if low < high {
		partitionIdx := findPartition(low, high, array)

		Quicksort(low, partitionIdx-1, array)
		Quicksort(partitionIdx+1, high, array)
	}

}

func findPartition(low, high int, array []int) int {
	pivot := array[low]

	i := low
	j := high

	for i < j {
		for array[i] <= pivot && i < high {
			i++
		}

		for array[j] > pivot && j > low {
			j--
		}

		if i < j {
			array[i], array[j] = array[j], array[i]
		}
	}

	array[low], array[j] = array[j], array[low] // swap with the pivot

	return j
}
