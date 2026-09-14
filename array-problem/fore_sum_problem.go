package arrayproblem

import (
	"github.com/munnaMia/Data-Structure-Algorithms/sorting"
)

func ForeSumProblem(target int, arr ...int) [][]int {
	sorting.Quicksort(0, len(arr)-1, arr)
	combinationMap := make([][]int, 0)

	for i := 0; i < len(arr); i++ {
		if i > 0 && arr[i] == arr[i-1] {
			continue
		}

		for j := i + 1; j < len(arr); j++ {
			if j > i+1 && arr[j] == arr[j-1] {
				continue
			}

			k := j + 1
			l := len(arr) - 1

			for k < l {
				sum := arr[i] + arr[j] + arr[k] + arr[l]
				if sum == target {
					temp := []int{arr[i], arr[j], arr[k], arr[l]}
					combinationMap = append(combinationMap, temp)
					k++
					l--
					for k < l && arr[k] == arr[k-1] {
						k++
					}
					for k < l && arr[l] == arr[l+1] {
						l--
					}
				} else if sum < target {
					k++
				} else {
					l--
				}
			}
		}
	}

	return combinationMap
}
