package sorting

func RecursiveInsertionSort(idx int, array []int){
	if len(array) == idx {
		return
	}

	j := idx 
	for j>0 && array[j-1] > array[j]{
		array[j-1] ,array[j]= array[j] , array[j-1]
		j--
	}

	RecursiveInsertionSort(idx+1, array)
}