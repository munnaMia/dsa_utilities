package sorting

func MergeSort(low, high int, array []int) {
	if low >= high {
		return
	}
	mid := (low + high) / 2

	// dividing the array
	MergeSort(low, mid, array)
	MergeSort(mid+1, high, array)

	merge(low, mid, high, array)
}

func merge(low, mid, high int, array []int) {
	temp_array := []int{}

	left, right := low, mid+1

	for left <= mid && right <= high {
		if array[left] <= array[right] {
			temp_array = append(temp_array, array[left])
			left++
		} else {
			temp_array = append(temp_array, array[right])
			right++
		}
	}

	for left <= mid {
		temp_array = append(temp_array, array[left])
		left++
	}

	for right <= high {
		temp_array = append(temp_array, array[right])
		right++
	}

	for i := low; i < high+1; i++ {
		array[i] = temp_array[i-low]
	}
}
