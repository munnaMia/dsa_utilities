package arrayproblem

import "fmt"

func LargestElem(arr []int) int {
	l := arr[0]

	for _, v := range arr {
		if l < v {
			l = v
		}
	}

	return l
}

func MaximuxConsOnes(arr []int) int {
	f := 0
	count := 0
	tempC := 0
	for f < len(arr) {
		if arr[f] == 1 {
			tempC++
		} else {
			if count < tempC {
				count = tempC
			}
			tempC = 0
		}
		f++
	}

	return count
}

func MissingNum(arr []int) int {
	xor1 := 1
	for i := 2; i <= len(arr)+1; i++ {
		xor1 = xor1 ^ i
	}

	xor2 := arr[0]
	for i := 1; i < len(arr); i++ {
		xor2 = xor2 ^ arr[i]
	}

	return xor1 ^ xor2
}

func SecoundElem(arr []int) int {
	l := arr[0]
	sl := 0

	for _, v := range arr {
		if l < v {
			sl = l
			l = v

		} else if sl < v && sl < l {
			sl = v
		}
	}

	return sl
}

func IsSorted(arr []int) bool {
	for i := 1; i < len(arr); i++ {
		if arr[i-1] > arr[i] {
			return false
		}
	}

	return true
}

func RemoveDuplicateFromSorted(arr []int) []int {
	f := 0
	s := f + 1
	for s < len(arr) {
		if arr[f] == arr[s] {
			s++
		} else {
			f++
			arr[f] = arr[s]
		}
	}
	return arr[:f+1]
}

func FindAppareOnce(arr []int) int {
	n := 0
	for _, v := range arr {
		n = n ^ v
	}

	return n
}

func LeftRotateByOne(arr []int) {
	temp := arr[0]
	for i := 1; i < len(arr); i++ {
		arr[i-1] = arr[i]
	}

	arr[len(arr)-1] = temp
}

func LeftRotateByK(k int, arr []int) {
	if k < 0 {
		return
	}
	k = k % len(arr)

	temp := append(make([]int, 0), arr[:k]...)

	fmt.Println(temp)

	for i := k; i < len(arr); i++ {
		arr[i-k] = arr[i]
	}

	for i, v := range temp {
		arr[len(arr)-k+i] = v
	}
}

func LeftRotateByKoptimal(k int, arr []int) {
	if k < 0 {
		return
	}
	k = k%len(arr) - 1

	reverse(0, k, arr)
	reverse(k+1, len(arr)-1, arr)
	reverse(0, len(arr)-1, arr)
}

func reverse(start, end int, arr []int) {
	for start < end {
		temp := arr[start]
		arr[start] = arr[end]
		arr[end] = temp
		start++
		end--
	}
}

func MoveZeroToEnd(arr []int) {
	j := -1
	for i, v := range arr {
		if v == 0 {
			j = i
			break
		}
	}

	for i := j + 1; i < len(arr); {
		if arr[i] == 0 {
			i++
			continue
		}
		temp := arr[j]
		arr[j] = arr[i]
		arr[i] = temp
		j++
	}
}

func FindUnionOfSortedArray(arr1, arr2 []int) (result []int) {
	for _, v := range arr1 {
		if len(result) == 0 {
			result = append(result, v)
		}
		if v != result[len(result)-1] {
			result = append(result, v)
		}
	}
	for _, v := range arr2 {
		if len(result) == 0 {
			result = append(result, v)
		}
		if v > result[len(result)-1] && v != result[len(result)-1] {
			result = append(result, v)
		}
	}

	return
}
