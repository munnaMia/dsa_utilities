package sorting

func BubbleSort(array []int) {
	for i := len(array); i >= 0; i-- {
		didSwap := false

		for j := 0; j < i-1; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
				didSwap = true
			}
		}

		if !didSwap {
			break
		}
	}
}
